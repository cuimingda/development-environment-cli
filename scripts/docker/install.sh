#!/bin/sh

apk update
apk add nodejs-lts npm icu-data-full \
    php83 php83-session php83-tokenizer php83-dom php83-xml php83-xmlwriter php83-fileinfo php83-pdo php83-pdo_sqlite php83-pcntl \
    composer jq curl go git

git clone \
    "https://github.com/cuimingda/development-environment-cli.git" \
    "/tmp/development-environment-cli"

cd /tmp/development-environment-cli
mkdir -p bin
go build -o bin/dev ./cmd

ln -sf $(pwd)/bin/dev /usr/local/bin/dev

# 清理临时目录
cd /
rm -rf /tmp/development-environment-cli
