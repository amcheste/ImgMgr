sources = src/main.go \
	src/handlers.go \
	src/routes.go \
	src/router.go \
	src/logger.go 

program = imgMgr
container_name = img-mgr
test_container = img-mgr-test

$(program): $(soures)
	go build -o $(program) $(sources)

docker: $(program)
	docker build --pull -t $(container_name) .

acceptance_tests: $(program)
	docker build --pull -t $(test_container) -f acceptance_tests/Dockerfile .
	@echo
	@echo "Running test_index"
	./acceptance_tests/test_index
	@echo
	@echo "Running test_albums"
	./acceptance_tests/test_albums
	@echo
	@echo "Running test_health"
	./acceptance_tests/test_health
	@echo
	@echo "Running test_state"
	./acceptance_tests/test_state

clean:
	@rm -f $(program)

clean-docker:
	docker rmi img-mgr-test -f
