package tencent

import (
	"fmt"

	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/galaxy-future/BridgX/pkg/utils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

func (p *TencentCloud) AllocateEip(req cloud.AllocateEipRequest) (ids []string, err error) {
	if req.Charge == nil {
		return nil, fmt.Errorf("charge is nil")
	}
	request := vpc.NewAllocateAddressesRequest()
	request.AddressCount = common.Int64Ptr(int64(req.Num))
	request.InternetServiceProvider = common.StringPtr(req.InternetServiceProvider)
	request.InternetChargeType = common.StringPtr(_eipChargeType[req.Charge.ChargeType])
	request.InternetMaxBandwidthOut = common.Int64Ptr(int64(req.Bandwidth))
	if req.Name != "" {
		request.AddressName = common.StringPtr(req.Name)
	}
	if req.Charge.ChargeType == cloud.BandwidthPrePaid {
		request.AddressChargePrepaid = &vpc.AddressChargePrepaid{
			Period:        common.Int64Ptr(int64(req.Charge.Period)),
			AutoRenewFlag: common.Int64Ptr(1),
		}
	}

	rsp, err := p.vpcClient.AllocateAddresses(request)
	if err != nil {
		return nil, err
	}
	for _, v := range rsp.Response.AddressSet {
		ids = append(ids, utils.StringValue(v))
	}
	return ids, nil
}

func (p *TencentCloud) GetEips(ids []string, regionId string) (map[string]cloud.Eip, error) {
	idNum := len(ids)
	eipMap := make(map[string]cloud.Eip, idNum)
	if idNum < 1 {
		return eipMap, nil
	}

	request := vpc.NewDescribeAddressesRequest()
	request.AddressIds = common.StringPtrs(ids)
	rsp, err := p.vpcClient.DescribeAddresses(request)
	if err != nil {
		return nil, err
	}

	for _, v := range rsp.Response.AddressSet {
		eipMap[utils.StringValue(v.AddressId)] = eip2Cloud(v)
	}
	return eipMap, nil
}

func (p *TencentCloud) ReleaseEip(ids []string) (err error) {
	request := vpc.NewReleaseAddressesRequest()
	request.AddressIds = common.StringPtrs(ids)

	_, err = p.vpcClient.ReleaseAddresses(request)
	if err != nil {
		return err
	}
	return nil
}

func (p *TencentCloud) AssociateEip(id, instanceId string) error {
	request := vpc.NewAssociateAddressRequest()
	request.AddressId = common.StringPtr(id)
	request.InstanceId = common.StringPtr(instanceId)

	_, err := p.vpcClient.AssociateAddress(request)
	if err != nil {
		return err
	}
	return nil
}

func (p *TencentCloud) DisassociateEip(id string) error {
	request := vpc.NewDisassociateAddressRequest()
	request.AddressId = common.StringPtr(id)

	_, err := p.vpcClient.DisassociateAddress(request)
	if err != nil {
		return err
	}
	return nil
}

func (p *TencentCloud) DescribeEip(req cloud.DescribeEipRequest) (cloud.DescribeEipResponse, error) {
	request := vpc.NewDescribeAddressesRequest()
	if req.InstanceId != "" {
		request.Filters = []*vpc.Filter{
			{
				Name:   common.StringPtr("instance-id"),
				Values: common.StringPtrs([]string{req.InstanceId}),
			},
		}
	} else {
		if req.PageNum < 1 {
			req.PageNum = 1
		}
		if req.PageSize < 1 || req.PageSize > _pageSize {
			req.PageSize = _pageSize
		}
		request.Offset = common.Int64Ptr(int64((req.PageNum - 1) * req.PageSize))
		request.Limit = common.Int64Ptr(int64(req.PageSize))
		request.Filters = []*vpc.Filter{
			{
				Name:   common.StringPtr("address-status"),
				Values: common.StringPtrs([]string{"UNBIND", "BIND_ENI"}),
			},
		}
	}

	rsp, err := p.vpcClient.DescribeAddresses(request)
	if err != nil {
		return cloud.DescribeEipResponse{}, err
	}

	ret := cloud.DescribeEipResponse{
		List:       []cloud.Eip{},
		TotalCount: int(utils.Int64Value(rsp.Response.TotalCount)),
	}
	for _, v := range rsp.Response.AddressSet {
		ret.List = append(ret.List, eip2Cloud(v))
	}
	return ret, nil
}

func (p *TencentCloud) ConvertPublicIpToEip(req cloud.ConvertPublicIpToEipRequest) error {
	request := vpc.NewTransformAddressRequest()
	request.InstanceId = common.StringPtr(req.InstanceId)

	_, err := p.vpcClient.TransformAddress(request)
	if err != nil {
		return err
	}
	return nil
}

func eip2Cloud(eip *vpc.Address) cloud.Eip {
	return cloud.Eip{
		Id:         utils.StringValue(eip.AddressId),
		Name:       utils.StringValue(eip.AddressName),
		Ip:         utils.StringValue(eip.AddressIp),
		InstanceId: utils.StringValue(eip.InstanceId),
	}
}
