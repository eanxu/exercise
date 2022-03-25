import os
os.environ['NLS_LANG'] = 'AMERICAN_AMERICA.AL32UTF8'

import datetime

import geopandas as gpd
from geoalchemy2 import WKTElement, Geometry


def column(path=None, engine=None, table_name=None):
    print("path: ", path)
    if path.endswith(".shp"):
        map_data = gpd.read_file(path, encoding="utf8")
    else:
        map_data = gpd.read_file(path, encoding="utf8", layer="DLTB")
    epsg = int(map_data.crs.srs.split(':')[-1])  # 读取shp的空间参考

    # 入库
    map_data['the_geom'] = map_data['geometry'].apply(lambda x: WKTElement(x.wkt, epsg))
    map_data.pop("geometry")
    map_data.to_sql(
        name=table_name,
        con=engine,
        # if_exists='replace',  # 如果表存在，则替换原有表
        chunksize=1000,  # 设置一次入库大小，防止数据量太大卡顿
        # 指定geometry的类型,这里直接指定geometry_type='GEOMETRY'，防止MultiPolygon无法写入
        dtype={'the_geom': Geometry(  # the_geom为数据库表格geometry字段名称
            geometry_type='GEOMETRY', srid=epsg)},
        method='multi',
        index="fid"  # 数据库索引名称
    )


def main(path=None, engine=None, table_name=None):
    column(path, engine, table_name)


if __name__ == '__main__':
    # 超算
    # table_name = f"shp_DLTB_1_row"
    # path = """/home/xuyi/数据保密.gdb"""
    # engine = """postgresql://gpadmin:gpadmin@10.30.117.12:15432/vector_test?sslmode=disable"""

    # 本地
    table_name = f"adcode_test"
    path = """/home/ean/Desktop/adcode84/adcode.shp"""
    # engine = """postgresql://gpadmin:gpadmin@192.168.0.238:15432/vector_test?sslmode=disable"""
    engine = """postgresql://postgres:123456@localhost:5432/vector_test?sslmode=disable"""

    starttime = datetime.datetime.now()

    main(path, engine, table_name)

    endtime = datetime.datetime.now()
    print((endtime - starttime).seconds)
