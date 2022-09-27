package aws

import (
	"fmt"
	"testing"

	"github.com/galaxy-future/BridgX/pkg/cloud"
)

func TestBatchCreate(t *testing.T) {
	aws, _ := New("", "", "cn-north-1")
	m := cloud.Params{
		Provider:     cloud.AwsCloud,
		InstanceType: "t2.micro",
		ImageId:      "ami-0a5e581c2158fe57d",
		Network: &cloud.Network{
			VpcId:                   "vpc-0d8c6a0bd621bf4c4",
			SubnetId:                "subnet-09fe97713f59f89ef",
			SecurityGroup:           "sg-07cdd57dd38d31672",
			InternetChargeType:      "",
			InternetMaxBandwidthOut: 10,
			InternetIpType:          "",
		},
		Zone:   "cnn1-az1",
		Region: "cn-north-1",
		Disks: &cloud.Disks{
			SystemDisk: cloud.DiskConf{
				Category:         "gp2",
				Size:             100,
				PerformanceLevel: "",
			},
			DataDisk: nil,
		},
		Charge: &cloud.Charge{
			Period:     1,
			PeriodUnit: "Month",
			ChargeType: "PostPaid",
		},
		Password:    "Ivgg87892789!",
		Tags:        nil,
		DryRun:      false,
		KeyPairId:   "key-06e93c1ff9818d34c",
		KeyPairName: "test_key_pair",
	}
	ids, err := aws.BatchCreate(m, 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ids)

}
