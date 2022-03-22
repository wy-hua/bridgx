package service

import (
	"context"
	"errors"

	"github.com/galaxy-future/BridgX/internal/constants"

	"github.com/galaxy-future/BridgX/pkg/cloud"

	"gorm.io/gorm"

	"github.com/galaxy-future/BridgX/internal/model"
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

func ListKeyPairs(ctx context.Context, provider, regionId string, pageNumber, pageSize int) ([]*model.KeyPair, int64, error) {
	var keyPairs = make([]*model.KeyPair, 0)
	where := map[string]interface{}{"provider": provider, "region_id": regionId}
	count, err := model.Query(where, pageNumber, pageSize, &keyPairs, "id desc", true)
	if err != nil {
		return nil, 0, err
	}
	return keyPairs, count, nil
}
