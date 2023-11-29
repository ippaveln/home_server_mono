./bin/home_server_mono: 
	go build -mod vendor -o ./bin/home_server_mono ./cmd/home_server_mono

build: ./bin/home_server_mono

run: build
	./bin/home_server_mono


clean:
	rm -rf bin/*