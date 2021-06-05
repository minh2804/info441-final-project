FROM nginx
ADD default.conf /etc/nginx/conf.d/default.conf
COPY app /app
COPY build /build
ENTRYPOINT [ "/app" ]
