
# RUN REDIS
docker-compose -f redis/docker-compose.yml -p redis up -d
docker-compose -f minio/docker-compose.yml -p minio up -d
