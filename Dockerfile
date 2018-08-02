FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Kleister API" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

EXPOSE 8080 8090
VOLUME ["/var/lib/kleister"]

ENTRYPOINT ["/usr/bin/kleister-api"]
CMD ["server"]

ENV KLEISTER_API_UPLOAD_DSN file://var/lib/kleister/

RUN apk add --no-cache ca-certificates mailcap bash

COPY dist/binaries/kleister-api-*-linux-amd64 /usr/bin/
