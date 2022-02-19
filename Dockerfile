FROM alpine:3

ADD gotestmetrics /
EXPOSE 8080
# ENTRYPOINT ["/gotestmetrics"]
WORKDIR /
CMD ["/gotestmetrics"]
