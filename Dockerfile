FROM golang:latest

WORKDIR /practice-server

ADD cmd ./cmd
ADD internal ./internal
ADD ui ./ui
ADD go.mod .
ADD Makefile .

RUN make build
RUN chmod +x ./bin/practice-server

CMD ./bin/practice-server

EXPOSE 4000/tcp

