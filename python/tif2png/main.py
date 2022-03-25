from osgeo import gdal


def main():
    path = "/home/ean/Desktop/aaa11111.tif"
    ds = gdal.Open(path)


if __name__ == '__main__':
    main()
