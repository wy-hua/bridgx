package service

import (
	"context"
	"errors"

	"github.com/galaxy-future/BridgX/internal/clients"
	"github.com/galaxy-future/BridgX/internal/constants"
	"github.com/galaxy-future/BridgX/internal/model"
	"github.com/galaxy-future/BridgX/internal/types"
	"github.com/galaxy-future/BridgX/pkg/cloud"

	"gorm.io/gorm"
)

func GetKeyPair(ctx context.Context, keyId int64) (*model.KeyPair, error) {
	var keyPair model.KeyPair
	err := model.Get(keyId, &keyPair)
	if err != nil {
		return nil, err
	}
	return &keyPair, nil
}

func CreateKeyPair(ctx context.Context, ak, provider, regionId, keyPairName string) error {
	var keyPair model.KeyPair
	err := model.QueryFirst(map[string]interface{}{"provider": provider, "region_id": regionId, "key_pair_name": keyPairName}, &keyPair)
	if err == nil {
		return errors.New("key_pair_name already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		p, err := getProvider(provider, ak, regionId)
		if err != nil {
			return err
		}
		keyPairResponse, err := p.CreateKeyPair(cloud.CreateKeyPairRequest{RegionId: regionId, KeyPairName: keyPairName})
		if err != nil {
			return err
		}
		keyPair = model.KeyPair{
			Provider:    provider,
			RegionId:    regionId,
			KeyPairId:   keyPairResponse.KeyPairId,
			KeyPairName: keyPairResponse.KeyPairName,
			PrivateKey:  keyPairResponse.PrivateKey,
			PublicKey:   keyPairResponse.PublicKey,
			KeyType:     constants.KeyTypeAuto,
		}
		return model.Save(&keyPair)
	}
	return err
}

func CreateKeyPairByPrivateKey(ctx context.Context, provider, regionId, keyPairId, keyPairName, privateKey string) (*model.KeyPair, error) {
	var keyPair model.KeyPair
	err := model.QueryFirst(map[string]interface{}{"provider": provider, "region_id": regionId, "key_pair_name": keyPairName}, &keyPair)
	if err == nil {
		return nil, errors.New("key_pair_name already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		keyPair = model.KeyPair{
			Provider:    provider,
			RegionId:    regionId,
			KeyPairId:   keyPairId,
			KeyPairName: keyPairName,
			PrivateKey:  privateKey,
			KeyType:     constants.KeyTypeSync,
		}
		err := model.Save(&keyPair)
		if err != nil {
			return nil, err
		}
		return &keyPair, nil
	}
	return nil, err
}

func ImportKeyPair(ctx context.Context, ak, provider, regionId, keyPairName, publicKey, privateKey string) error {
	var keyPair model.KeyPair
	err := model.QueryFirst(map[string]interface{}{"provider": provider, "region_id": regionId, "key_pair_name": keyPairName}, &keyPair)
	if err == nil {
		return errors.New("key_pair_name already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		p, err := getProvider(provider, ak, regionId)
		if err != nil {
			return err
		}
		keyPairResponse, err := p.ImportKeyPair(cloud.ImportKeyPairRequest{RegionId: regionId, KeyPairName: keyPairName, PublicKey: publicKey})
		if err != nil {
			return err
		}
		keyPair = model.KeyPair{
			Provider:    provider,
			RegionId:    regionId,
			KeyPairId:   keyPairResponse.KeyPairId,
			KeyPairName: keyPairName,
			PrivateKey:  privateKey,
			PublicKey:   publicKey,
			KeyType:     constants.KeyTypeImport,
		}
		return model.Save(&keyPair)
	}
	return err
}

func ListKeyPairs(ctx context.Context, ak, provider, regionId string, page types.PageReq) ([]*model.KeyPair, *types.PageRsp, error) {
	var keyPairs = make([]*model.KeyPair, 0)
	p, err := getProvider(provider, ak, regionId)
	if err != nil {
		return nil, nil, err
	}
	keyPairsResponse, err := p.DescribeKeyPairs(cloud.DescribeKeyPairsRequest{RegionId: regionId,
		PageNumber: page.PageNum, PageSize: page.PageSize, OlderMarker: page.Marker})
	if err != nil {
		return nil, nil, err
	}
	if len(keyPairsResponse.KeyPairs) == 0 {
		return keyPairs, &types.PageRsp{
			NextMarker: keyPairsResponse.NewMarker,
			Total:      keyPairsResponse.TotalCount,
		}, nil
	}

	keyPairsMap, err := getKeyPairsMapByProviderAndRegionId(ctx, provider, regionId)
	if err != nil {
		return nil, nil, err
	}

	for _, keyPair := range keyPairsResponse.KeyPairs {
		tempKeyPair := &model.KeyPair{
			KeyPairId:   keyPair.KeyPairId,
			KeyPairName: keyPair.KeyPairName,
		}
		dbKeyPair, ok := keyPairsMap[keyPair.KeyPairName]
		if ok {
			tempKeyPair.Id = dbKeyPair.Id
		}
		keyPairs = append(keyPairs, tempKeyPair)
	}

	return keyPairs, &types.PageRsp{
		NextMarker: keyPairsResponse.NewMarker,
		Total:      keyPairsResponse.TotalCount,
	}, nil
}

func getKeyPairsByProviderAndRegionId(ctx context.Context, provider, regionId string) ([]*model.KeyPair, error) {
	var keyPairs = make([]*model.KeyPair, 0)
	where := map[string]interface{}{"provider": provider, "region_id": regionId}
	err := model.QueryAll(where, &keyPairs, "")
	if err != nil {
		return nil, err
	}
	return keyPairs, nil
}

func getKeyPairsMapByProviderAndRegionId(ctx context.Context, provider, regionId string) (map[string]model.KeyPair, error) {
	keyPairs, err := getKeyPairsByProviderAndRegionId(ctx, provider, regionId)
	if err != nil {
		return nil, err
	}
	keyPairsMap := make(map[string]model.KeyPair, len(keyPairs))
	for _, pair := range keyPairs {
		keyPairsMap[pair.KeyPairName] = *pair
	}
	return keyPairsMap, nil
}

func GetKeyPairByKeyIds(ctx context.Context, keyIds []int64) (map[int64]model.KeyPair, error) {
	pairs := make([]model.KeyPair, 0)
	err := clients.ReadDBCli.WithContext(ctx).
		Model(model.KeyPair{}).
		Where("id in(?)", keyIds).
		Find(&pairs).
		Error
	if err != nil {
		return nil, err
	}
	pairMap := make(map[int64]model.KeyPair)
	for _, p := range pairs {
		pairMap[p.Id] = p
	}
	return pairMap, nil
}
