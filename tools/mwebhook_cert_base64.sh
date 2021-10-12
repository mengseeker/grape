caBundle=$(base64 install/injector_cert.pem | tr -d "\n")
sed -i "s/caBundle: .*/caBundle: ${caBundle}/" install/injector-mwebhook.yaml