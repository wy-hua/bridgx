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
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

const (
	_securotyGroupPrefix = _v1Api + "security_group/"
)

func TestSecurityGroupCreate(t *testing.T) {
	tests := []request.CreateSecurityGroupRequest{
		{
			AK:                AKGenerator(cloud.BaiduCloud),
			VpcId:             "vpc-i21un0x7mmtz",
			RegionId:          "bj",
			SecurityGroupName: "test_SecurityGroup",
			SecurityGroupType: "",
		},
		{
			AK:                AKGenerator(cloud.AlibabaCloud),
			VpcId:             "vpc-2zexksa5gr5bxtufd61oz",
			RegionId:          "cn-beijing",
			SecurityGroupName: "test_SecurityGroup",
			SecurityGroupType: "",
		},
		{
			AK:                AKGenerator(cloud.AwsCloud),
			VpcId:             "vpc-0d8c6a0bd621bf4c4",
			RegionId:          "cn-north-1",
			SecurityGroupName: "test_SecurityGroup",
			SecurityGroupType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.VpcId, func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _securotyGroupPrefix+"create", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestDescribeSecurityGroup(t *testing.T) {
	tests := []struct {
		vpcId             string
		securityGroupName string
		accountKey        string
	}{
		{
			vpcId:             "vpc-i21un0x7mmtz",
			securityGroupName: "g-xy2ttwa9hqsb",
			accountKey:        AKGenerator(cloud.BaiduCloud),
		},
		{
			vpcId:             "vpc-0d8c6a0bd621bf4c4",
			securityGroupName: "test_SecurityGroup",
			accountKey:        AKGenerator(cloud.AwsCloud),
		},
	}
	for _, tt := range tests {
		t.Run(tt.securityGroupName, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _securotyGroupPrefix+fmt.Sprintf("describe?vpc_id=%s&security_group_name=%s&account_key=%s", tt.vpcId, tt.securityGroupName, tt.accountKey), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestAddSecurityGroupRuleAPI(t *testing.T) {
	tests := []request.AddSecurityGroupRuleRequest{
		{
			AK:              "xx",
			VpcId:           "vpc-i21un0x7mmtz",
			RegionId:        "bj",
			SecurityGroupId: "g-xy2ttwa9hqsb",
			Rules: []service.GroupRule{
				{
					Protocol:     "tcp",
					PortFrom:     1024,
					PortTo:       2048,
					Direction:    "ingress",
					GroupId:      "g-xy2ttwa9hqsb",
					CidrIp:       "192.168.1.0/24",
					PrefixListId: "",
				},
			},
		},
		{
			AK:              AKGenerator(cloud.AlibabaCloud),
			VpcId:           "vpc-2zexksa5gr5bxtufd61oz",
			RegionId:        "cn-beijing",
			SecurityGroupId: "sg-2ze8na66wsies9tkfd3w",
			Rules: []service.GroupRule{
				{
					Protocol:     "tcp",
					PortFrom:     22,
					PortTo:       22,
					Direction:    "ingress",
					GroupId:      "sg-2ze8na66wsies9tkfd3w",
					CidrIp:       "192.168.1.0/24",
					PrefixListId: "",
				},
			},
		},
		{
			AK:              AKGenerator(cloud.AwsCloud),
			VpcId:           "vpc-0d8c6a0bd621bf4c4",
			RegionId:        "cn-north-1",
			SecurityGroupId: "sg-07cdd57dd38d31672",
			Rules: []service.GroupRule{
				{
					Protocol:     "tcp",
					PortFrom:     1024,
					PortTo:       2048,
					Direction:    "ingress",
					GroupId:      "sg-07cdd57dd38d31672",
					CidrIp:       "10.0.0.0/24",
					PrefixListId: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.SecurityGroupId, func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _securotyGroupPrefix+"rule/add", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestGetSecurityGroupWithRules(t *testing.T) {
	tests := []struct {
		securityGroupId string
	}{
		{
			securityGroupId: "g-xy2ttwa9hqsb",
		},
		{
			securityGroupId: "sg-07cdd57dd38d31672",
		},
	}
	for _, tt := range tests {
		t.Run(tt.securityGroupId, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _securotyGroupPrefix+tt.securityGroupId+"/rules", nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
