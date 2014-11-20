# Perdocker
#
# VERSION               0.0.1

FROM      ubuntu:14.04
MAINTAINER vagmi

# make sure the package repository is up to date
RUN apt-get update
RUN apt-get install -y wget ca-certificates build-essential

# RUBY
RUN apt-get install -y curl
RUN gpg --keyserver hkp://keys.gnupg.net --recv-keys D39DC0E3
RUN curl -sSL https://get.rvm.io | bash -s stable --ruby

# C & C++
RUN apt-get install -y gcc g++

# GO
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /usr/local/go/

RUN wget --no-verbose https://go.googlecode.com/files/go1.2.src.tar.gz
RUN tar -v -C /usr/local -xzf go1.2.src.tar.gz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1

# NODEJS
RUN apt-get install -y lamp-server^
RUN curl http://nodejs.org/dist/v0.10.24/node-v0.10.24-linux-x64.tar.gz > node-v0.10.24-linux-x64.tar.gz
RUN tar -xvzf node-v0.10.24-linux-x64.tar.gz
RUN mv node-v0.10.24-linux-x64 /usr/local/node
RUN ln -s /usr/local/node/bin/node /usr/bin/node

#PHP
RUN apt-get install -y php5

#PYTHON 
RUN apt-get install -y python

#JAVA
RUN apt-get install -y openjdk-7-jdk

#C#
RUN apt-get install -y mono-mcs

# USER
RUN groupadd perdocker
RUN useradd -g perdocker perdocker

USER perdocker
