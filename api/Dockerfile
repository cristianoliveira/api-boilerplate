# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# Copy the local package files to the container's workspace.
ADD . /go/src/api

# dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson

# Install app
RUN go install api

# Run the outyet command by default when the container starts.
CMD sh

ENTRYPOINT $GOPATH/bin/api

# Document that the service listens on port 8080.
EXPOSE 8080
