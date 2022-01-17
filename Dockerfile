FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod . 
COPY Shadow.txt .
COPY Standard.txt .
COPY Thinkertoy.txt . 
COPY readme.md .

RUN go build -o main .
ENTRYPOINT ["/app/main"]

#./nameoffile.sh to build shell file
#  in terminal "docker build -t name-of-app ."  **first step* ./exec.sh
# docker image ls *to ckeck*  **2nd step**
#$ docker run -p 8080:8081 -it/ -tid(tty console, interactive,detach) name-of-app / **container created within this step*

# To get rid off both, 
# Docker image prune -a (y/N) //docker images check
# kill container (docker kill <name of container>) then 
# docker container prune // docker ps(to check)

#docker system prune -a //delete images

# EDITING IMAGES/CONTAINERS https://www.codegrepper.com/code-examples/shell/how+to+remove+docker+images

#We can now verify that our image exists on our machine by typing docker images:

#In order to run this newly created image, we can use the docker run command and pass in the ports 
#we want to map to and the image we wish to run

#-p 8080:8081 - This exposes our application which is running on port 8081 within our container on http://localhost:8080 on our local machine.
#-it - This flag specifies that we want to run this image in interactive mode with a tty for this container process.
#my-go-app - This is the name of the image that we want to run in a container.


#To remove all images which are not used by existing containers, use the -a flag:
#docker image prune -a

#export docker as pdf

#In order to view the list of containers running in the background you can use docker ps

# If we want to have it run permanently in the background, you can replace -it with -d to run this container in detached mode.

#In order to view the list of containers running in the background you can use docker ps which should output something like this:

#$ docker ps
#CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
#70fcc9195865        my-go-app           "/app/main"              5 seconds ago       Up 3 seconds        0.0.0.0:8080->8081/tcp   silly_swirles
#If we then wanted to kill this container, we could do so by using the docker kill command and pass in that container ID that is prints out in the terminal.