# base go image 
# This stage is for just build
# FROM golang:1.18-alpine as builder 

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# # build using standard go
# RUN CGO_ENABLED=0 go build -o brokerapp ./cmd/api

# RUN chmod +x /app/brokerapp



# build tiny docker image 
# take build from previous stage and copy in tiny docker image and run the app.
# See it's caller docker-compose.yml in project directory
# FROM alpine:latest 

# RUN mkdir /app

# COPY --from=builder /app/brokerapp /app

# # Run the application 

# CMD ["/app/brokerapp"]



# |-- Copy Object file from terminal to container 

FROM alpine:latest 

RUN mkdir /app

COPY brokerApp /app

# Run the application 

CMD ["/app/brokerApp"]