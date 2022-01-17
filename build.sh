#!/bin/bash
docker image build -t aaweb .

 docker container run -p 8080:8080 -detach --name asciicontainer aaweb

 docker image prune -a