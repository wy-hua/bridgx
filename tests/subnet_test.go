package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/galaxy-future/BridgX/cmd/api/request"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

const (
	_subnetPrefix = _v1Api + "subnet/"
)

func TestCreateSubnetAPI(t *testing.T) {
	tests := []request.CreateSwitchRequest{
		{
			SwitchName: "test_Switch",
			RegionId:   "bj",
			VpcId:      "vpc-i21un0x7mmtz",
			CidrBlock:  "192.168.1.0/24",
			GatewayIp:  "192.168.1.1",
			ZoneId:     "cn-bj-d",
			AK:         AKGenerator(cloud.BaiduCloud),
		},
		{
			SwitchName: "test_Switch",
			RegionId:   "cn-beijing",
			VpcId:      "vpc-2zexksa5gr5bxtufd61oz",
			CidrBlock:  "192.168.1.0/24",
			GatewayIp:  "192.168.1.1",
			ZoneId:     "cn-beijing-h",
			AK:         AKGenerator(cloud.AlibabaCloud),
		},
		{
			SwitchName: "test_Switch",
			RegionId:   "cn-north-1",
			VpcId:      "vpc-0d8c6a0bd621bf4c4",
			CidrBlock:  "10.0.0.0/24",
			GatewayIp:  "",
			ZoneId:     "cnn1-az1",
			AK:         AKGenerator(cloud.AWSCloud),
		},
	}
	for _, tt := range tests {
		t.Run(tt.SwitchName, func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _subnetPrefix+"create", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}

func TestDescribeSubnet(t *testing.T) {
	tests := []struct {
		name       string
		vpcId      string
		subnetName string
		zone       string
	}{
		{
			name:       "baidu",
			vpcId:      "vpc-i21un0x7mmtz",
			subnetName: "test_Switch",
			zone:       "cn-bj-d",
		},
		{
			name:       "aliyun",
			vpcId:      "vpc-2zexksa5gr5bxtufd61oz",
			subnetName: "test_Switch",
			zone:       "cn-beijing-h",
		},
		{
			name:       "aws",
			vpcId:      "vpc-0d8c6a0bd621bf4c4",
			subnetName: "test_Switch",
			zone:       "cnn1-az1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _subnetPrefix+fmt.Sprintf("describe?vpc_id=%s&switch_name=%s&zone_id=%s", tt.vpcId, tt.subnetName, tt.zone), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestGetSwitchById(t *testing.T) {
	tests := []struct {
		name     string
		switchId string
		vpcId    string
	}{
		{
			name:     "baidu",
			switchId: "sbn-6pk6bngtzvtg",
			vpcId:    "vpc-i21un0x7mmtz",
		},
		{
			name:     "aws",
			switchId: "subnet-09fe97713f59f89ef",
			vpcId:    "vpc-0d8c6a0bd621bf4c4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _subnetPrefix+fmt.Sprintf("info/%s?vpc_id=%s", tt.switchId, tt.vpcId), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
