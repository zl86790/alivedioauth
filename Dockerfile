FROM golang:alpine as golang   

RUN mkdir -p /root/vedioAuth
RUN mkdir -p /root/vedioAuth/output
COPY ./ /root/vedioAuth

WORKDIR /root/vedioAuth/output
RUN go build /root/vedioAuth/main.go


FROM alpine as alpine

RUN mkdir -p /root/vedioAuth
COPY --from=golang --chown=root:root /root/vedioAuth/output /root/vedioAuth

CMD /root/vedioAuth/main

