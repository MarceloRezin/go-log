PACKAGE=https://github.com/MarceloRezin/go-log
VERSION=v0.0.0

test:
	@ echo "Executing unit tests"
	go test ./...

deploy:
	@ echo "Deploying"
	git add .
	git commit -m "Changes to $(VERSION)"
	git tag $(VERSION)
	git push origin $(VERSION)
	GOPROXY=proxy.golang.org go list -m $(PACKAGE)@$(VERSION)
	git push