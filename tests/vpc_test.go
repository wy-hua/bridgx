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
	_vpcPrefix = _v1Api + "vpc/"
)

func TestCreateVPCl(t *testing.T) {
	tests := []request.CreateVpcRequest{
		{
			Provider:  cloud.BaiduCloud,
			RegionId:  "bj",
			VpcName:   "test_vpc",
			CidrBlock: "192.168.0.0/16",
			AK:        AKGenerator(cloud.BaiduCloud),
		},
		{
			Provider:  cloud.AlibabaCloud,
			RegionId:  "cn-beijing",
			VpcName:   "test_vpc_1",
			CidrBlock: "192.168.0.0/16",
			AK:        AKGenerator(cloud.AlibabaCloud),
		},
		{
			Provider:  cloud.AWSCloud,
			RegionId:  "cn-north-1",
			VpcName:   "test_vpc",
			CidrBlock: "10.0.0.0/16",
			AK:        AKGenerator(cloud.AWSCloud),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Provider, func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _vpcPrefix+"create", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(1 * time.Minute)
		})
	}

}
func TestDescribeVPC(t *testing.T) {
	tests := []struct {
		name        string
		provider    string
		region_id   string
		vpc_name    string
		account_key string
	}{
		{
			name:        "baidu",
			provider:    cloud.BaiduCloud,
			region_id:   "bj",
			vpc_name:    "test_vpc",
			account_key: AKGenerator(cloud.BaiduCloud),
		},
		{
			name:        "aliyun",
			provider:    cloud.AlibabaCloud,
			region_id:   "cn-beijing",
			vpc_name:    "test_vpc_1",
			account_key: AKGenerator(cloud.AlibabaCloud),
		},
		{
			name:        "aws",
			provider:    cloud.AWSCloud,
			region_id:   "cn-north-1",
			vpc_name:    "test_vpc",
			account_key: AKGenerator(cloud.AWSCloud),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _vpcPrefix+fmt.Sprintf("describe?provider=%s&region_id=%s&vpc_name=%s&account_key=%s", tt.provider, tt.region_id, tt.vpc_name, tt.account_key), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestGetVpcById(t *testing.T) {
	tests := []struct {
		name  string
		vpcId string
	}{
		{
			name:  "baidu",
			vpcId: "vpc-i21un0x7mmtz",
		},
		{
			name:  "aws",
			vpcId: "vpc-0d8c6a0bd621bf4c4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _vpcPrefix+"info/"+tt.vpcId, nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
