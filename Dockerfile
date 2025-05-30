# syntax=docker/dockerfile:1
# Build Container
FROM golang:1.23-alpine AS build

WORKDIR /workspace
COPY . /workspace
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o fsh-lint .

# Runtime Container
FROM alpine:latest AS runtime

COPY --from=build /workspace/fsh-lint /fsh-lint

ENTRYPOINT ["/fsh-lint"]
