## dev/batch executes batch.
dev/batch:
	go run cmd/main.go

## test/lint executes linter.
test/lint:
	# コーディングスタイルの問題をチェックする
	(! gofmt -s -d `find . -name vendor -prune -type f -o -name '*.go'` | grep '^')
	# コーディングスタイルの問題をチェックする
	golint -set_exit_status `go list ./...`
	# import 文のフォーマットをチェックする
	(! goimports -l `find . -name vendor -prune -type f -o -name '*.go'` | grep 'go')
	# コンパイラが検出しないエラーをチェックする。
	go vet ./...
	# shadowed 変数をチェックする。
	go vet -vettool=$(which shadow) ./...
	# 関数のエラー戻り値をハンドリングしているかをチェックする。
	errcheck ./...
	# 未使用のグローバル変数と定数がないかチェックする。
	varcheck ./...
	# 代入した値を無視していないかチェックする
	ineffassign .
	# 不要な型変換をしていないかチェックする。
	unconvert -v ./...