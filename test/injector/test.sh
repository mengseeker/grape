#!/bin/sh
curl -v \
  --cacert install/injector_cert.pem \
  -H "Content-Type: application/json" \
  --resolve grape-injector.grape-system.svc:8443:127.0.0.1 \
  -X POST https://grape-injector.grape-system.svc:8443/inject \
  -d @test/injector/admission_review.json \
  | jq .response.patch | sed 's/"//g' | base64 -d | jq