package main

import (
	"fmt"
	"github.com/dadadamarine/orb/encoding/mvt"
	"github.com/dadadamarine/orb/encoding/wkt"
	"io/ioutil"
	"log"
)

func main() {
	// pbf, err = mvt.Marshal(layers)
	//b, err := ioutil.ReadFile("/home/ean/Desktop/code/exercise/golang/demo/pbf_demo/1687")
	b, err := ioutil.ReadFile("/home/ean/Desktop/code/exercise/golang/demo/pbf_demo/1687_postgis")
	if err != nil {
		log.Fatalln("read file error, err: ", err)
	}
	layers, err := mvt.Unmarshal(b)
	if err != nil {
		log.Fatalln("unmarshal error, err: ", err)
	}
	fc := layers.ToFeatureCollections()
	for k, v := range fc {
		fmt.Println("layer_name: ", k)
		for _, f := range v.Features {
			//fmt.Printf("feature: %T\n", i)
			wktStr := wkt.MarshalString(f.Geometry)
			fmt.Printf("%v: %s\n", f.ID, wktStr)
		}
	}
}
