.PHONY: build

build:
	go build -o build/grape grape/main.go
	go build -o build/extauth extauth/main.go
	go build -o build/pilot pilot/main.go

updatedist:
	rm -rf grape/server/ui/static/*
	cp -r ../grapeui/dist/* grape/server/ui/static/

dockerbuild-logtrans:
	docker build -f docker/Dockerfile.logtrans -t repo.nexttao.com.cn/mc/logtrans .

dockerpush-logtrans:
	docker push repo.nexttao.com.cn/mc/logtrans

dockerbuild-envoybase:
	cd test/envoybase &&\
	docker build -t envoybase .

# extauth
dockerbuild-extauth:
	docker build -t extauth -f docker/Dockerfile.extauth .

dockercompose-extauth:
	cd test/extauth && docker-compose build && docker-compose up