package huawei

import (
	"github.com/galaxy-future/BridgX/pkg/cloud"
)

func (p *HuaweiCloud) AllocateEip(req cloud.AllocateEipRequest) (ids []string, err error) {

	return nil, nil
}

func (p *HuaweiCloud) GetEips(ids []string, regionId string) (map[string]cloud.Eip, error) {
	idNum := len(ids)
	eipMap := make(map[string]cloud.Eip, idNum)
	if idNum < 1 {
		return eipMap, nil
	}
	return eipMap, nil
}

func (p *HuaweiCloud) ReleaseEip(ids []string) (err error) {

	return nil
}

func (p *HuaweiCloud) AssociateEip(id, instanceId string) error {

	return nil
}

func (p *HuaweiCloud) DisassociateEip(id string) error {

	return nil
}

func (p *HuaweiCloud) DescribeEip(req cloud.DescribeEipRequest) (cloud.DescribeEipResponse, error) {
	ret := cloud.DescribeEipResponse{}
	return ret, nil
}

func (p *HuaweiCloud) ConvertPublicIpToEip(req cloud.ConvertPublicIpToEipRequest) error {

	return nil
}
