# build image from base image
FROM golang:alpine

# set default working directory
WORKDIR /app

# copy all project files in /app
COPY . /app/

# download the necessary dependencies for application
RUN go mod download

# run the app
CMD go run main.go
