FROM alpine
RUN apk add --no-cache ca-certificates
COPY todolist /todolist
EXPOSE 443
ENTRYPOINT [ "/todolist" ]
