FROM golang:1.10.0

RUN mkdir -p /g/src/washi
WORKDIR /go/src/washi

RUN go get -u github.com/ivpusic/rerun

COPY go.mod go.sum ./
COPY main.go ./
COPY cmd cmd
COPY api api
RUN go get -u golang.org/x/vgo && \
    CC=gcc vgo install . 

CMD [ "washi", "serve" ]
