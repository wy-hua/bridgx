package aws

import (
	"fmt"

	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/galaxy-future/BridgX/pkg/cloud"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cast"
)

// CreateSecurityGroup output missing field: RequestId
func (p *AwsCloud) CreateSecurityGroup(req cloud.CreateSecurityGroupRequest) (cloud.CreateSecurityGroupResponse, error) {
	input := &ec2.CreateSecurityGroupInput{
		Description: aws.String(req.SecurityGroupName),
		GroupName:   aws.String(req.SecurityGroupName),
		VpcId:       aws.String(req.VpcId),
	}

	output, err := p.ec2Client.CreateSecurityGroup(input)
	if err != nil {
		logs.Logger.Errorf("CreateSecurityGroup AwsCloud failed.err:[%v] req:[%v]", err, req)
		return cloud.CreateSecurityGroupResponse{}, err
	}
	return cloud.CreateSecurityGroupResponse{SecurityGroupId: aws.StringValue(output.GroupId)}, nil
}

// AddIngressSecurityGroupRule req:PrefixListId isn't use
func (p *AwsCloud) AddIngressSecurityGroupRule(req cloud.AddSecurityGroupRuleRequest) error {
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(req.SecurityGroupId),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(int64(req.PortFrom)),
				IpProtocol: aws.String(req.IpProtocol),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String(req.CidrIp),
					},
				},
				ToPort: aws.Int64(int64(req.PortTo)),
			},
		},
	}
	_, err := p.ec2Client.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		logs.Logger.Errorf("AddIngressSecurityGroupRule AwsCloud failed.err:[%v] req:[%v]", err, req)
		return err
	}
	return nil
}

// AddEgressSecurityGroupRule req:PrefixListId isn't use
func (p *AwsCloud) AddEgressSecurityGroupRule(req cloud.AddSecurityGroupRuleRequest) error {
	input := &ec2.AuthorizeSecurityGroupEgressInput{
		GroupId: aws.String(req.SecurityGroupId),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(int64(req.PortFrom)),
				IpProtocol: aws.String(req.IpProtocol),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String(req.CidrIp),
					},
				},
				ToPort: aws.Int64(int64(req.PortTo)),
			},
		},
	}
	_, err := p.ec2Client.AuthorizeSecurityGroupEgress(input)
	if err != nil {
		logs.Logger.Errorf("AddEgressSecurityGroupRule AwsCloud failed. err:[%v] req:[%v]", err, req)
		return err
	}
	return nil
}

// DescribeSecurityGroups output missing field: CreateAt
func (p *AwsCloud) DescribeSecurityGroups(req cloud.DescribeSecurityGroupsRequest) (cloud.DescribeSecurityGroupsResponse, error) {
	pageSize := _pageSize * 10
	var awsSecurityGroups = make([]*ec2.SecurityGroup, 0, pageSize)
	input := &ec2.DescribeSecurityGroupsInput{
		MaxResults: aws.Int64(int64(pageSize)),
	}
	if req.VpcId != "" {
		input.Filters = []*ec2.Filter{{Name: aws.String(_filterNameVpcId), Values: []*string{&req.VpcId}}}
	}
	err := p.ec2Client.DescribeSecurityGroupsPages(input, func(output *ec2.DescribeSecurityGroupsOutput, b bool) bool {
		awsSecurityGroups = append(awsSecurityGroups, output.SecurityGroups...)
		return output.NextToken != nil
	})
	if err != nil {
		logs.Logger.Errorf("DescribeSecurityGroups AwsCloud failed.err:[%v] req:[%v]", err, req)
		return cloud.DescribeSecurityGroupsResponse{}, err
	}
	if len(awsSecurityGroups) == 0 {
		logs.Logger.Warnf("DescribeSecurityGroups AwsCloud failed. req:[%v] len(awsSubnets) is zero", req)
		return cloud.DescribeSecurityGroupsResponse{}, nil
	}
	var securityGroups = make([]cloud.SecurityGroup, 0, len(awsSecurityGroups))
	for _, awsGroup := range awsSecurityGroups {
		securityGroups = append(securityGroups, buildSecurityGroup(req.RegionId, awsGroup))
	}
	return cloud.DescribeSecurityGroupsResponse{Groups: securityGroups}, nil
}

func (p *AwsCloud) describeSecurityGroups(regionId, groupId string) (cloud.SecurityGroup, error) {
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{aws.String(groupId)},
	}
	output, err := p.ec2Client.DescribeSecurityGroups(input)
	if err != nil {
		logs.Logger.Errorf("DescribeSecurityGroups AwsCloud failed. err:[%v] groupId:[%v]", err, groupId)
		return cloud.SecurityGroup{}, err
	}
	if output == nil || len(output.SecurityGroups) == 0 {
		logs.Logger.Warnf("DescribeSecurityGroups AwsCloud failed. groupId:[%v] output:[%v]", groupId, output)
		return cloud.SecurityGroup{}, nil
	}
	awsGroup := output.SecurityGroups[0]
	securityGroup := buildSecurityGroup(regionId, awsGroup)
	return securityGroup, nil
}

func buildSecurityGroup(regionId string, awsGroup *ec2.SecurityGroup) cloud.SecurityGroup {
	return cloud.SecurityGroup{
		SecurityGroupId:   aws.StringValue(awsGroup.GroupId),
		SecurityGroupName: aws.StringValue(awsGroup.GroupName),
		SecurityGroupType: "normal",
		VpcId:             aws.StringValue(awsGroup.VpcId),
		RegionId:          regionId,
		//CreateAt: "",
	}
}

// DescribeGroupRules output missing field: CreateAt
func (p *AwsCloud) DescribeGroupRules(req cloud.DescribeGroupRulesRequest) (cloud.DescribeGroupRulesResponse, error) {
	pageSize := _pageSize * 10
	var awsSecurityGroupRules = make([]*ec2.SecurityGroupRule, 0, pageSize)
	input := &ec2.DescribeSecurityGroupRulesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String(_filterNameGroupId),
				Values: []*string{&req.SecurityGroupId},
			},
		},
		MaxResults: aws.Int64(int64(pageSize)),
	}
	err := p.ec2Client.DescribeSecurityGroupRulesPages(input, func(output *ec2.DescribeSecurityGroupRulesOutput, b bool) bool {
		awsSecurityGroupRules = append(awsSecurityGroupRules, output.SecurityGroupRules...)
		return output.NextToken != nil
	})
	if err != nil {
		logs.Logger.Errorf("DescribeGroupRules AwsCloud failed.err:[%v] req:[%v]", err, req)
		return cloud.DescribeGroupRulesResponse{}, err
	}
	if len(awsSecurityGroupRules) == 0 {
		logs.Logger.Errorf("DescribeGroupRules AwsCloud failed. req:[%v] len(awsSecurityGroupRules) is zero", req)
		return cloud.DescribeGroupRulesResponse{}, nil
	}
	var rules = make([]cloud.SecurityGroupRule, 0, len(awsSecurityGroupRules))
	for _, rule := range awsSecurityGroupRules {
		var vpcId string
		if securityGroup, err := p.describeSecurityGroups(req.RegionId, req.SecurityGroupId); err == nil {
			vpcId = securityGroup.VpcId
		}
		rules = append(rules, cloud.SecurityGroupRule{
			VpcId:           vpcId,
			SecurityGroupId: aws.StringValue(rule.GroupId),
			PortFrom:        int(aws.Int64Value(rule.FromPort)),
			PortTo:          int(aws.Int64Value(rule.ToPort)),
			Protocol:        aws.StringValue(rule.IpProtocol),
			Direction:       _secGrpRuleDirection[aws.BoolValue(rule.IsEgress)],
			GroupId:         aws.StringValue(rule.GroupId),
			CidrIp:          aws.StringValue(rule.CidrIpv4),
			PrefixListId:    aws.StringValue(rule.PrefixListId),
			//CreateAt:
		})
	}
	return cloud.DescribeGroupRulesResponse{Rules: rules}, nil
}

func formatPortRange(fromPort, toPort int64) string {
	if fromPort == toPort {
		return cast.ToString(fromPort)
	}
	return fmt.Sprintf("%d-%d", fromPort, toPort)
}
