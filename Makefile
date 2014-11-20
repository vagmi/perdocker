install: pull-image

docker-stop:
	sudo kill -QUIT `cat /var/run/docker.pid`

run: build run-perdocker

run-docker:
	sudo ./bin/docker -d
run-perdocker:
	./bin/perdocker

build:
	go build && mv perdocker ./bin/perdocker && chmod +x ./bin/perdocker

build-image: 
	docker --dns 8.8.8.8 build -rm -t="vagmi/perduniversal:latest" ./images/universal/

build-images-lang: build-image-ruby build-image-nodejs build-image-go build-image-python build-image-c build-image-php

build-image-java:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdjava:attach" ./images/java/
build-image-ruby:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdruby:attach" ./images/ruby/
build-image-nodejs:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdnodejs:attach" ./images/nodejs/
build-image-go:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdgo:attach" ./images/go/
build-image-python:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdpython:attach" ./images/python/
build-image-c:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdc:attach" ./images/c/
build-image-php:
	docker --dns 8.8.8.8 build -rm -t="vagmi/perdphp:attach" ./images/php/

#pull-image: 
	#docker pull perdocker/universal
#pull-images-lang: pull-image-ruby pull-image-nodejs pull-image-go pull-image-python pull-image-c pull-image-php

#pull-image-ruby:
	#docker pull perdocker/ruby
#pull-image-nodejs:
	#docker pull perdocker/nodejs
#pull-image-go:
	#docker pull perdocker/go
#pull-image-python:	
	#docker pull perdocker/python
#pull-image-c:	
	#docker pull perdocker/c
#pull-image-php:	
	#docker pull perdocker/php

