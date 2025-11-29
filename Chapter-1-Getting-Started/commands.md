## Commands

```sh
go env -w GOPATH=/Volumes/Data/golib
go env -w GOBIN=/Volumes/Data/golib/bin
go env -w GO111MODULE=on
go env -w GOMODCACHE=/Volumes/Data/golib/pkg/mod
go env -w GOSUMDB=off
go env -w GONOPROXY=github.com/yourusername/*
go env -w GOPROXY=https://proxy.golang.org,direct
go env -w GOINSECURE=github.com/yourusername/*
go env -w GONOSUMDB=github.com/yourusername/*
go env -w CGO_CFLAGS="-I/Volumes/Data/golib/include"
go env -w CGO_LDFLAGS="-L/Volumes/Data/golib/lib"
go env -w PKG_CONFIG_PATH=/Volumes/Data/golib/lib/pkgconfig
go env -w LD_LIBRARY_PATH=/Volumes/Data/golib/lib
go env -w DYLD_LIBRARY_PATH=/Volumes/Data/golib/lib
go env -w PATH=$PATH:/Volumes/Data/golib/bin
go env -w GODEBUG="cgocheck=2"
go env -w GOGC=100
go env -w GOMAXPROCS=4
go env -w GOTRACEBACK=all
go env -w GODEBUG="madvdontneed=1"
go env -w GOFLAGS="-mod=mod"
go env -w GOROOT=/usr/local/go
go env -w GOOS=darwin
go env -w GOARCH=amd64
go env -w CGO_ENABLED=1
go env -w GOARM=7
go env -w GOMOD=/Volumes/Data/golib/src/github.com/yourusername/yourproject/go.mod
go env -w GOWORK=/Volumes/Data/golib/src/github.com/yourusername/yourproject/go.work
go env -w GONOPROXY=github.com/yourusername/*
go env -w GOINSECURE=github.com/yourusername/*
go env -w GONOSUMDB=github.com/yourusername/*
go env -w GOPRIVATE=github.com/yourusername/*
go env -w GIT_TERMINAL_PROMPT=1
```
## Shell setup

Run these in your shell to persist GOPATH (example uses `~/.bash_profile`).

```sh
echo 'export GOPATH=/Volumes/Data/golib\n' >> ~/.bash_profile
source ~/.bash_profile
```

## Environment configuration (using `go env -w`)

Set Go environment variables persistently. Replace `yourusername` and paths as needed.

```sh
go env -w GOPATH=/Volumes/Data/golib
```
- Explanation: Set the workspace root where Go will store packages and modules.

```sh
go env -w GOBIN=/Volumes/Data/golib/bin
```
- Explanation: Where `go install` will place binaries.

```sh
go env -w GO111MODULE=on
```
- Explanation: Enable Go modules (recommended for newer projects).

```sh
go env -w GOMODCACHE=/Volumes/Data/golib/pkg/mod
```
- Explanation: Location for the module cache.

```sh
go env -w GOSUMDB=off
```
- Explanation: Disable the Go checksum database (useful in air-gapped setups).

```sh
go env -w GOPROXY=https://proxy.golang.org,direct
```
- Explanation: Configure module proxy fallback to direct if proxy fails.

```sh
go env -w GONOPROXY=github.com/yourusername/*
go env -w GOINSECURE=github.com/yourusername/*
go env -w GONOSUMDB=github.com/yourusername/*
go env -w GOPRIVATE=github.com/yourusername/*
```
- Explanation: Treat your private repos specially (no proxy, checksum, etc.).

```sh
go env -w CGO_CFLAGS="-I/Volumes/Data/golib/include"
go env -w CGO_LDFLAGS="-L/Volumes/Data/golib/lib"
```
- Explanation: Compiler/linker flags for Cgo when using local C libraries.

```sh
go env -w PKG_CONFIG_PATH=/Volumes/Data/golib/lib/pkgconfig
go env -w LD_LIBRARY_PATH=/Volumes/Data/golib/lib
go env -w DYLD_LIBRARY_PATH=/Volumes/Data/golib/lib
```
- Explanation: Help find native libraries and pkg-config files on macOS/Linux.

```sh
go env -w PATH=$PATH:/Volumes/Data/golib/bin
```
- Explanation: Ensure installed Go binaries are on your shell `PATH`.

```sh
go env -w GODEBUG="cgocheck=2"
go env -w GODEBUG="madvdontneed=1"
```
- Explanation: Debug flags for runtime behavior and stricter Cgo checks.

```sh
go env -w GOGC=100
go env -w GOMAXPROCS=4
go env -w GOTRACEBACK=all
```
- Explanation: Tune garbage collector, CPU concurrency, and tracebacks.

```sh
go env -w GOFLAGS="-mod=mod"
```
- Explanation: Default flags passed to `go` (here forcing module mode).

```sh
go env -w GOROOT=/usr/local/go
go env -w GOOS=darwin
go env -w GOARCH=amd64
go env -w CGO_ENABLED=1
go env -w GOARM=7
```
- Explanation: Toolchain root and target platform/architecture settings.

```sh
go env -w GOMOD=/Volumes/Data/golib/src/github.com/yourusername/yourproject/go.mod
go env -w GOWORK=/Volumes/Data/golib/src/github.com/yourusername/yourproject/go.work
```
- Explanation: Explicit module/workspace files (optional; usually auto-detected).

## Git-related environment settings

```sh
go env -w GIT_TERMINAL_PROMPT=1
go env -w GIT_SSL_NO_VERIFY=1
go env -w GIT_CONFIG_NOSYSTEM=1
go env -w GIT_CONFIG_GLOBAL=/Volumes/Data/golib/src/github.com/yourusername/.gitconfig
go env -w GIT_CONFIG_LOCAL=/Volumes/Data/golib/src/github.com/yourusername/yourproject/.git/config
go env -w GIT_TRACE=1
go env -w GIT_CURL_VERBOSE=1
```
- Explanation: Various Git settings useful for debugging and custom configs.

```sh
go env -w GIT_SSH_COMMAND="ssh -i /Volumes/Data/golib/.ssh/id_rsa"
go env -w SSH_AUTH_SOCK=/Volumes/Data/golib/.ssh/ssh_auth_sock
go env -w GIT_SSH_VARIANT=ssh
go env -w GIT_ALLOW_PROTOCOL=file:http:https:ssh
```
- Explanation: Configure SSH for Git operations and allowed protocols.

## Misc Git LFS / config examples

```sh
go env -w GIT_LFS_SKIP_SMUDGE=1
go env -w GIT_LFS_FETCH_INCLUDE="*.bin"
go env -w GIT_LFS_FETCH_EXCLUDE="*.txt"
go env -w GIT_LFS_PUSH_INCLUDE="*.bin"
go env -w GIT_LFS_PUSH_EXCLUDE="*.txt"
```
- Explanation: Control how Git LFS fetches/pushes large files.

## Quick verification commands

```sh
go env
```
- Explanation: Print effective Go environment variables.

```sh
go version
```
- Explanation: Show the installed Go toolchain version.

```sh
go help
```
- Explanation: Show general help for the `go` command.

```sh
go list -m all
```
- Explanation: List all modules required by the current module.

```sh
go mod tidy
```
- Explanation: Add missing and remove unused module requirements.

```sh
go build ./...
```
- Explanation: Build all packages in the module (recursively).

```sh
go test ./...
```
- Explanation: Run tests for all packages.

```sh
go install github.com/yourusername/yourproject/cmd/yourcommand
```
- Explanation: Install the specified package binary to `GOBIN`.

```sh
go run github.com/yourusername/yourproject/cmd/yourcommand
```
- Explanation: Compile & run a package main from module path.

```sh
go get github.com/yourusername/yourproject/pkg/yourpackage
```
- Explanation: Download and add a dependency (pre-module or for updating).

```sh
go clean -modcache
```
- Explanation: Remove cached module downloads.

```sh
go doc fmt
```
- Explanation: Show documentation for the `fmt` package.

## `go tool` utilities

```sh
go tool compile -Version
go tool link -version
```
- Explanation: Show versions/info for compile/link tools.

```sh
go tool nm yourbinary
```
- Explanation: List symbols in a binary (`yourbinary` must exist).

```sh
go tool pprof yourbinary cpu.prof
```
- Explanation: Run pprof interactive profiling on a binary and profile data.

```sh
go tool vet ./...
```
- Explanation: Run `vet` analyzer to find suspicious code patterns.

```sh
go generate ./...
```
- Explanation: Run generator commands (//go:generate) across packages.

```sh
go fix ./...
```
- Explanation: Update code to use newer APIs (automated fixes).

## Module editing and inspection

```sh
go mod vendor
```
- Explanation: Copy module dependencies into `vendor/`.

```sh
go mod verify
```
- Explanation: Verify module downloads against local cache checksums.

```sh
go mod graph
```
- Explanation: Print module requirement graph.

```sh
go mod why github.com/some/dependency
```
- Explanation: Explain why a module is needed.

```sh
go mod edit -replace=old=new
go mod edit -require=github.com/some/dependency@v1.2.3
go mod edit -droprequire=github.com/some/dependency
go mod edit -go=1.18
go mod edit -module=github.com/yourusername/yourproject
go mod edit -json
go mod edit -fmt
go mod edit -print
```
- Explanation: Manipulate `go.mod` programmatically (replace, require, set go version, print JSON, etc.).

```sh
go mod edit -replace=github.com/old/dependency=github.com/new/dependency@v1.2.3
```
- Explanation: Replace a dependency with another fork or version locally.

```
```
*** End Patch
