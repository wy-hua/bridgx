package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/galaxy-future/BridgX/cmd/api/request"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/stretchr/testify/assert"

	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/service"
)

const (
	_networkPrefix = _v1Api + "network_config/"
)

func TestCreateNetworkConfig(t *testing.T) {
	tests := []request.CreateNetworkRequest{
		{
			Provider:          cloud.BaiduCloud,
			RegionId:          "bj",
			CidrBlock:         "192.168.0.0/16",
			VpcName:           "network_config_vpc",
			ZoneId:            "cn-bj-d",
			SwitchCidrBlock:   "192.168.0.0/24",
			GatewayIp:         "192.168.0.1",
			SwitchName:        "network_config_switch",
			SecurityGroupName: "network_config_security_group",
			SecurityGroupType: "",
			AK:                AKGenerator(cloud.BaiduCloud),
			Rules: []service.GroupRule{
				{
					Protocol:     "tcp",
					PortFrom:     1024,
					PortTo:       2048,
					Direction:    "ingress",
					GroupId:      "",
					CidrIp:       "192.168.0.0/24",
					PrefixListId: "",
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _networkPrefix+"create", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestSyncNetworkConfig(t *testing.T) {
	tests := []request.SyncNetworkRequest{
		{
			Provider:   cloud.BaiduCloud,
			RegionId:   "bj",
			AccountKey: AKGenerator(cloud.BaiduCloud),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _networkPrefix+"sync", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestGetNetCfgTemplate(t *testing.T) {
	var tests = []string{
		"BaiduCloud",
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _networkPrefix+"template?provider="+tt, nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestCreateVPC(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.CreateVPCRequest
	}
	tests := []struct {
		name      string
		args      args
		wantVpcId string
		wantErr   bool
	}{
		{
			name: "测试创建vpc",
			args: args{
				ctx: nil,
				req: service.CreateVPCRequest{
					Provider:  "AlibabaCloud",
					RegionId:  "cn-qingdao",
					VpcName:   "vpc测试自动更新",
					CidrBlock: "",
					AK:        "LTAI5t7qCv6L8ZFh3hzSYpSv",
				},
			},
			wantVpcId: "",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVpcId, err := service.CreateVPC(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateVPC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVpcId == "" {
				t.Errorf("CreateVPC() gotVpcId = %v, want %v", gotVpcId, tt.wantVpcId)
			}
		})
	}
}

func TestCreateSwitch(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.CreateSwitchRequest
	}
	tests := []struct {
		name         string
		args         args
		wantSwitchId string
		wantErr      bool
	}{
		{
			name: "创建交换机",
			args: args{
				ctx: nil,
				req: service.CreateSwitchRequest{
					SwitchName: "第一台交换机",
					ZoneId:     "cn-qingdao-b",
					VpcId:      "vpc-m5ey3pofeclswmv796tgd",
					CidrBlock:  "172.16.0.0/24",
				},
			},
			wantSwitchId: "",
			wantErr:      false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSwitchId, err := service.CreateSwitch(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSwitch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSwitchId == tt.wantSwitchId {
				t.Errorf("CreateSwitch() gotSwitchId = %v, want %v", gotSwitchId, tt.wantSwitchId)
			}
		})
	}
}

func TestCreateSecurityGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.CreateSecurityGroupRequest
	}
	tests := []struct {
		name                string
		args                args
		wantSecurityGroupId string
		wantErr             bool
	}{
		{
			name: "测试创建安全组",
			args: args{
				ctx: nil,
				req: service.CreateSecurityGroupRequest{
					VpcId:             "vpc-m5ey3pofeclswmv796tgd",
					SecurityGroupName: "测试的第一个安全组",
					SecurityGroupType: "normal",
				},
			},
			wantSecurityGroupId: "",
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSecurityGroupId, err := service.CreateSecurityGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSecurityGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSecurityGroupId == tt.wantSecurityGroupId {
				t.Errorf("CreateSecurityGroup() gotSecurityGroupId = %v, want %v", gotSecurityGroupId, tt.wantSecurityGroupId)
			}
		})
	}
}

func TestAddSecurityGroupRule(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.AddSecurityGroupRuleRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试添加一条入安全组策略",
			args: args{
				ctx: nil,
				req: service.AddSecurityGroupRuleRequest{
					RegionId:        "cn-qingdao",
					VpcId:           "vpc-m5ey3pofeclswmv796tgd",
					SecurityGroupId: "sg-m5ebcpsx22mv2cu2f5x8",
					//Protocol:        "tcp",
					//PortRange:       "80/80",
					//Direction:       service.DirectionIn,
					//CidrIp:          "0.0.0.0/0",
				},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.AddSecurityGroupRule(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSecurityGroupRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddSecurityGroupRule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNetwork(t *testing.T) {
	type args struct {
		ctx context.Context
		req *service.CreateNetworkRequest
	}
	tests := []struct {
		name       string
		args       args
		wantVpcRes service.CreateNetworkResponse
		wantErr    bool
	}{
		{
			name: "测试一键创建",
			args: args{
				ctx: nil,
				req: &service.CreateNetworkRequest{
					Provider:          "AlibabaCloud",
					RegionId:          "cn-qingdao",
					VpcName:           "测试一键创建",
					ZoneId:            "cn-qingdao-b",
					SwitchCidrBlock:   "172.16.0.0/24",
					SwitchName:        "一键创建的switch",
					SecurityGroupName: "一键创建的安全组",
					SecurityGroupType: "normal",
					AK:                "xxx",
				},
			},
			wantVpcRes: service.CreateNetworkResponse{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVpcRes, err := service.CreateNetwork(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(gotVpcRes, tt.wantVpcRes) {
				t.Errorf("CreateNetwork() gotVpcRes = %v, want %v", gotVpcRes, tt.wantVpcRes)
			}
		})
	}
}

func TestGetVPC(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.GetVPCRequest
	}
	tests := []struct {
		name     string
		args     args
		wantVpcs []model.Vpc
		wantErr  bool
	}{
		{
			name: "测试 Vpc 查询",
			args: args{
				ctx: nil,
				req: service.GetVPCRequest{
					Provider:   "AlibabaCloud",
					RegionId:   "cn-qingdao",
					VpcName:    "测试一键创建",
					PageNumber: 0,
					PageSize:   20,
					AccountKey: "",
				},
			},
			wantVpcs: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVpc, err := service.GetVPC(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVPC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotVpc.Vpcs) == 0 {
				t.Errorf("GetVPC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSwitch(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.GetSwitchRequest
	}
	tests := []struct {
		name         string
		args         args
		wantSwitches []model.Switch
		wantErr      bool
	}{
		{
			name: "测试查询交换机",
			args: args{
				ctx: nil,
				req: service.GetSwitchRequest{
					SwitchName: "一键创建的switch",
					VpcId:      "vpc-m5e9zap8y3afp3aatzsje",
					PageNumber: 0,
					PageSize:   0,
				},
			},
			wantSwitches: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSwitches, err := service.GetSwitch(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSwitch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotSwitches)
			if len(gotSwitches.Switches) < 1 {
				t.Errorf("GetSwitch() gotSwitches = %v, want %v", gotSwitches, tt.wantSwitches)
			}
		})
	}
}

func TestGetSecurityGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req service.GetSecurityGroupRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []model.SecurityGroup
		wantErr bool
	}{
		{
			name: "测试查询安全组",
			args: args{
				ctx: nil,
				req: service.GetSecurityGroupRequest{
					SecurityGroupName: "一键创建的安全组",
					VpcId:             "vpc-m5e9zap8y3afp3aatzsje",
					PageNumber:        0,
					PageSize:          0,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetSecurityGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecurityGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			if len(got.Groups) < 1 {
				t.Errorf("GetSecurityGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//使用本函数前需要将 updateOrCreateSecurityGroupRules 改为同步调用
func TestSyncNetwork(t *testing.T) {
	accounts := make([]model.Account, 0)
	if err := model.QueryAll(map[string]interface{}{}, &accounts, ""); err != nil {
		t.Log("query account,", err)
		return
	}

	for _, account := range accounts {
		err := service.RefreshAccount(&service.SimpleTask{
			ProviderName: account.Provider,
			AccountKey:   account.AccountKey,
		})
		if err != nil {
			t.Log(account.AccountKey, err)
			continue
		}
	}
}
