
# ********************** all **************************
.PHONY: buildall
buildall:
	go build -o build/grape grape/main.go
	go build -o build/extauth extauth/main.go
	go build -o build/pilot pilot/main.go

.PHONY: protobuf
protobuf:
	cd .. && protoc -I=. --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative \
		grape/api/**/*.proto


.PHONY: generate
generate:
	go generate ./...



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



