package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/galaxy-future/BridgX/cmd/api/request"
	"github.com/galaxy-future/BridgX/pkg/cloud"

	"github.com/galaxy-future/BridgX/internal/constants"
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestListInstanceType(t *testing.T) {
	tests := []struct {
		provider string
		regionId string
		zoneId   string
	}{
		{
			provider: cloud.BaiduCloud,
			regionId: "bj",
			zoneId:   "cn-bj-d",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _v1Api+fmt.Sprintf("instance_type/list?provider=%s&region_id=%s&zone_id=%s", tt.provider, tt.regionId, tt.zoneId), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestSyncInstanceExpireTimeAPI(t *testing.T) {
	tests := []request.SyncInstanceExpireTimeRequest{
		{
			ClusterName: "test_cluster",
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _v1Api+"instance/sync_expire_time", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}

func TestCreateInstance(t *testing.T) {
	instances := make([]model.Instance, 0)
	now := time.Now()
	instances = append(instances, model.Instance{
		Base: model.Base{
			CreateAt: &now,
		},
		InstanceId:  "test1",
		Status:      constants.Pending,
		ClusterName: "test_cluster1",
		DeleteAt:    &now,
	})
	instances = append(instances, model.Instance{
		Base: model.Base{
			CreateAt: &now,
		},
		InstanceId:  "test2",
		Status:      constants.Pending,
		ClusterName: "test_cluster2",
		DeleteAt:    &now,
	})
	err := model.BatchCreateInstance(instances)
	t.Log(err)
}

func TestUpdateInstance(t *testing.T) {
	instance1 := model.Instance{
		InstanceId:  "test1",
		Status:      constants.Running,
		ClusterName: "test_cluster1",
		IpInner:     "10.0.0.1",
	}
	err := model.UpdateByInstanceId(instance1)
	t.Log(err)

	instance2 := model.Instance{
		InstanceId: "test2",
		Status:     constants.Deleted,
	}
	err = model.UpdateByInstanceId(instance2)
	t.Log(err)
}

func TestGetInstanceByIp(t *testing.T) {
	instance, err := model.GetInstanceByIpInner("10.0.0.1")
	t.Log(err)
	t.Log(instance)
}

func TestGetInstanceByIps(t *testing.T) {
	ips := []string{"10.0.0.1", "10.0.0.2"}
	instances, err := model.GetInstancesByIPs(ips, "")
	t.Log(err)
	t.Log(instances)
}

func TestBatchUpdate(t *testing.T) {
	instanceIds := []string{"test1", "test2"}
	err := model.BatchUpdateByInstanceIds(instanceIds, model.Instance{Status: constants.Deleted})
	t.Log(err)
}

func TestSyncInstanceExpireTime(t *testing.T) {
	err := service.SyncInstanceExpireTime(context.Background(), "gf.cloud_4")
	assert.Nil(t, err)
}
