.PHONY: prepare

prepare:
	@echo "Installing Taskfile"
	@go install github.com/go-task/task/v3/cmd/task@latest
	@echo "Install Husky"
	@go install github.com/go-courier/husky/cmd/husky@latest && husky init
