FROM alpine:3.21

RUN apk update && \
    apk add --no-cache \
    nodejs-lts npm icu-data-full \
    php83 php83-session php83-tokenizer php83-dom php83-xml php83-xmlwriter \
    php83-fileinfo php83-pdo php83-pdo_sqlite php83-pcntl \
    composer jq curl go && \
    rm -rf /var/cache/apk/*