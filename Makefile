OBJ_NAME = 
LDFLAGS = 
install:
	$(eval OBJ_NAME += chess-engine)
	$(eval LDFLAGS += "-w -s")
	go mod tidy
	cd ./frontend/ && npm install
	cd ./chessapi/cmd/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
build:
	$(eval OBJ_NAME += chess-engine)
	$(eval LDFLAGS += "-w -s")
	cd ./chessapi/cmd/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
run:
	$(eval OBJ_NAME += chess-engine)
	./chessapi/bin/$(OBJ_NAME)	
dev:
	cd ./frontend/; npm run dev
doc:
	cd ./cmd/; godoc -http=:6060
