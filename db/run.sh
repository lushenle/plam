docker run --name postgres -d \
    -e POSTGRES_PASSWORD=mypass \
    -e POSTGRES_USER=root \
    -e POSTGRES_DB=plam \
    -e TZ=Asia/Shanghai \
    -e PGTZ=Asia/Shanghai \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v ${PWD}/data:/var/lib/postgresql/data \
    -v /etc/localtime:/etc/localtime:ro \
    -p 127.0.0.1:5432:5432 \
    postgres:16.2-bullseye
