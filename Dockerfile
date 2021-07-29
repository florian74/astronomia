FROM golang:1.16

ENV GOPATH=/go
ENV GOBIN=/go/bin

WORKDIR $GOPATH/src/astronomia


# Copy the code into the container
COPY . .
# Copy and download dependency using go mod
RUN go get -d -v ./...

# Build the application
RUN go install ./...
#RUN go install github.com/florian74/astronomia/main/...


# Run the executable
COPY ./main/gorillaz/configs /go/src/astronomia/configs
RUN ["chmod", "+x", "/go/bin/manual"]
RUN ["chmod", "+x", "/go/bin/gorillaz"]


ENTRYPOINT ["gorillaz"]

#in terminal
#run docker build -t <image_name> .
#run docker push <image_name>