FROM golang:latest as builder

WORKDIR /app
COPY app/*.go . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vote-app-base vote-app-base.go

FROM scratch

COPY --from=builder /app/vote-app-base /

VOLUME /vote_data/

CMD ["./vote-app-base"]
