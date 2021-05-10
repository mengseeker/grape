FROM ubuntu

WORKDIR /server

RUN apt update && apt install -y net-tools inetutils-ping curl

COPY envoy /usr/local/bin/envoy
COPY run .