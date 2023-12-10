./bin/home_server_mono: 
	go build -mod vendor -o ./bin/home_server_mono ./cmd/home_server_mono

build: ./bin/home_server_mono

run: 
	go run ./cmd/home_server_mono/main.go
	# ./bin/home_server_mono

clean:
	rm -rf bin/*