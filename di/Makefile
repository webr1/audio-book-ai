wire-build:
	@echo "Running wire in the di folder..."
	@wiregenx --root ../src --out ../di/wire/provider.go
	@cd wire/ && wire && mv wire_gen.go ../../cmd/app/container.go
