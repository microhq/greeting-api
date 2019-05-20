FROM alpine:3.2
ADD greeting-api /greeting-api
ENTRYPOINT [ "/greeting-api" ]
