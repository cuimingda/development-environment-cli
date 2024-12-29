FROM alpine:3.21

COPY ./scripts/docker/install.sh /install.sh
RUN sh /install.sh

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["sh"]