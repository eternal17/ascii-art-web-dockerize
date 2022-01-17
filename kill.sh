#!/bin/bash
docker kill asciicontainer

docker container prune 

docker system prune -a