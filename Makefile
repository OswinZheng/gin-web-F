gen:
	go get github.com/99designs/gqlgen/cmd@v0.13.0
	go get github.com/vektah/gqlparser/v2@v2.1.0
	go install github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen

run_dev:
	bash ./scripts/run.sh

build_dev:
	docker build -t oswin/jinwei-go-seed:dev .