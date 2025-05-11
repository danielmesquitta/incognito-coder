.PHONY: run
run:
	@wails dev -tags webkit2_41

.PHONY: build
build:
	@wails build -tags webkit2_41
