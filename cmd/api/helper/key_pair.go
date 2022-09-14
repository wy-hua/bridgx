package helper

import (
	"github.com/galaxy-future/BridgX/cmd/api/response"
	"github.com/galaxy-future/BridgX/internal/model"
)

func ConvertToKeyPairInfo(keyPair *model.KeyPair) response.KeyPairInfo {
	return response.KeyPairInfo{
		KeyId:       keyPair.GetIdStr(),
		KeyPairId:   keyPair.KeyPairId,
		KeyPairName: keyPair.KeyPairName,
		PublicKey:   keyPair.PublicKey,
		PrivateKey:  keyPair.PrivateKey,
		KeyType:     keyPair.KeyType,
	}
}

func ConvertToKeyPairList(keyPairs []*model.KeyPair) []response.KeyPair {
	var keyPairList = make([]response.KeyPair, 0, len(keyPairs))
	for _, keyPair := range keyPairs {
		var keyId string
		if keyPair.Id != 0 {
			keyId = keyPair.GetIdStr()
		}
		keyPairList = append(keyPairList, response.KeyPair{
			KeyId:       keyId,
			KeyPairId:   keyPair.KeyPairId,
			KeyPairName: keyPair.KeyPairName,
		})
	}
	return keyPairList
}
