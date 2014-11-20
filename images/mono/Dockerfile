# Perdocker
#
# VERSION               0.0.1

FROM      ubuntu
MAINTAINER vagmi

RUN apt-get update
RUN apt-get install -y mono-mcs

RUN groupadd perdocker
RUN useradd -g perdocker perdocker

USER perdocker
