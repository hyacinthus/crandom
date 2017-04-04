set -e
dep ensure
go test
go build -v -o crandom

