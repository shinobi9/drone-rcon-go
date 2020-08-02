FROM plugins/base:linux-amd64

LABEL maintainer="shinobi9 <shinobi9c@outlook.com>" \
  org.label-schema.name="Drone RCON Support" \
  org.label-schema.vendor="shinobi9" \
  org.label-schema.schema-version="0.0.1"


COPY release/linux/amd64/drone-rcon /bin/

ENTRYPOINT ["/bin/drone-rcon"]