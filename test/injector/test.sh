#!/bin/sh
curl -v \
  --cacert install/injector_cert.pem \
  -H "Content-Type: application/json" \
  --resolve graped.grape-system.svc:8082:127.0.0.1 \
  -X POST https://graped.grape-system.svc:8082/inject \