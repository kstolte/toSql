hello:
	echo "Hello"

build-wasm:
	echo "building dependencies, wasms, adding to asset folder"
	go build ./wasm/toSql
	GOOS=js GOARCH=wasm go build -o ./assets/toSql.wasm ./wasm/main.go
run:
	go run main.go
dev: build-wasm
	go run ./server/main.go
