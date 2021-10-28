GOLANG_LINT_COMMAND := $(shell { command -v golangci-lint; } 2>/dev/null)
PRECOMMIT_COMMAND := $(shell { command -v pre-commit; } 2>/dev/null)

setup: pre-commit setup-test
	@go get github.com/codegangsta/gin
	@go get github.com/swaggo/swag/cmd/swag@v1.6.7
	@go mod tidy
	@go get .

run:
	 @DEBUG=True gin -b gitlab-lint -a 8888 -p 3000 -i

collector:
	@DEBUG=True go run collector/collector.go

run-docker:
	@docker-compose up -d
	@${MAKE} run

docker-stop:
	@docker-compose stop

docker-down:
	@docker-compose down

clean:
	@find . -name "*.swp" -delete

swagger:
	@swag init

test:
	go run github.com/onsi/ginkgo/ginkgo@v1.16.4 -r .

lint:
ifndef GOLANG_LINT_COMMAND
	@echo "Command golangci-lint not found"
	@echo "Please, run the following command as sudo to install it"
	@echo "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.33.1"
	@exit 1
endif
	@golangci-lint run --out-format=github-actions

setup-test:
	@go install github.com/onsi/ginkgo/ginkgo@v1.16.4

pre-commit:
ifndef PRECOMMIT_COMMAND
	@echo "\nCommand 'pre-commit' not found!\n"
	@echo "Please, run the following command to install it:"
	@echo "\nMacOSX:"
	@echo "brew install pre-commit"
	@echo "\nGNU/Linux:"
	@echo "aptitude install pre-commit"
	@echo "\nMore info, take a look at: https://pre-commit.com/index.html#install\n\n"
	@exit 1
endif
	@pre-commit install --install-hooks
	@pre-commit install --hook-type pre-push


.PHONY: setup run collector run-docker docker-stop docker-down lint clean pre-commit setup-test
