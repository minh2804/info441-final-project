FROM alpine
RUN apk add --no-cache ca-certificates
COPY app /app
EXPOSE 443
ENTRYPOINT [ "/app" ]
