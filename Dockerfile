FROM golang:latest as builder

WORKDIR /app
COPY ./app /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vote-app-base vote-app-base.go

VOLUME /vote_data/

# FROM scratch

# COPY --from=builder /app /

EXPOSE 8080

CMD ["/app/vote-app-base"]
# CMD ["/bin/bash"]