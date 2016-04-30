FROM golang:alpine
ADD nexmo-srv /
ENTRYPOINT ./nexmo-srv
