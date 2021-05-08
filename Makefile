.PHONY: build protobuf builddocker

build:
	echo '~~~~~~~~~~~~~~~~~~'

builddocker:
	cd grape &&\
	docker build -t grape -f docker/Dockerfile .

migrate:
	cd grape &&\
	rake db:migrate

migrateredo:
	cd grape &&\
	rake db:migrate VERSION=0 &&\
	rake db:migrate

generate_models:
	cd grape &&\
	sqlboiler -c config/sqlboiler.toml --add-global-variants --add-panic-variants psql &&\
	go generate models/services_helper.go

updatedist:
	rm -rf grape/server/ui/static/*
	cp -r ../grapeui/dist/* grape/server/ui/static/