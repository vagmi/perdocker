# Perdocker
#
# VERSION               0.0.1

FROM      ubuntu
MAINTAINER vagmi

# make sure the package repository is up to date
RUN apt-get update
RUN apt-get install -y php5

RUN groupadd perdocker
RUN useradd -g perdocker perdocker

USER perdocker
