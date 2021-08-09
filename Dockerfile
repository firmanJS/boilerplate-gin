# # ======================
# #  GO FIRST STAGE
# # ======================

# FROM golang:latest as builder
# USER ${USER}
# WORKDIR /usr/src/app
# COPY go.mod \
#   go.sum ./
# RUN go mod download
# COPY . ./
# ENV GO111MODULE="on" \
#   GOARCH="amd64" \
#   GOOS="linux" \
#   CGO_ENABLED="0"
# RUN apt-get clean \
#   && apt-get remove

# # ======================
# #  GO FINAL STAGE
# # ======================

# FROM builder
# WORKDIR /usr/src/app
# RUN apt-get update \
#   && apt-get install -y \
#   make \
#   vim \
#   build-essential
# COPY --from=builder . ./usr/src/app
# RUN make goprod
# EXPOSE 3000
# CMD ["./main"]

# Builder
FROM golang:latest as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/engine /app

CMD /app/engine