FROM alpine:3.21

COPY ./scripts/docker/install-packages.sh /install-packages.sh
RUN sh /install-packages.sh

COPY . /.development-environment-cli
RUN cd /.development-environment-cli && mkdir -p bin && go build -o bin/dev ./cmd && cp bin/dev /usr/local/bin

COPY ./scripts/docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["sh"]