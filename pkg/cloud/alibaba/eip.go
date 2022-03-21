package alibaba

import (
	"fmt"
	"strings"

	ecsClient "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vpcClient "github.com/alibabacloud-go/vpc-20160428/v2/client"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/galaxy-future/BridgX/pkg/utils"
	"github.com/spf13/cast"
)

func (p *AlibabaCloud) AllocateEip(req cloud.AllocateEipRequest) (ids []string, err error) {
	if req.Charge == nil {
		return nil, fmt.Errorf("charge is nil")
	}
	allocateEipAddressRequest := &vpcClient.AllocateEipAddressRequest{
		RegionId:  tea.String(req.RegionId),
		Bandwidth: tea.String(cast.ToString(req.Bandwidth)),
		ISP:       tea.String(req.InternetServiceProvider),
		AutoPay:   tea.Bool(true),
	}
	if req.Name != "" {
		allocateEipAddressRequest.Name = tea.String(req.Name)
	}
	insChargeType, internetChargeType := getEipChargeType(req.Charge.ChargeType)
	if insChargeType == "" {
		return nil, fmt.Errorf("invalid charge type")
	}
	allocateEipAddressRequest.InstanceChargeType = tea.String(insChargeType)
	allocateEipAddressRequest.InternetChargeType = tea.String(internetChargeType)
	if insChargeType == cloud.PrePaid {
		allocateEipAddressRequest.PricingCycle = tea.String(req.Charge.PeriodUnit)
		allocateEipAddressRequest.Period = tea.Int32(int32(req.Charge.Period))
	}

	idChan := make(chan string, req.Num)
	errChan := make(chan error, req.Num)
	for i := 0; i < req.Num; i++ {
		go func(req *vpcClient.AllocateEipAddressRequest) {
			rsp, err1 := p.vpcClient.AllocateEipAddress(req)
			if err1 != nil {
				errChan <- err1
				return
			}
			idChan <- utils.StringValue(rsp.Body.AllocationId)
		}(allocateEipAddressRequest)
	}

	for i := 0; i < req.Num; i++ {
		select {
		case err = <-errChan:
		case id := <-idChan:
			ids = append(ids, id)
		}
	}
	return ids, err
}

func (p *AlibabaCloud) GetEips(ids []string, regionId string) (map[string]cloud.Eip, error) {
	idNum := len(ids)
	eipMap := make(map[string]cloud.Eip, idNum)
	if idNum < 1 {
		return eipMap, nil
	}

	var maxNum int64 = 50
	batchIds := utils.StringSliceSplit(ids, maxNum)
	for _, onceIds := range batchIds {
		describeEipAddressesRequest := &vpcClient.DescribeEipAddressesRequest{
			RegionId:     tea.String(regionId),
			AllocationId: tea.String(strings.Join(onceIds, ",")),
		}
		rsp, err := p.vpcClient.DescribeEipAddresses(describeEipAddressesRequest)
		if err != nil {
			return nil, err
		}

		for _, v := range rsp.Body.EipAddresses.EipAddress {
			eipMap[tea.StringValue(v.AllocationId)] = eip2Cloud(v)
		}
	}
	return eipMap, nil
}

// ReleaseEip 包年包月类型的EIP不支持释放
func (p *AlibabaCloud) ReleaseEip(ids []string) (err error) {
	num := len(ids)
	if num < 1 {
		return nil
	}

	idChan := make(chan string, num)
	errChan := make(chan error, num)
	for _, id := range ids {
		go func(id string) {
			releaseEipAddressRequest := &vpcClient.ReleaseEipAddressRequest{
				AllocationId: tea.String(id),
			}
			_, err1 := p.vpcClient.ReleaseEipAddress(releaseEipAddressRequest)
			if err1 != nil {
				errChan <- err1
				return
			}
			idChan <- id
		}(id)
	}

	for i := 0; i < num; i++ {
		select {
		case err = <-errChan:
		case <-idChan:
		}
	}
	return err
}

func (p *AlibabaCloud) AssociateEip(id, instanceId string) error {
	associateEipAddressRequest := &vpcClient.AssociateEipAddressRequest{
		AllocationId: tea.String(id),
		InstanceId:   tea.String(instanceId),
		InstanceType: tea.String("EcsInstance"),
	}

	_, err := p.vpcClient.AssociateEipAddress(associateEipAddressRequest)
	if err != nil {
		return err
	}
	return nil
}

func (p *AlibabaCloud) DisassociateEip(id string) error {
	unassociateEipAddressRequest := &vpcClient.UnassociateEipAddressRequest{
		AllocationId: tea.String(id),
	}

	_, err := p.vpcClient.UnassociateEipAddress(unassociateEipAddressRequest)
	if err != nil {
		return err
	}
	return nil
}

func (p *AlibabaCloud) DescribeEip(req cloud.DescribeEipRequest) (cloud.DescribeEipResponse, error) {
	describeEipAddressesRequest := &vpcClient.DescribeEipAddressesRequest{
		RegionId: tea.String(req.RegionId),
	}
	if req.InstanceId != "" {
		describeEipAddressesRequest.AssociatedInstanceType = tea.String("EcsInstance")
		describeEipAddressesRequest.AssociatedInstanceId = tea.String(req.InstanceId)

	} else {
		if req.PageNum < 1 {
			req.PageNum = 1
		}
		if req.PageSize < 1 || req.PageSize > _pageSize {
			req.PageSize = _pageSize
		}
		describeEipAddressesRequest.Status = tea.String("Available")
		describeEipAddressesRequest.PageNumber = tea.Int32(int32(req.PageNum))
		describeEipAddressesRequest.PageSize = tea.Int32(int32(req.PageSize))
	}

	rsp, err := p.vpcClient.DescribeEipAddresses(describeEipAddressesRequest)
	if err != nil {
		return cloud.DescribeEipResponse{}, err
	}

	ret := cloud.DescribeEipResponse{
		List:       []cloud.Eip{},
		TotalCount: int(tea.Int32Value(rsp.Body.TotalCount)),
	}
	for _, v := range rsp.Body.EipAddresses.EipAddress {
		ret.List = append(ret.List, eip2Cloud(v))
	}
	return ret, nil
}

func (p *AlibabaCloud) ConvertPublicIpToEip(req cloud.ConvertPublicIpToEipRequest) error {
	convertNatPublicIpToEipRequest := &ecsClient.ConvertNatPublicIpToEipRequest{
		RegionId:   tea.String(req.RegionId),
		InstanceId: tea.String(req.InstanceId),
	}

	_, err := p.ecsClient.ConvertNatPublicIpToEip(convertNatPublicIpToEipRequest)
	if err != nil {
		return err
	}
	return nil
}

func eip2Cloud(eip *vpcClient.DescribeEipAddressesResponseBodyEipAddressesEipAddress) cloud.Eip {
	return cloud.Eip{
		Id:         tea.StringValue(eip.AllocationId),
		Name:       tea.StringValue(eip.Name),
		Ip:         tea.StringValue(eip.IpAddress),
		InstanceId: tea.StringValue(eip.InstanceId),
	}
}

func getEipChargeType(s string) (string, string) {
	switch s {
	case cloud.BandwidthPayByFix:
		return cloud.PostPaid, _payByBandwidth
	case cloud.BandwidthPayByTraffic:
		return cloud.PostPaid, _payByTraffic
	case cloud.BandwidthPrePaid:
		return cloud.PrePaid, _payByBandwidth
	}
	return "", ""
}
