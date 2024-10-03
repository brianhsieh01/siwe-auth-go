docker run --name larryx-db \
    -e POSTGRES_DB=larryx-db \
    -e POSTGRES_USER=user \
    -e POSTGRES_PASSWORD=password \
    -p 5432:5432 \
    -v pgdata:/var/lib/postgresql/data \
    -d postgres:16
