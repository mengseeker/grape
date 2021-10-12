
# ********************** all **************************
.PHONY: buildall
buildall:
	go build -o .build/grape grape/main.go
	go build -o .build/extauth extauth/main.go
	go build -o .build/pilot pilot/main.go

.PHONY: protobuf
protobuf:
	cd .. && protoc -I=. --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative \
		grape/api/**/**/*.proto


.PHONY: generate_src
generate_src:
	go generate ./...

.PHONY: generate_injector_tls
generate_injector_tls:
	mkdir -p .build/injector_tls && cd .build/injector_tls &&\
	sh -x ../../tools/generate_injector_tls.sh
	cp .build/injector_tls/cert.pem install/injector_cert.pem
	cp .build/injector_tls/key.pem install/injector_key.pem

.PHONY: mwebhook_cert_base64
mwebhook_cert_base64:
	sh -x tools/mwebhook_cert_base64.sh

# ********************** test **************************
dockercompose-extauth:
	cd test/extauth && docker-compose build && docker-compose up



.PHONY: updatedist
updatedist:
	rm -rf grape/server/ui/static/*
	cp -r ../grapeui/dist/* grape/server/ui/static/

.PHONY: dockerbuild-logtrans
dockerbuild-logtrans:
	docker build -f docker/Dockerfile.logtrans -t repo.nexttao.com.cn/mc/logtrans .

.PHONY: dockerpush-logtrans
dockerpush-logtrans:
	docker push repo.nexttao.com.cn/mc/logtrans

.PHONY: dockerbuild-envoybase
dockerbuild-envoybase:
	cd test/envoybase &&\
	docker build -t envoybase .

# extauth
dockerbuild-extauth:
	docker build -t extauth -f docker/Dockerfile.extauth .



