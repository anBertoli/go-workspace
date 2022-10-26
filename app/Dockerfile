FROM golang:1.19-alpine AS BUILDER

WORKDIR /workspace

COPY /go.work ./go.work
COPY ./app ./app
COPY ./my_box ./my_box
COPY ./my_unpub ./my_unpub

WORKDIR /workspace/app

RUN go build -o ./app_bin

FROM alpine as RUNNER

WORKDIR /app

COPY --from=BUILDER /workspace/app/app_bin ./app_bin

EXPOSE 8080

CMD ["./app_bin"]