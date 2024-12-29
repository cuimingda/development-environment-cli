FROM alpine:3.21

COPY ./scripts/docker/install-packages.sh /install-packages.sh
RUN sh /install-packages.sh

RUN git clone "https://github.com/cuimingda/development-environment-cli.git" "/tmp/development-environment-cli" && \
    cd /tmp/development-environment-cli && \
    mkdir -p bin && \
    go build -o bin/dev ./cmd && \
    ln -sf $(pwd)/bin/dev /usr/local/bin/dev

COPY ./scripts/docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["sh"]