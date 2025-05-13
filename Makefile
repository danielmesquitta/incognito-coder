.PHONY: run
run:
	@cd cmd/app && wails dev -tags webkit2_41

.PHONY: build
build:
	@cd cmd/app && wails build -tags webkit2_41
