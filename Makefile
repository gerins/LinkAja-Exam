container:
	docker-compose up -d

stop:
	docker-compose down --rmi all

test:
	go test -cover ./...

run:
	export MAIN_SERVER_HOST=localhost;\
	export MAIN_SERVER_PORT=8080;\
	export DB_DRIVER=mysql;\
	export DB_USER=root;\
	export DB_PASSWORD=admin;\
	export DB_HOST=localhost;\
	export DB_PORT=3306;\
	export DB_NAME=LinkAja;\
	go run server.go

fresh:
	export MAIN_SERVER_HOST=localhost;\
	export MAIN_SERVER_PORT=8080;\
	export DB_DRIVER=mysql;\
	export DB_USER=root;\
	export DB_PASSWORD=admin;\
	export DB_HOST=localhost;\
	export DB_PORT=3306;\
	export DB_NAME=LinkAja;\
	fresh
