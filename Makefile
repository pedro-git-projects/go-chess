OBJ_NAME = 
LDFLAGS = 
install:
	$(eval SOCKET_APP += chess-engine)
	$(eval BLOG_APP += blog-engine)
	$(eval LDFLAGS += "-w -s")
	go mod tidy
	cd ./frontend/ && npm install
	cd ./chessapi/cmd/; go build -v -ldflags $(LDFLAGS) -o $(SOCKET_APP); mv $(SOCKET_APP) ../bin; 
	cd ./blogapi/cmd/; go build -v -ldflags $(LDFLAGS) -o $(BLOG_APP); mv $(BLOG_APP) ../bin 
build:
	$(eval SOCKET_APP += chess-engine)
	$(eval BLOG_APP += blog-engine)
	$(eval LDFLAGS += "-w -s")
	cd ./chessapi/cmd/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
	cd ./blogapi/cmd/; go build -v -ldflags $(LDFLAGS) -o $(BLOG_APP); mv $(BLOG_APP) ../bin 
run:
	$(eval OBJ_NAME += chess-engine)
	./chessapi/bin/$(OBJ_NAME)	
dev:
	cd ./frontend/; npm run dev
doc:
	cd ./cmd/; godoc -http=:6060
