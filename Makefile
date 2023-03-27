OBJ_NAME = 
LDFLAGS = 
install:
	$(eval OBJ_NAME += chess-engine)
	$(eval LDFLAGS += "-w -s")
	go mod tidy
	cd ./frontend/ && npm install
	cd ./cmd/api; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
run:
	$(eval OBJ_NAME += chess-engine)
	./cmd/bin/$(OBJ_NAME)	
dev:
	cd ./frontend/; npm run dev
doc:
	cd ./cmd/; godoc -http=:6060
