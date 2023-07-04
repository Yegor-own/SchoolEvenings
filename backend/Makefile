bclear:
	rm -rf build


build:
	mkdir build build/database
	cp src/periphery/database/backend.db build/database
	cp -r src/periphery/driver/sqlite/scripts/ build/
	go build -o build/app cmd/app/main.go

rebuild:
	make bclear
	make build