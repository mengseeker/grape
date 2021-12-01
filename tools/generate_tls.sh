GOROOT=${GOROOT:-/usr/local/go}
go run ${GOROOT}/src/crypto/tls/generate_cert.go \
  --rsa-bits 2048 \
  --host $1 \
  --ca \
  --start-date "Jan 1 00:00:00 1970" \
  --duration=1000000h