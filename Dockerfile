FROM ubuntu:latest

RUN mkdir /certs
RUN mkdir /config
RUN mkdir /policies

EXPOSE 443/tcp

ADD main /main

CMD /main
