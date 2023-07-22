bclear:
	rm -rf build


build:
	mkdir build build/database
	cp src/periphery/database/backend.db build/database
	cp -r src/periphery/driver/sqlite/scripts/ build/
	cp .env build/
	go build -o build/app cmd/app/main.go

cprsd:
	go build -ldflags "-s -w" -o build/app cmd/app/main.go

ccgo:
	go build -compiler gccgo -gccgoflags "-s -w" -o build/app cmd/app/main.go

win:
	 GOOS=windows GOARCH=amd64 CC=gcc CGO_ENABLED=1 go build -ldflags "-s -w" -o build/appWin.exe cmd/app/main.go

rebuild:
	make bclear
	make build