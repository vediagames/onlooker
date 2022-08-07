FROM golang:1.18 as builder

ARG GITHUB_TOKEN

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
COPY go.mod go.sum ./
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
RUN go env -w GOPRIVATE=github.com/vediagames/*
RUN go mod download
COPY . .

# Build app
RUN go build -a -installsuffix cgo -o server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/server .

CMD [ "./server" ]