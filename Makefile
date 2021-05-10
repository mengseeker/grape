.PHONY: build protobuf builddocker

build:
	echo '~~~~~~~~~~~~~~~~~~'

updatedist:
	rm -rf grape/server/ui/static/*
	cp -r ../grapeui/dist/* grape/server/ui/static/

build-logtrans:
	go build && docker build -t logtrans .

build-envoybase:
	cd test/envoybase &&\
	docker build -t envoybase .

# extauth
build-extauth:
	docker build -t extauth -f docker/Dockerfile.extauth .

dockercompose-extauth:
	cd test/extauth && docker-compose build && docker-compose up