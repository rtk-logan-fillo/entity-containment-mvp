FROM golang:1.10 AS build
WORKDIR /go/src
COPY entity-containment ./entity-containment
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o entitycontainment .

FROM scratch AS runtime
COPY --from=build /go/src/entitycontainment ./
EXPOSE 8080/tcp
ENTRYPOINT ["./entitycontainment"]
