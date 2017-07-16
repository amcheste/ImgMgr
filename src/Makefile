sources = main.go \
    handlers.go \
    routes.go \
    router.go \
    logger.go

program = imgMgr

$(program): $(sources)
	go build -o $(program) $(sources)

clean:
	@rm -f $(program)
