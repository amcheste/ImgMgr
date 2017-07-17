sources = main.go \
    handlers.go \
    routes.go \
    router.go \
    logger.go

program = imgMgr

$(program): $(sources)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(program) $(sources)

clean:
	@rm -f $(program)
