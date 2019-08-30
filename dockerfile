# 1st stage
FROM sevendollar/golang AS build-env
COPY index.html main.go /go/src/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o web . 

# 2nd stage
FROM busybox
COPY --from=build-env /go/src/web /go/src/index.html /
EXPOSE 8080
CMD /web
