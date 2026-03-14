# Build
FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY main.go ./

# Must build without cgo because libc is unavailable in runtime image
ENV CGO_ENABLED=0
RUN go build -ldflags "-s -w" -o ipapi main.go

# Run
FROM scratch
EXPOSE 8080

COPY --from=build /src/ipapi /ipapi

WORKDIR /
ENTRYPOINT ["/ipapi"]