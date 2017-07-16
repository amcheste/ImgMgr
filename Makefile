program = imgMgr
container_name = img-mgr
test_container = img-mgr-test
build_container = img-mgr-build

.PHONY: build clean docker acceptance_tests

all: build docker acceptance_tests clean
build:
	docker build -t $(build_container) -f build/Dockerfile .
	docker run --name $(build_container) $(build_container)
	docker cp $(build_container):/root/scratch/src/ImgMgr/imgMgr .
	docker rm $(build_container)
	docker rmi $(build_container)

docker:
	docker build --pull -t $(container_name) .

acceptance_tests:
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
	rm $(program)
	docker rmi $(test_container)
