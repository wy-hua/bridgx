package service

import (
	"context"
	"fmt"

	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/pkg/cloud"
)

const EipIdUnknown = "UNKNOWN"

type CloudAccount struct {
	Provider string `json:"provider" form:"provider"`
	RegionId string `json:"region_id" form:"region_id"`
	AK       string `json:"ak" form:"ak"`
}
type Eip struct {
	CloudAccount
	Id                      string        `json:"id"`
	Name                    string        `json:"name"`
	Ip                      string        `json:"ip"`
	InstanceId              string        `json:"instance_id"`
	InternetServiceProvider string        `json:"internet_service_provider"`
	Bandwidth               int           `json:"bandwidth"`
	Charge                  *cloud.Charge `json:"charge"`
}

func (p *Eip) CreateEip(ctx context.Context, num int) error {
	cloudCli, err := getProvider(p.Provider, p.AK, p.RegionId)
	if err != nil {
		logs.Logger.Errorf("getProvider failed, %v", err)
		return err
	}

	_, err = cloudCli.AllocateEip(cloud.AllocateEipRequest{
		RegionId:                p.RegionId,
		Name:                    p.Name,
		InternetServiceProvider: p.InternetServiceProvider,
		Bandwidth:               p.Bandwidth,
		Charge:                  p.Charge,
		Num:                     num,
	})
	if err != nil {
		logs.Logger.Errorf("AllocateEip failed, %v", err)
		return err
	}
	return nil
}

func (p *Eip) DescribeEip(ctx context.Context, pageNumber, pageSize int) (cloud.DescribeEipResponse, error) {
	cloudCli, err := getProvider(p.Provider, p.AK, p.RegionId)
	if err != nil {
		logs.Logger.Errorf("getProvider failed, %v", err)
		return cloud.DescribeEipResponse{}, err
	}

	rsp, err := cloudCli.DescribeEip(cloud.DescribeEipRequest{
		RegionId:   p.RegionId,
		InstanceId: p.InstanceId,
		PageNum:    pageNumber,
		PageSize:   pageSize,
	})
	if err != nil {
		logs.Logger.Errorf("DescribeEip failed, %v", err)
		return cloud.DescribeEipResponse{}, err
	}

	if p.InstanceId != "" {
		go func(instanceId string, eipList []cloud.Eip) {
			instance, err := GetInstance(ctx, instanceId)
			if err != nil {
				logs.Logger.Errorf("GetInstance failed, %v", err)
				return
			}
			if instance.EipId != EipIdUnknown || len(eipList) == 0 {
				return
			}
			instance.IpOuter = eipList[0].Ip
			instance.EipId = eipList[0].Id
			if err = model.Save(&instance); err != nil {
				logs.Logger.Errorf("model.Save failed, %v", err)
				return
			}
		}(p.InstanceId, rsp.List)
	}
	return rsp, nil
}

func (p *Eip) BindEip(ctx context.Context) error {
	instance, err := GetInstance(ctx, p.InstanceId)
	if err != nil {
		logs.Logger.Errorf("GetInstance failed, %v", err)
		return err
	}
	if instance.IpOuter != "" {
		return fmt.Errorf("public ip already exist")
	}

	cloudCli, err := getProvider(p.Provider, p.AK, p.RegionId)
	if err != nil {
		logs.Logger.Errorf("getProvider failed, %v", err)
		return err
	}
	eip, err := cloudCli.GetEips([]string{p.Id}, p.RegionId)
	if err != nil {
		logs.Logger.Errorf("GetEips failed, %v", err)
		return err
	}
	if len(eip) == 0 {
		return fmt.Errorf("eip id doesn't exist")
	}

	if err = cloudCli.AssociateEip(p.Id, p.InstanceId); err != nil {
		logs.Logger.Errorf("AssociateEip failed, %v", err)
		return err
	}

	instance.IpOuter = eip[p.Id].Ip
	instance.EipId = p.Id
	if err = model.Save(&instance); err != nil {
		logs.Logger.Errorf("model.Save failed, %v", err)
		return err
	}
	return nil
}

func (p *Eip) UnBindEip(ctx context.Context) error {
	instance, err := GetInstance(ctx, p.InstanceId)
	if err != nil {
		logs.Logger.Errorf("GetInstance failed, %v", err)
		return err
	}
	cloudCli, err := getProvider(p.Provider, p.AK, p.RegionId)
	if err != nil {
		logs.Logger.Errorf("getProvider failed, %v", err)
		return err
	}
	if p.Id == EipIdUnknown {
		rsp, err := cloudCli.DescribeEip(cloud.DescribeEipRequest{
			RegionId:   p.RegionId,
			InstanceId: p.InstanceId,
		})
		if err != nil {
			logs.Logger.Errorf("DescribeEip failed, %v", err)
			return err
		}
		if len(rsp.List) == 0 {
			return fmt.Errorf("DescribeEip rsp is empty")
		}
		p.Id = rsp.List[0].Id
	} else {
		eip, err := cloudCli.GetEips([]string{p.Id}, p.RegionId)
		if err != nil {
			logs.Logger.Errorf("GetEips failed, %v", err)
			return err
		}
		if eip[p.Id].InstanceId != p.InstanceId {
			return fmt.Errorf("eip and instance don't match")
		}
	}

	if err = cloudCli.DisassociateEip(p.Id); err != nil {
		logs.Logger.Errorf("DisassociateEip failed, %v", err)
		return err
	}
	instance.IpOuter = ""
	instance.EipId = ""
	if err = model.Save(&instance); err != nil {
		logs.Logger.Errorf("model.Save failed, %v", err)
		return err
	}
	return nil
}

func (p *Eip) ConvertPublicIp2Eip(ctx context.Context) error {
	instance, err := GetInstance(ctx, p.InstanceId)
	if err != nil {
		logs.Logger.Errorf("GetInstance failed, %v", err)
		return err
	}
	if instance.IpOuter == "" {
		return fmt.Errorf("public ip doesn't exist")
	}

	cloudCli, err := getProvider(p.Provider, p.AK, p.RegionId)
	if err != nil {
		logs.Logger.Errorf("getProvider failed, %v", err)
		return err
	}
	if err = cloudCli.ConvertPublicIpToEip(cloud.ConvertPublicIpToEipRequest{
		RegionId:   p.RegionId,
		InstanceId: p.InstanceId,
	}); err != nil {
		logs.Logger.Errorf("ConvertPublicIpToEip failed, %v", err)
		return err
	}
	rsp, err := cloudCli.DescribeEip(cloud.DescribeEipRequest{
		RegionId:   p.RegionId,
		InstanceId: p.InstanceId,
	})
	if err != nil {
		logs.Logger.Errorf("DescribeEip failed, %v", err)
		return err
	}

	if len(rsp.List) == 0 {
		instance.EipId = EipIdUnknown
	} else {
		eip := rsp.List[0]
		instance.IpOuter = eip.Ip
		instance.EipId = eip.Id
	}
	if err = model.Save(&instance); err != nil {
		logs.Logger.Errorf("model.Save failed, %v", err)
		return err
	}
	return nil
}
