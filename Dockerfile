FROM alpine:latest 
WORKDIR /app
COPY main .
CMD ["/app/main"]