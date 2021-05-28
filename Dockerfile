FROM ubuntu:14.04
LABEL author="michaelpeterswa"

COPY . usr/src/kochava_postback
CMD [ "bash", "kochava_postback/docker_entry.sh" ]