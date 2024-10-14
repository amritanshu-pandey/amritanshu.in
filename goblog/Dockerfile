FROM docker.io/library/golang:1.23

WORKDIR /src
COPY . /src/
RUN ls -lrtah
ENV GOBLOG_CONFIG_PATH=./goblog.json
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go mod download
RUN go mod tidy && go build -o /bin/ ./...

FROM scratch
COPY --from=0 /bin/goblog /bin/goblog
CMD ["/bin/goblog"]
