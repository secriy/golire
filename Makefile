build:
	go build -o golire.exe ./cmd
release:
	goreleaser --snapshot --skip-publish --rm-dist