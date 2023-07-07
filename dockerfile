
### BUILD

FROM golang:1.20 as Build-Stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./*.go ./
COPY ./model/*.go ./model/
COPY ./controller/*.go ./controller/

RUN CGO_ENABLED=0 GOOS=linux go build -o /aplication

### DEPLOY

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=Build-Stage /aplication /aplication

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/aplication"]