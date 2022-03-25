package utils

import (
	"github.com/airmap/gdal"
	"github.com/airmap/gdal/ogr"
	"log"
	"strings"
)

func SetupInputSRS(inputDataSet gdal.Dataset, inputSRS ogr.SpatialReference) string {
	inputSRSWTK := inputDataSet.Projection()
	if inputSRSWTK != "" && inputDataSet.GDALGetGCPCount() != 0 {
		inputSRSWTK = inputDataSet.GDALGetGCPProjection()
	}
	if inputSRSWTK != "" {
		err := inputSRS.SetFromUserInput(inputSRSWTK)
		if err != nil {
			log.Fatal("SetFromUserInput error", err)
		}
	}
	return inputSRSWTK  // TODO 这个值也没用 为啥要返回值??
}

func SetupOutputSRS(inputSRS ogr.SpatialReference, isGeodetic bool) ogr.SpatialReference {
	outputSRS := ogr.CreateSpatialReference("")
	if !isGeodetic {
		outputSRS.FromEPSG(3857)
	} else {
		outputSRS.FromEPSG(4326)
	}
	return outputSRS
}

func ReprojectDataset(fromDataSet gdal.Dataset, fromSRS ogr.SpatialReference, toSRS ogr.SpatialReference) (toDataSet gdal.Dataset) {
	fromProj4, err := fromSRS.ToProj4()
	if err != nil {
		log.Fatal("fromSRS.ToProj4 error", err)
	}
	fromProj4Low := strings.ToLower(fromProj4)
	toProj4, err := toSRS.ToProj4()
	if err != nil {
		log.Fatal("toSRS.ToProj4 error", err)
	}
	toProj4Low := strings.ToLower(toProj4)
	if fromProj4Low != toProj4Low || fromDataSet.GDALGetGCPCount() != 0 {
		fromWKT, err := fromSRS.ToWKT()
		if err != nil {
			log.Fatal("fromSRS.ToWKT error", err)
		}
		toWKT, err := toSRS.ToWKT()
		if err != nil {
			log.Fatal("toSRS.ToWKT error", err)
		}
		toDataSet, err = fromDataSet.AutoCreateWarpedVRT(fromWKT, toWKT, gdal.GRA_NearestNeighbour)
		if err != nil {
			log.Fatal("fromDataSet.AutoCreateWarpedVRT error", err)
		}
		return
	}
	return
}