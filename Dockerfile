FROM golang:1.6-alpine

COPY build/app /app/
WORKDIR /app/

EXPOSE 3000

CMD ["/app/app"]
