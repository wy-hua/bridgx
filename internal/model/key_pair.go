package model

import (
	"time"

	"github.com/spf13/cast"

	"github.com/galaxy-future/BridgX/internal/logs"
	"github.com/galaxy-future/BridgX/pkg/utils"
	"gorm.io/gorm"
)

type KeyPair struct {
	Base
	Provider    string `gorm:"column:provider"`      //云厂商
	RegionId    string `gorm:"column:region_id"`     //区域ID
	KeyPairName string `gorm:"column:key_pair_name"` //秘钥对名称
	KeyPairId   string `gorm:"column:key_pair_id"`   //秘钥对ID
	PublicKey   string `gorm:"column:public_key"`    //公钥
	PrivateKey  string `gorm:"column:private_key"`   //私钥
	KeyType     string `gorm:"column:key_type"`      //秘钥类型 0:自动创建  1:导入
}

func (t *KeyPair) TableName() string {
	return "key_pair"
}

func (t *KeyPair) GetIdStr() string {
	return cast.ToString(t.Id)
}

func (r *KeyPair) BeforeSave(*gorm.DB) (err error) {
	if r.Id == 0 {
		r.Id = int64(utils.GetNextId())
	}
	now := time.Now()
	r.CreateAt = &now
	r.UpdateAt = &now
	encryptPrivateKey, err := utils.AESEncrypt(r.KeyPairName, r.PrivateKey)
	if err != nil {
		logs.Logger.Errorf("utils.AESEncrypt failed, err:%v", err)
		return err
	}
	r.PrivateKey = encryptPrivateKey
	return
}

func (r *KeyPair) AfterFind(*gorm.DB) (err error) {
	if r.PrivateKey != "" {
		decryptPrivateKey, err := utils.AESDecrypt(r.KeyPairName, r.PrivateKey)
		if err != nil {
			logs.Logger.Errorf("utils.AESEncrypt failed, err:%v", err)
			return err
		}
		r.PrivateKey = decryptPrivateKey
	}
	return
}
