FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /demo .

FROM gcr.io/distroless/static-debian12
COPY --from=build /demo /demo
EXPOSE 8080
ENTRYPOINT ["/demo"]
