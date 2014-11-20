# Perdocker
#
# VERSION               0.0.1

FROM      ubuntu
MAINTAINER vagmi

RUN apt-get update
RUN apt-get install -y curl python

RUN groupadd perdocker
RUN useradd -g perdocker perdocker

USER perdocker
