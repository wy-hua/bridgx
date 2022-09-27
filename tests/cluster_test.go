package tests

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/galaxy-future/BridgX/cmd/api/request"
	"github.com/galaxy-future/BridgX/internal/constants"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/service"
	"github.com/galaxy-future/BridgX/internal/types"
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/galaxy-future/BridgX/pkg/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const (
	_cluster = _v1Api + `cluster/`
	_keyPair = _v1Api + `key_pair/`
	_vpc     = "vpc-iyneigjruzy5"
)

func TestCreate(t *testing.T) {
	tests := []types.ClusterInfo{
		{
			Id:           1,
			Name:         "test_cluster",
			Desc:         "no description",
			RegionId:     "bj",
			ZoneId:       "cn-bj-d",
			ClusterType:  "",
			InstanceType: "bcc.ic4.c1m1",
			Image:        "m-OWQC4wwM",
			Provider:     cloud.BaiduCloud,
			Username:     "root",
			Password:     "I1235677!",
			AccountKey:   AKGenerator(cloud.BaiduCloud),
			ImageConfig: &types.ImageConfig{
				Id:       "m-OWQC4wwM",
				Name:     "7.9 x86_64 (64bit)",
				Type:     "global",
				Platform: "Centos",
				Size:     0,
			},

			NetworkConfig: &types.NetworkConfig{
				Vpc:                     "vpc-i21un0x7mmtz",
				SubnetId:                "sbn-mgiqutgye6ui",
				SecurityGroup:           "g-xy2ttwa9hqsb",
				InternetChargeType:      "",
				InternetMaxBandwidthOut: 10,
				InternetIpType:          "",
			},
			StorageConfig: &types.StorageConfig{
				MountPoint: "",
				NAS:        "",
				Disks: &cloud.Disks{
					SystemDisk: cloud.DiskConf{
						Category:         "enhanced_ssd_pl1",
						Size:             20,
						PerformanceLevel: "",
					},
					DataDisk: nil,
				},
			},
			ChargeConfig: &types.ChargeConfig{ChargeType: "PostPaid"},
			ExtendConfig: &types.ExtendConfig{
				Core:    1,
				Memory:  1,
				CpuType: "cpu",
			},
			Tags: map[string]string{"myTest": "1"},
		},
		{
			Id:           2,
			Name:         "test_import_key",
			Desc:         "no description",
			RegionId:     "cn-north-1",
			ZoneId:       "cnn1-az1",
			ClusterType:  "",
			InstanceType: "t2.micro",
			Image:        "ami-0a5e581c2158fe57d",
			Provider:     cloud.AwsCloud,
			Username:     "",
			Password:     "Ivgg87892789!",
			AccountKey:   AKGenerator(cloud.AwsCloud),
			ImageConfig: &types.ImageConfig{
				Id:       "ami-0a5e581c2158fe57d",
				Name:     "AWS Deep Learning AMI GPU CUDA 11 (Ubuntu 20.04) 20220317",
				Type:     "global",
				Platform: "Linux/UNIX",
				Size:     0,
			},

			NetworkConfig: &types.NetworkConfig{
				Vpc:                     "vpc-0d8c6a0bd621bf4c4",
				SubnetId:                "subnet-09fe97713f59f89ef",
				SecurityGroup:           "sg-07cdd57dd38d31672",
				InternetChargeType:      "",
				InternetMaxBandwidthOut: 10,
				InternetIpType:          "",
			},
			StorageConfig: &types.StorageConfig{
				MountPoint: "",
				NAS:        "",
				Disks: &cloud.Disks{
					SystemDisk: cloud.DiskConf{
						Category:         "gp2",
						Size:             100,
						PerformanceLevel: "",
					},
					DataDisk: nil,
				},
			},
			ChargeConfig: &types.ChargeConfig{
				ChargeType: "PostPaid",
				Period:     1,
				PeriodUnit: "Month",
			},
			ExtendConfig: &types.ExtendConfig{
				Core:    1,
				Memory:  1,
				CpuType: "cpu",
			},
			KeyPairId:   "key-xx",
			KeyId:       "xx",
			KeyPairName: "import_key_pair",
			AuthType:    constants.AuthTypeKeyPair,
			Tags:        map[string]string{"myTest": "1"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			json, _ := json.Marshal(tt)
			req, _ := http.NewRequest("POST", _cluster+`create`, bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestCreateKeyPair(t *testing.T) {
	tests := []request.CreateKeyPairRequest{
		{
			AK:          AKGenerator(cloud.AwsCloud),
			Provider:    cloud.AwsCloud,
			RegionId:    "cn-north-1",
			KeyPairName: "test_key_pair",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _keyPair+"create", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}
}
func TestGetClusterAuth(t *testing.T) {
	tests := []request.ClusterAuthRequest{
		{
			ClusterName: "aws_cluster_test",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _cluster+"auth?cluster_name="+tt.ClusterName, nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}
func TestImportKeyPair(t *testing.T) {
	tests := []request.ImportKeyPairRequest{
		{
			AK:          AKGenerator(cloud.AwsCloud),
			Provider:    cloud.AwsCloud,
			RegionId:    "cn-north-1",
			KeyPairName: "import_key_pair",
			PrivateKey:  `__test_private_key_content__`,
			PublicKey:   `__test_public_key_content__`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _keyPair+"import", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}
}
func TestGetClusterByName(t *testing.T) {
	tests := []string{"aws_cluster_test"}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _cluster+"name/"+tt, nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}

}

func TestExpandCluster(t *testing.T) {
	tests := []request.ExpandClusterRequest{
		{
			TaskName:    "task1",
			ClusterName: "clusterA",
			Count:       2,
		},
		{
			TaskName:    "task2",
			ClusterName: "test_ali_cluster",
			Count:       2,
		},
		{
			TaskName:    "task3",
			ClusterName: "aws_cluster_test",
			Count:       2,
		},
		{
			TaskName:    "task4",
			ClusterName: "test_import_key",
			Count:       2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _cluster+"expand", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(1 * time.Minute)
		})
	}

}
func TestShrinkCluster(t *testing.T) {
	tests := []request.ShrinkClusterRequest{
		{
			TaskName:    "task1",
			ClusterName: "test_cluster",
			Count:       2,
		},
		{
			TaskName:    "task2",
			ClusterName: "test_ali_cluster",
			Count:       2,
		},
		{
			TaskName:    "task3",
			ClusterName: "aws_cluster_test",
			Count:       2,
		},
		{
			TaskName:    "task4",
			ClusterName: "clusterA",
			Count:       2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			json, _ := json.Marshal(tt)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", _cluster+"shrink", bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(1 * time.Minute)
		})
	}

}

func TestGetClusterTags(t *testing.T) {
	tests := []struct {
		clusterName string
		tagKey      string
		pageNum     string
		pageSize    string
	}{
		{
			clusterName: "test_cluster",
			tagKey:      "myTest",
			pageNum:     "1",
			pageSize:    "100",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", _cluster+fmt.Sprintf("get_tags?cluster_name=%s&tag_key=%s&page_number=%s&page_size=%s", tt.clusterName, tt.tagKey, tt.pageNum, tt.pageSize), nil)
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestEditCluster(t *testing.T) {
	tests := []types.ClusterInfo{
		{
			Id:           1,
			Name:         "test_cluster",
			Desc:         "edit",
			RegionId:     "bj",
			ZoneId:       "cn-bj-d",
			ClusterType:  "",
			InstanceType: "bcc.ic4.c1m1",
			Image:        "centos",
			Provider:     cloud.BaiduCloud,
			Username:     "root",
			Password:     "Idfjafh81!",
			AccountKey:   AKGenerator(cloud.BaiduCloud),
			ImageConfig: &types.ImageConfig{
				Id: "m-OWQC4wwM",
			},

			NetworkConfig: &types.NetworkConfig{
				Vpc:                     "xx",
				SubnetId:                "sbn-6pk6bngtzvtg",
				SecurityGroup:           "g-xy2ttwa9hqsb",
				InternetChargeType:      "",
				InternetMaxBandwidthOut: 10,
				InternetIpType:          "",
			},
			StorageConfig: &types.StorageConfig{
				MountPoint: "",
				NAS:        "",
				Disks: &cloud.Disks{
					SystemDisk: cloud.DiskConf{
						Category:         "enhanced_ssd_pl1",
						Size:             40,
						PerformanceLevel: "",
					},
					DataDisk: nil,
				},
			},
			ChargeConfig: &types.ChargeConfig{ChargeType: "PostPaid"},
			ExtendConfig: &types.ExtendConfig{
				Core:    1,
				Memory:  1,
				CpuType: "cpu",
			},
			Tags: map[string]string{"myTest": "1"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			json, _ := json.Marshal(tt)
			req, _ := http.NewRequest("POST", _cluster+`edit`, bytes.NewReader(json))
			req.Header.Set("Authorization", "Bear "+_Token)
			req.Header.Set("content-type", "application/json")
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			assert.Equal(t, 200, w.Code)
			time.Sleep(7 * time.Second)
		})
	}

}
func TestCreateCluster(t *testing.T) {
	err := service.CreateCluster4Test("TEST_CLUSTER")
	assert.Nil(t, err, "should be nil", err)
}

func TestCreateClusterByApi(t *testing.T) {
	cluster := types.ClusterInfo{
		Name: fmt.Sprintf("gf.metrics.pi"),
		//Name:         fmt.Sprintf("gf.metrics.pi.cluster-%v", time.Now().Unix()),
		RegionId: "cn-qingdao",
		//ZoneId:       "cn-beijing-h",
		InstanceType: "ecs.g6.large",
		Image:        "centos_8_4_uefi_x64_20G_alibase_20210611.vhd",
		Provider:     cloud.AlibabaCloud,
		Password:     "xxx",
		AccountKey:   "xxx",
		NetworkConfig: &types.NetworkConfig{
			Vpc:           "vpc-2zelmmlfd5c5duibc2xb2",
			SubnetId:      "vsw-2zennaxawzq6sa2fdj8l5",
			SecurityGroup: "sg-2zefbt9tw0yo1r7vc3ac",
		},
		StorageConfig: &types.StorageConfig{
			Disks: &cloud.Disks{
				SystemDisk: cloud.DiskConf{Size: 40, Category: "cloud_efficiency"},
				DataDisk: []cloud.DiskConf{{
					Size:     100,
					Category: "cloud_efficiency",
				}},
			},
		},
		ChargeConfig: &types.ChargeConfig{ChargeType: "PostPaid"},
	}
	b, _ := jsoniter.MarshalToString(cluster)
	t.Logf(b)
	ret, _ := utils.HttpPostJsonDataT("http://0.0.0.0:9090/api/v1/cluster/create", []byte(b), 3)
	t.Logf("Response:%v", string(ret))
}

func TestCreateClusterTagsByApi(t *testing.T) {
	name := time.Now().Unix()
	req := fmt.Sprintf(`{"name":"Cluster-%v","desc":"k","region_id":"cn-bj","zone_id":"cn-bj-h","instance_type":"2c4g","charge_type":"by_month","network_config":{"vpc":"vpc-ikw1swp1"},"storage_config":{"mountPoint":"/opt/data","nas":""},"tags":{"dc":"lf","env":"prod"}}`, name)
	ret, err := utils.HttpPostJsonDataT("http://0.0.0.0:9090/api/v1/cluster/create", []byte(req), 3)
	t.Logf("Response:%v", string(ret))
	assert.Nil(t, err, "err not nil")
	tagReq := fmt.Sprintf(`{"cluster_name": "Cluster-%v", "tags": {"k1": "v1", "k2": "v2"}}`, name)
	ret2, err := utils.HttpPostJsonDataT("http://0.0.0.0:9090/api/v1/cluster/add_tags", []byte(tagReq), 3)
	t.Logf("Response:%v", string(ret2))
	assert.Nil(t, err, "err not nil")
}

func TestCreateClusterErr(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = service.CreateCluster4Test("TEST_CLUSTER")
		r := 100 + rand.Int31n(50)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}
}

func TestExpandClusterUseMockCluster(t *testing.T) {
	cluster := types.ClusterInfo{
		Name:         fmt.Sprintf("cluster-%v", time.Now()),
		RegionId:     "cn-beijing",
		ZoneId:       "cn-beijing-h",
		InstanceType: "ecs.s6-c1m1.small",
		Image:        "centos_7_9_x64_20G_alibase_20210623.vhd",
		Provider:     cloud.AlibabaCloud,
		Password:     "xxx",
		AccountKey:   "xxx",
		NetworkConfig: &types.NetworkConfig{
			Vpc:           "vpc-2zelmmlfd5c5duibc2xb2",
			SubnetId:      "vsw-2zennaxawzq6sa2fdj8l5",
			SecurityGroup: "sg-2zefbt9tw0yo1r7vc3ac",
		},
		StorageConfig: &types.StorageConfig{
			Disks: &cloud.Disks{
				SystemDisk: cloud.DiskConf{Size: 40, Category: "cloud_efficiency"},
				DataDisk: []cloud.DiskConf{{
					Size:     100,
					Category: "cloud_efficiency",
				}},
			},
		},
		ChargeConfig: &types.ChargeConfig{ChargeType: "PostPaid"},
	}
	instanceIds, err := service.Expand(&cluster, nil, 2)
	t.Logf("instanceIds: %v", strings.Join(instanceIds, ","))
	t.Log("err: ", err)
}

func TestGetInstance(t *testing.T) {
	cluster := types.ClusterInfo{
		RegionId:   "cn-beijing",
		Provider:   cloud.AlibabaCloud,
		AccountKey: "xxx",
	}
	res, err := service.GetInstances(&cluster, []string{"i-2ze5ysm1hx7o9q3mz218", "i-2ze5ysm1hx7o9q3mz219"})
	t.Logf("infos: %v", res)
	t.Log("err: ", err)
}

func TestShrink(t *testing.T) {
	cluster := types.ClusterInfo{
		RegionId:   "cn-beijing",
		Provider:   cloud.AlibabaCloud,
		AccountKey: "xxx",
	}
	err := service.Shrink(&cluster, []string{"i-2ze5ysm1hx7o9q3mz218", "i-2ze5ysm1hx7o9q3mz219"})
	t.Log("err: ", err)
}

func TestCreateExpandTask(t *testing.T) {
	req := fmt.Sprintf(`{"cluster_name":"gf.bridgx.online", "count": 1}`)
	ret, err := utils.HttpPostJsonDataT("http://10.192.219.2:9090/api/v1/cluster/expand", []byte(req), 3)
	t.Logf("Response:%v", string(ret))
	assert.Nil(t, err, "err not nil")
}

func TestCreateShrinkTask(t *testing.T) {
	req := fmt.Sprintf(`{"cluster_name":"gf.bridgx.online", "ips":[], "count": 2}`)
	ret, err := utils.HttpPostJsonDataT("http://10.192.219.2:9090/api/v1/cluster/shrink", []byte(req), 3)
	t.Logf("Response:%v", string(ret))
	assert.Nil(t, err, "err not nil")
}

func TestGetClusterCount(t *testing.T) {
	cnt, err := service.GetClusterCount(context.Background(), []string{"LTAI5tAwAMpXAQ78pePcRb6t"})
	t.Logf("get account cnt:%v", cnt)
	assert.Nil(t, err)
	assert.NotZero(t, cnt)
	cnt, err = service.GetClusterCount(context.Background(), []string{"account_not_exist"})
	assert.Nil(t, err)
	assert.Zero(t, cnt)
}

func TestGetInstanceCount(t *testing.T) {
	cnt, err := model.CountActiveInstancesByClusterName(context.Background(), nil)
	assert.Nil(t, err)
	assert.EqualValues(t, cnt, 0)
}

func TestListClustersByCond(t *testing.T) {
	cond := model.ClusterSearchCond{
		AccountKeys: nil,
		ClusterName: "",
		Provider:    "",
		Usage:       "unused",
		PageNum:     1,
		PageSize:    10,
	}

	res, total, _ := model.ListClustersByCond(context.Background(), cond)
	t.Logf("res:%v", res)
	t.Logf("total:%v", total)
}

func TestDeleteClusters(t *testing.T) {
	err := service.DeleteClusters(context.Background(), []int64{1355}, 0)
	assert.Nil(t, err)
}
