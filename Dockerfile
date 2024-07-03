FROM golang:latest

WORKDIR /space_service

COPY . .

RUN go mod tidy

COPY start.sh .

RUN chmod +x start.sh

CMD ["./start.sh"]

# EXPOSE 50051 50052 50061