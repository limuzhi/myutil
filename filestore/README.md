//go--文档：
[](https://min.io/docs/minio/linux/developers/go/API.html#ComposeObject)

#windows docker 安装部署minio 

docker pull minio/minio


docker run --name minio \
-p 9000:9000 \
-p 9999:9999 \
-d --restart=always \
-e "MINIO_ROOT_USER=admin" \
-e "MINIO_ROOT_PASSWORD=admin123" \
-v D:\minio\data:/data \
-v D:\minio\config:/root/.minio \
minio/minio server /data \
--console-address '0.0.0.0:9999'


或者

docker run --name minio -p 9000:9000 -p 9999:9999 -d --restart=always -e "MINIO_ROOT_USER=admin" -e "MINIO_ROOT_PASSWORD=admin123" -v D:\minio\data:/data -v D:\minio\config:/root/.minio minio/minio server /data --console-address "0.0.0.0:9999"

localhost:9000 访问会自动跳转到 localhost:9999 ，打开登陆页面


