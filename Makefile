all: linux windows darwin arm7 openbsd freebsd

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/github-stats-linux-amd64 github-stats.go

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/github-stats-windows-amd64.exe github-stats.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/github-stats-macos-amd64 github-stats.go

arm7:
	GOOS=linux GOARCH=arm GOARM=7 go build -o bin/github-stats-linux-arm7 github-stats.go

openbsd:
	GOOS=openbsd GOARCH=amd64 go build -o bin/github-stats-openbsd-amd64 github-stats.go

freebsd:
	GOOS=freebsd GOARCH=amd64 go build -o bin/github-stats-freebsd-amd64 github-stats.go

clean:
	rm -f bin/github-stats-linux-arm7 bin/github-stats-macos-amd64 bin/github-stats-windows-amd64.exe bin/github-stats-linux-amd64 bin/github-stats-openbsd-amd64 bin/github-stats-freebsd-amd64
