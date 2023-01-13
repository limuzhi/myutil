
docker pull elasticsearch:7.4.0

docker run --name elasticsearch -d -e ES_JAVA_OPTS="-Xms512m -Xmx512m" -e "discovery.type=single-node" -p 9200:9200 -p 9300:9300 elasticsearch:7.4.0

docker pull kibana:7.4.0 

docker inspect elasticsearch | findstr IPAddress

docker run --name kibana -e ELASTICSEARCH_URL=http://172.17.0.4:9200 -p 5601:5601 -d kibana:7.4.0 
docker run --name kibana -e ELASTICSEARCH_URL=http://192.168.40.114:9200 -p 5601:5601 -d kibana:7.4.0 