package model

import (
	"time"

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

func (r *KeyPair) BeforeCreate(*gorm.DB) (err error) {
	now := time.Now()
	r.CreateAt = &now
	r.UpdateAt = &now
	return
}

func (r *KeyPair) BeforeSave(*gorm.DB) (err error) {
	now := time.Now()
	r.UpdateAt = &now
	return
}

func (r *KeyPair) BeforeUpdate(*gorm.DB) (err error) {
	now := time.Now()
	r.UpdateAt = &now
	return
}
