FROM gcr.io/distroless/static-debian12

ENV NEW_RELIC_DEV_KEY=""

EXPOSE 1990

COPY bin/app /app

CMD ["/app"]