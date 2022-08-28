# gen-tool

一个mysql数据表自动生成proto、curd、gorm、api文档工具；

### 使用说明：


#### 生成mysql以及文档
./scripts/gormgen.sh 127.0.0.1:3306 root 123456 db_ad_system ts_ad_cs_white

#### 生成prto


#### 通过proto生成文档
protoc --proto_path=. \
--proto_path=./third_party \
--openapi_out=naming=proto,title=广告相关接口,description=广告相关接口:. \
./proto/db_ad_system.proto
