FROM apline:3.12

LABEL maintainers="0xff-dev"

RUN apk add util-linux coreutils && apk update && apk upgrade
WORKDIR /
COPY ./csi-example .
ENTRYPOINT [ "/csi-example", "-v=4"]