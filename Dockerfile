FROM alpine
MAINTAINER verosk
COPY ./http-redirect /bin/http-redirect
EXPOSE 80
ENTRYPOINT [ "/bin/http-redirect" ]

