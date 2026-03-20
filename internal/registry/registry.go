package registry

import (
	"github.com/nicoki2004/d2-internal/internal/destiny"
)

var Manifest *destiny.ManifestRecord

func InitManifest(path string) error {
	m, err := destiny.LoadManifestDefinition(path)
	if err != nil {
		return err
	}
	Manifest = m
	return nil
}

var items map[uint32]destiny.ManifestItem

func InitItemRegistry(inventoryMap map[uint32]destiny.ManifestItem) {
	items = inventoryMap
}

func GetItem(hash uint32) (destiny.ManifestItem, bool) {
	val, ok := items[hash]
	return val, ok
}
