package common

import (
	"go_8_mage/week07/feture_extractor/transform"
	"time"
)

type Location struct{
	Province string
	City string
}

type User struct {
	Name string
	Age int
	Gender byte
	Address *Location
}

type Product struct {
	Id int
	Name string
	Sales int
	Feedback float32
	Seller *User
	OnShelfTime time.Time
	Tags []string
}

type FeatureConfig struct {
	Id int
	Path string
	Discretize string
	Hash string
	DiscretizeFunc transform.Discretizer
	HashFunc transform.Transformer
}
