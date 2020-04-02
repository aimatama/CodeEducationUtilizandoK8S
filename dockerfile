FROM golang:1.8-alpine as builder

RUN mkdir /src

RUN cd src

RUN mkdir /app 

ADD /src/app/. /src/app/

WORKDIR /src/app 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o goWebServer .

FROM scratch

COPY --from=builder /src/app/goWebServer /app/

WORKDIR /app

CMD ["./goWebServer"]