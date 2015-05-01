FROM golang
RUN go get github.com/jcrumb/turbocharger
RUN go install github.com/jcrumb/turbocharger
ENTRYPOINT /go/bin/turbocharger
EXPOSE 8228