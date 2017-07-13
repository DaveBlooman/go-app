FROM golang:1.8
# Add code

RUN apt-get update -y
RUN apt-get install postgresql postgresql-client wget -y
ADD . /go/src/github.com/DaveBlooman/go-app

# Set current working directory
WORKDIR /app

RUN cd /go/src/github.com/DaveBlooman/go-app; go build -o myapp; cp myapp /app/

RUN wget -nv https://s3-eu-west-1.amazonaws.com/fundapps-packer-deps/migrate && chmod a+x migrate && mv migrate /usr/bin/

EXPOSE 8080

CMD "/go/src/github.com/DaveBlooman/go-app/run.sh"
