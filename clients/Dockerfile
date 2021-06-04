FROM nginx
ADD default.conf /etc/nginx/conf.d/default.conf
COPY dist /dist
COPY clients /clients
ENTRYPOINT [ "/clients" ]
