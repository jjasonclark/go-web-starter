FROM golang:1.6-alpine

EXPOSE 3000
WORKDIR /app/
CMD ["/app/app"]

COPY dist/app /app/
