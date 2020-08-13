from golang:alpine

RUN mkdir /code
WORKDIR /code

COPY . /code
RUN go build main.go

ENTRYPOINT ["/code/main"]