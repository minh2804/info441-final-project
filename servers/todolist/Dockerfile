FROM alpine
RUN apk add --no-cache ca-certificates
COPY app /app
EXPOSE 80
EXPOSE 443
ENTRYPOINT [ "/app" ]
