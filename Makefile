clean:
	@rm -rfv ./build
	@rm -rfv ./server/statik.go

clean_frontend:
	@rm -rfv ./frontend/build
	@rm -rfv ./public

frontend: clean_frontend
	@cd ./frontend; \
	npm install && \
	npm run build && \
	cp -rv build ../public
	@go get -u -v github.com/rakyll/statik
	@statik -f -src ./public -p server

build: frontend
	@mkdir -p build
	@go build -v -o build/remonpi \
		./cmd/remonpi/main.go

