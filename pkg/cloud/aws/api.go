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
func (p *AwsCloud) CreateKeyPair(req cloud.CreateKeyPairRequest) (cloud.CreateKeyPairResponse, error) {
	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String(req.KeyPairName),
	}
	output, err := p.ec2Client.CreateKeyPair(input)
	if err != nil {
		logs.Logger.Errorf("CreateKeyPair AwsCloud failed.err:[%v] req:[%v]", err, req)
		return cloud.CreateKeyPairResponse{}, err
	}
	return cloud.CreateKeyPairResponse{
		KeyPairId:   *output.KeyPairId,
		KeyPairName: *output.KeyName,
		PrivateKey:  *output.KeyMaterial,
	}, nil
}

func (p *AwsCloud) ImportKeyPair(req cloud.ImportKeyPairRequest) (cloud.ImportKeyPairResponse, error) {
	input := &ec2.ImportKeyPairInput{
		KeyName:           aws.String(req.KeyPairName),
		PublicKeyMaterial: []byte(req.PublicKey),
	}
	output, err := p.ec2Client.ImportKeyPair(input)
	if err != nil {
		logs.Logger.Errorf("ImportKeyPair AwsCloud failed.err:[%v] req:[%v]", err, req)
		return cloud.ImportKeyPairResponse{}, err
	}
	return cloud.ImportKeyPairResponse{
		KeyPairId:   *output.KeyPairId,
		KeyPairName: *output.KeyName,
	}, nil
}

func (p *AwsCloud) DescribeKeyPairs(req cloud.DescribeKeyPairsRequest) (cloud.DescribeKeyPairsResponse, error) {
	input := &ec2.DescribeKeyPairsInput{}
	output, err := p.ec2Client.DescribeKeyPairs(input)
	if err != nil {
		logs.Logger.Errorf("DescribeKeyPairs AwsCloud failed.err:[%v] req:[%v]", err, req)
		return cloud.DescribeKeyPairsResponse{}, err
	}
	if len(output.KeyPairs) == 0 {
		return cloud.DescribeKeyPairsResponse{}, nil
	}
	totalCount := len(output.KeyPairs)
	keyPairs := make([]cloud.KeyPair, 0, totalCount)
	for _, pair := range output.KeyPairs {
		keyPairs = append(keyPairs, cloud.KeyPair{
			KeyPairId:   aws.StringValue(pair.KeyPairId),
			KeyPairName: aws.StringValue(pair.KeyName),
		})
	}
	return cloud.DescribeKeyPairsResponse{TotalCount: totalCount, KeyPairs: keyPairs}, nil
}
