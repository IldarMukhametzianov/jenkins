FROM ubuntu:20.04

WORKDIR /app

COPY app/netinfo .

EXPOSE 8080

CMD ["./netinfo"]
