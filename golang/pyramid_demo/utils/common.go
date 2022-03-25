package utils

import (
	"fmt"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/wkb"
)

func GetBoundGeometry(src []byte) (*geom.Extent, error) {
	gBound, err := wkb.DecodeBytes(src)
	if err != nil {
		return nil, fmt.Errorf("byte can not decode, err: %v", err)
	}
	ex, err := geom.NewExtentFromGeometry(gBound)
	if err != nil {
		return nil, fmt.Errorf("get extent from geometry error: %v", err)
	}
	return ex, nil
}
