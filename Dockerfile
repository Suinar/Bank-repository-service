FROM ubuntu:latest
LABEL authors="suinar"

ENTRYPOINT ["top", "-b"]