package baidu

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/galaxy-future/BridgX/pkg/cloud"
)

var (
	b         *BaiduCloud
	instances []string
)

const (
	_vpcID           = "vpc-i21un0x7mmtz"
	_switchID        = "sbn-mgiqutgye6ui"
	_securityGroupID = "g-xy2ttwa9hqsb"
	_instance        = "i-kfXpdJ87"
)

func setup() {
	var err error
	b, err = New("xxx", "xx", "bj")
	if err != nil {
		log.Fatalln(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	fmt.Println("in main")
	code := m.Run()
	os.Exit(code)
}
func TestProviderType(t *testing.T) {
	fmt.Println(b.ProviderType())
}

func TestBatchCreate(t *testing.T) {
	disk := cloud.Disks{
		SystemDisk: cloud.DiskConf{
			Category:         "enhanced_ssd_pl1",
			Size:             40,
			PerformanceLevel: "",
		},
		DataDisk: nil,
	}
	network := cloud.Network{
		VpcId:                   _vpcID,
		SubnetId:                _switchID,
		SecurityGroup:           _securityGroupID,
		InternetChargeType:      "",
		InternetMaxBandwidthOut: 10,
		InternetIpType:          "",
	}
	charge := cloud.Charge{
		Period:     1,
		PeriodUnit: "Month",
		ChargeType: "PostPaid",
	}
	m := cloud.Params{
		InstanceType: "bcc.ic4.c1m1",
		ImageId:      "m-OWQC4wwM",
		Network:      &network,
		Zone:         "cn-bj-d",
		Region:       "BJ",
		Disks:        &disk,
		Charge:       &charge,
		Password:     "asdasd",
		Tags:         []cloud.Tag{{Key: "myTest", Value: "1"}},
		DryRun:       false,
	}
	var err error
	instances, err = b.BatchCreate(m, 1)
	if err != nil {
		t.Error("batch create", err)
	}
	fmt.Println("instances id :", instances)

}
func TestGetInstancesByTags(t *testing.T) {
	resp, err := b.GetInstancesByTags("BJ", []cloud.Tag{{Key: "myTest", Value: "1"}})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("by tag instance :", resp)
}
func TestGetInstances(t *testing.T) {
	ins, err := b.GetInstances([]string{_instance})
	if err != nil {
		t.Error("get instances:", err)
	}
	fmt.Println("instances info:", ins, "network info :", *ins[0].Network)
}
func TestGetZones(t *testing.T) {
	req := cloud.GetZonesRequest{
		RegionId: "BJ",
	}
	resp, err := b.GetZones(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("get zones response : ", resp)
}
func TestStopInstances(t *testing.T) {
	if err := b.StopInstances([]string{_instance}); err != nil {
		t.Error(err)
	}
}
func TestStartInstances(t *testing.T) {
	if err := b.StartInstances([]string{_instance}); err != nil {
		t.Error(err)
	}
}
func TestCreatVPC(t *testing.T) {
	request := cloud.CreateVpcRequest{
		RegionId:  "BJ",
		VpcName:   "myVPC",
		CidrBlock: "192.168.0.0/16",
	}
	resp, err := b.CreateVPC(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("VPC response: ", resp)
}
func TestGetVPC(t *testing.T) {
	req := cloud.GetVpcRequest{
		VpcId:    _vpcID,
		RegionId: "BJ",
	}
	resp, err := b.GetVPC(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("get vpc response :", resp)
}
func TestCreateSwitch(t *testing.T) {
	req := cloud.CreateSwitchRequest{
		RegionId:    "BJ",
		ZoneId:      "cn-bj-d",
		CidrBlock:   "192.168.0.0/24",
		VSwitchName: "mySwitch",
		VpcId:       _vpcID,
		GatewayIp:   "192.168.0.1",
	}
	resp, err := b.CreateSwitch(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("create switch response :", resp)
}
func TestGetSwitch(t *testing.T) {
	req := cloud.GetSwitchRequest{
		SwitchId: _switchID,
	}
	resp, err := b.GetSwitch(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("get switch response :", resp)
}
func TestCreateSecurityGroup(t *testing.T) {
	req := cloud.CreateSecurityGroupRequest{
		RegionId:          "BJ",
		SecurityGroupName: "mySecurityGroup",
		VpcId:             _vpcID,
		SecurityGroupType: "",
	}
	resp, err := b.CreateSecurityGroup(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("create security group response: ", resp)
}
func TestDescribeSecurityGroups(t *testing.T) {
	req := cloud.DescribeSecurityGroupsRequest{
		VpcId:    _vpcID,
		RegionId: "BJ",
	}
	resp, err := b.DescribeSecurityGroups(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("describe security groups response:", resp)
}
func TestDescribeGroupRules(t *testing.T) {
	req := cloud.DescribeGroupRulesRequest{
		RegionId:        "BJ",
		SecurityGroupId: _securityGroupID,
	}
	resp, err := b.DescribeGroupRules(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("describe group rules : ", resp)
}
func TestAddIngressSecurityGroupRule(t *testing.T) {
	req := cloud.AddSecurityGroupRuleRequest{
		RegionId:        "BJ",
		VpcId:           _vpcID,
		SecurityGroupId: _securityGroupID,
		IpProtocol:      "tcp",
		PortFrom:        1024,
		PortTo:          2048,
		CidrIp:          "192.168.0.0/24",
	}
	err := b.AddIngressSecurityGroupRule(req)
	if err != nil {
		t.Error(err)
	}
}
func TestDescribeImages(t *testing.T) {
	req := cloud.DescribeImagesRequest{
		RegionId:  "BJ",
		InsType:   "bcc.ic4.c1m1",
		ImageType: "System",
	}
	resp, err := b.DescribeImages(req)
	if err != nil {
		t.Error("dexcribe image:", err)
	}
	fmt.Println("describe images response: ", resp)
}
func TestDescribeAvailableResource(t *testing.T) {
	req := cloud.DescribeAvailableResourceRequest{
		RegionId: "BJ",
		ZoneId:   "cn-bj-d",
	}
	resp, err := b.DescribeAvailableResource(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("describe availabe resoure :", resp)
}
func TestShutDown(t *testing.T) {
	err := b.BatchDelete([]string{_instance}, "bj")
	if err != nil {
		t.Error(err)
	}
}
