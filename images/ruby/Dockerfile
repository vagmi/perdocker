# Perdocker
#
# VERSION               0.0.1

FROM      ubuntu
MAINTAINER vagmi

# make sure the package repository is up to date
RUN apt-get update
RUN apt-get install -y curl
RUN gpg --keyserver hkp://keys.gnupg.net --recv-keys D39DC0E3
RUN curl -sSL https://get.rvm.io | bash -s stable --ruby

RUN groupadd perdocker
RUN useradd -g perdocker perdocker

USER perdocker
