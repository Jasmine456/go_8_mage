package common

import (
	"go_8_mage/week07/feture_extractor/transform"
	"time"
)

//地区
type Location struct {
	//省
	Province string
	//市
	City     string
}

// 卖家
type User struct {
	Name    string
	Age     int
	Gender  byte
	Address *Location
}

// 产品
type Product struct {
	Id          int
	Name        string
	Sales       int       // 销量
	Feedback    float32   // 好评率
	Seller      *User     //  商家
	OnShelfTime time.Time //上架时间
	Tags        []string
}

// 特征转换配置 文件
type FeatureConfig struct {
	Id             int `json:"id"`
	Path           string `json:"path"`
	Discretize     string `json:"discretize"`
	Hash           string `json:"hash"`
	DiscretizeFunc transform.Discretizer `json:"-"`
	HashFunc       transform.Transformer `json:"-"`
}

type FeatureConfigList []*FeatureConfig
