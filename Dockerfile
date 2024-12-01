FROM gcr.io/distroless/static-debian12

COPY bin/app /app

CMD ["/app"]