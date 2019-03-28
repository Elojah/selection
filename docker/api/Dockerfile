FROM golang:latest

WORKDIR /api
COPY . .
RUN make tidy
RUN make api
ENTRYPOINT ["bin/selection_api", "config/api.json"]
