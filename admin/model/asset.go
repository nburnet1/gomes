package model

import (
	"github.com/nburnet1/gomes/pkg/namespace"
	"github.com/nburnet1/gomes/config"
	"path"
	"time"
)

func init() {
	namespace.RegisterModels(Asset{})
	asset := Asset{AssetID: 10444}
	config.DB.First(&asset, asset.AssetID)
	// asset := Asset{
	// 	AssetID:    1,
	// 	Topic:      "Enterprise/Site/Area/Machine1",
	// 	Name:       "Machine1",
	// 	Type:       "CELL",
	// 	OEEEnabled: true,
	// 	Active:     true,
	// }

	assetNodes := []*namespace.NodeInstance{
		{
			Topic:        path.Join(asset.Topic, "Asset"),
			Value:        asset,
			EventHandler: someFunction,
		},
	}

	namespace.RegisterNodes(assetNodes...)

	namespace.RegisterEventHandlers(someFunction)
}

type Asset struct {
	AssetID    int `gorm:"primaryKey"`
	Topic      string
	Name       string
	Type       string
	OEEEnabled bool
	Active     bool
}

func someFunction(engine *namespace.NamespaceEngine, path string, oldValue, newValue interface{}, oldTimestamp time.Time) {
	newAsset := newValue.(Asset)
	config.DB.Save(&newAsset)
}
