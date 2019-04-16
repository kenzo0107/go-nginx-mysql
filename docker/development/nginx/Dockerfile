FROM nginx:alpine

RUN apk add --no-cache --update openssl
RUN mkdir /etc/nginx/ssl
COPY ssl /etc/nginx/ssl
