FROM golang:latest

WORKDIR /scorer
COPY . .
RUN make tidy
RUN make scorer
ENTRYPOINT ["bin/selection_scorer", "config/scorer.json"]
