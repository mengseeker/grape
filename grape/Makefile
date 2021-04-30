.PHONY: build protobuf builddocker

build:
	echo '~~~~~~~~~~~~~~~~~~'

builddocker:
	docker build -t grape -f docker/Dockerfile .

migrate:
	rake db:migrate

migrateredo:
	rake db:migrate VERSION=0
	rake db:migrate

generate_models:
	sqlboiler -c config/sqlboiler.toml psql
	go generate models/services_helper.go