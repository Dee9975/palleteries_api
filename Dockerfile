FROM golang:1.17 as build

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
COPY . .
RUN go build -o /build/app .

FROM golang:1.17 as run

COPY --from=build /build/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]