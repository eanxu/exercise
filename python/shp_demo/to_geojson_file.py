import geopandas as gpd


def main():
    path = "/home/ean/Desktop/shp/adcode.shp"
    gd = gpd.GeoDataFrame.from_file(path, encoding="utf8")
    path = "/home/ean/Desktop/shp/adcode.geojson"
    gd.to_file(path, driver='GeoJSON')


if __name__ == '__main__':
    main()
