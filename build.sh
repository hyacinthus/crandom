set -e
dep ensure
go test
go build -o crandom

