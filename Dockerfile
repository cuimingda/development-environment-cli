FROM alpine:3.21

RUN apk update && \
    apk add --no-cache \
    nodejs-lts npm icu-data-full \
    php83 php83-session php83-tokenizer php83-dom php83-xml php83-xmlwriter \
    php83-fileinfo php83-pdo php83-pdo_sqlite php83-pcntl \
    composer jq curl go git

RUN git clone https://github.com/cuimingda/development-environment-cli.git /tmp/development-environment-cli && \
    cd /tmp/development-environment-cli && \
    sh setup.sh && \
    cd / && \
    rm -rf /tmp/development-environment-cli

# 将 entrypoint.sh 脚本复制到容器中
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

# 设置 ENTRYPOINT 来执行脚本
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

# 设置默认命令为 sh
CMD ["sh"]