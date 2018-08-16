FROM alpine

RUN apk update && apk add bash

COPY pipeline-utilities-linux /usr/bin/pipeline-utilities
RUN chmod +x /usr/bin/pipeline-utilities
