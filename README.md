## hello
- [x] hello module with `cmd/hello/main.go`, `internal/greet/(greet.go, greet_test.go)`.
- [x] `go.mod`, `go.sum`.
- [x] working `go build -o bin/hello ./cmd/hello`, `./bin/hello Jane` prints greeting with UUID.
- [x] `gofmt -l .`, `go vet ./...` pass.
- [x] `.golangci.yml` exists, `golangci-lint run ./...` pass.
- [x] `go test -race -count=1 ./...` pass.
- [x] `Taskfile.yml` exists, `check` runs fmt, vet, lint, and test.
- [x] `.gitignore` exists, merged pull request done.
- [x] transcription of dlv submitted.
