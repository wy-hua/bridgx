package aws

import (
	"strings"

	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/galaxy-future/BridgX/pkg/cloud"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AwsCloud struct {
	ec2Client *ec2.EC2
}

func New(ak, sk, regionId string) (*AwsCloud, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(ak, sk, ""),
		Region:      aws.String(regionId),
	})
	if err != nil {
		logs.Logger.Errorf("AwsCloud new session failed. err:[%v]", err)
		return nil, err
	}
	svc := ec2.New(sess)
	return &AwsCloud{svc}, nil
}

func (*AwsCloud) ProviderType() string {
	return cloud.AwsCloud
}

// GetRegions
func (p *AwsCloud) GetRegions() (cloud.GetRegionsResponse, error) {
	input := &ec2.DescribeRegionsInput{}
	output, err := p.ec2Client.DescribeRegions(input)
	if err != nil {
		logs.Logger.Errorf("GetRegions AwsCloud failed. err:[%v]", err)
		return cloud.GetRegionsResponse{}, err
	}
	var regions = make([]cloud.Region, 0, len(output.Regions))
	for _, region := range output.Regions {
		regions = append(regions, cloud.Region{
			RegionId:  aws.StringValue(region.RegionName),
			LocalName: _regionLocalName[aws.StringValue(region.RegionName)],
		})
	}
	return cloud.GetRegionsResponse{Regions: regions}, nil
}

//DescribeImages req:InsType isn't use
func (p *AwsCloud) DescribeImages(req cloud.DescribeImagesRequest) (cloud.DescribeImagesResponse, error) {
	instanceType, err := p.describeInstanceType(req.InsType)
	if err != nil {
		return cloud.DescribeImagesResponse{}, err
	}
	var images = make([]cloud.Image, 0, _pageSize)
	input := &ec2.DescribeImagesInput{
		Owners: append([]*string{}, aws.String(_imageType[req.ImageType])),
		Filters: append([]*ec2.Filter{}, &ec2.Filter{
			Name:   aws.String(_filterNameState),
			Values: []*string{aws.String("available")},
		}, &ec2.Filter{
			Name:   aws.String(_filterNameArchitecture),
			Values: instanceType.ProcessorInfo.SupportedArchitectures,
		}),
	}
	result, err := p.ec2Client.DescribeImages(input)
	if err != nil {
		logs.Logger.Errorf("DescribeImages AwsCloud failed. err:[%v] req:[%v]", err, req)
		return cloud.DescribeImagesResponse{}, err
	}
	for _, image := range result.Images {
		images = append(images, cloud.Image{
			Platform:  aws.StringValue(image.PlatformDetails),
			OsType:    formatOsType(aws.StringValue(image.Platform), aws.StringValue(image.PlatformDetails)),
			OsName:    aws.StringValue(image.Name),
			ImageId:   aws.StringValue(image.ImageId),
			ImageName: aws.StringValue(image.Name),
		})
	}
	return cloud.DescribeImagesResponse{Images: images}, nil
}

func formatOsType(platfrom, platformDetails string) string {
	if strings.EqualFold(platfrom, cloud.OsWindows) {
		return cloud.OsWindows
	}
	if strings.ContainsAny(strings.ToLower(platformDetails), cloud.OsLinux) {
		return cloud.OsLinux
	}
	return cloud.OsOther
}

func (p *AwsCloud) GetOrders(req cloud.GetOrdersRequest) (cloud.GetOrdersResponse, error) {
	return cloud.GetOrdersResponse{}, nil
}
