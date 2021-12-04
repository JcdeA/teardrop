#!/bin/bash

NR_OF_SERVERS=$1
TAG_NAME=$2
API_TOKEN=$3

tee /etc/nomad.d/config/telemetry.hcl >/dev/null <<EOF
telemetry {
  collection_interval = "1s"
  disable_hostname = true
  prometheus_metrics = true
  publish_allocation_metrics = true
  publish_node_metrics = true
}
EOF

hashi-up consul install \
  --version 1.9.5 \
  --local \
  --server \
  --bootstrap-expect ${NR_OF_SERVERS} \
  --client-addr 0.0.0.0 \
  --advertise-addr "{{ GetInterfaceIP \"eth1\" }}" \
  --connect \
  --retry-join "provider=digitalocean region=${REGION} tag_name=${TAG_NAME} api_token=${API_TOKEN}"

hashi-up nomad install \
  --version 1.1.0 \
  --local \
  --server \
  --bootstrap-expect ${NR_OF_SERVERS} \
  --advertise "{{ GetInterfaceIP \"eth1\" }}"

systemctl start consul
systemctl start nomad




