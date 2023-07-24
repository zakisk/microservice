
# Base image to build image 
FROM golang:alpine

# set the working directory
WORKDIR /app

#copy all the project files to the /app folder
COPY . /app

# install git tool in order to fetch private repositories
RUN apk update && apk add --no-cache git

#download the necessary Go dependencies
RUN go get github.com/gorilla/handlers
RUN go get github.com/hashicorp/go-hclog
RUN go get github.com/zakisk/microservice/currency/protos/currency
RUN go get google.golang.org/grpc
RUN go get google.golang.org/grpc/codes
RUN go get google.golang.org/grpc/status

RUN go mod download

#0 run the application
CMD go run main.go