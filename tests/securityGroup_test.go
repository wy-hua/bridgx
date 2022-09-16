package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
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
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
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
			Rules: []service.GroupRule{{
				Protocol:     "tcp",
				PortFrom:     1024,
				PortTo:       2048,
				Direction:    "ingress",
				GroupId:      "g-xy2ttwa9hqsb",
				CidrIp:       "192.168.1.0/24",
				PrefixListId: ""},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
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
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
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
