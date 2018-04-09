FROM golang
WORKDIR /opt
ADD main.go /opt
RUN go build /opt/main.go
CMD /opt/main
EXPOSE 8080
