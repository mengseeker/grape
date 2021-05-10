./run &
envoy -c service-envoy.yaml --service-cluster service${SERVICE_NAME}