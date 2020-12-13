FROM alpine:3.9
RUN apk add ca-certificates

COPY main /app/my-app

CMD ["/app/my-app"]