#
# Frontend
#
clean_frontend:
	@rm -rfv ./frontend/build
	@rm -rfv ./public

frontend: clean_frontend
	@cd ./frontend && \
	npm install && \
	npm run build && \
	cp -rv build ../public

#
# Backend
#
clean:
	@rm -rfv ./build
	@rm -rfv ./server/statik.go

statik:
	@echo ":: Generate bindata from statik..."
	@go get -u -v github.com/rakyll/statik
	@statik -f -src ./public -p server

build: build_linux_amd64 build_linux_armv6
build_linux_amd64: clean statik
	@mkdir -p ./build/linux_amd64
	@GOOS=linux GOARCH=amd64 \
		go build -v -o build/linux_amd64/remonpi \
		./cmd/remonpi/main.go

build_linux_armv6: clean statik
	@mkdir -p ./build/linux_armv6
	@GOOS=linux GOARCH=arm GOARM=6 \
		go build -v -o build/linux_armv6/remonpi \
		./cmd/remonpi/main.go

