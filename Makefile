define get_version
$(shell cat current_version)
endef

VERSION=$(call get_version,)

run-rebase:
	git rebase -p master 2>/dev/null | grep "Your branch is up-to-date with 'origin/master'." || echo "\nPlease rebase your branch with master!"

run-tests:
	go test -timeout 10s -failfast ./...

build-package:
	dep ensure

install-dep:
	go get -u github.com/golang/dep/cmd/dep