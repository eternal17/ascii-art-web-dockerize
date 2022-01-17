#!/bin/bash
docker build -t aaweb .

 docker run -p 8080:8080 -tid aaweb