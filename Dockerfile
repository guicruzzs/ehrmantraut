FROM debian:jessie
MAINTAINER Guilherme Cruz "guicruz.zs@gmail.com"

RUN apt-get update && apt-get install -y wget gnupg

RUN wget http://repo.mosquitto.org/debian/mosquitto-repo.gpg.key
RUN apt-key add mosquitto-repo.gpg.key

RUN cd /etc/apt/sources.list.d/ && wget http://repo.mosquitto.org/debian/mosquitto-jessie.list

RUN apt-get update && apt-get install -y mosquitto mosquitto-clients
