# command line for development-environment

build
```
go build -o bin/dev ./cmd
```

run web without port
```
go run ./cmd web
```

run web with port
```
go run ./cmd web --port 5173
```

## Publish
```sh
docker login
docker build --no-cache --tag cuimingda/development-environment:latest .

docker pull alpine:3.21 && \
docker build --tag cuimingda/development-environment:latest .

docker push cuimingda/development-environment:latest
```