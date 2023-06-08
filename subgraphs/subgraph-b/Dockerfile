FROM golang:1.18.0 as go-build
WORKDIR /workdir/
COPY . /workdir/
RUN go get && go build
RUN ls -la

CMD [ "/workdir/subgraph-b" ]
