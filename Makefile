build:
	go build -o golire.exe ./cmd
release:
	go build -v -a -ldflags '-s -w' \
                   -gcflags="all=-trimpath=${PWD}" \
                   -asmflags="all=-trimpath=${PWD}" \
                   -o golire.exe ./cmd