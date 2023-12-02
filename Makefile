
GO_CMD = gofumpt -w . && go fmt ./... && go generate ./... && govulncheck ./... && golint ./... && go vet ./... && staticcheck ./... && go build ./... && golangci-lint run --timeout 3m --enable-all --disable lll,rowserrcheck,sqlclosecheck,wastedassign,wrapcheck,gomoddirectives,deadcode,interfacer,structcheck,varcheck,golint,ifshort,maligned,nosnakecase,exhaustivestruct,scopelint,testpackage,gochecknoglobals,paralleltest,exhaustruct,varnamelen,forbidigo,funlen,ireturn,depguard,nolintlint -e .*G114.* --out-format line-number --path-prefix `pwd` ./...

GO_FILES = $(shell find ./ -name .git -prune -o -name \*.go )
all: all_mod ${GO_FILES}
	$(GO_CMD)

all_mod:
	go mod download

tidy: 
	go mod tidy

test:
	go test ./... -v -v -race -coverprofile fmtcoverage.html
	gotestsum  --junitfile junit.xml ./...
	cd release_files ; python3 junit_to_csv.py
	cat junit.xml
	cat test_cases.csv

.PHONY: tidy proto all_mod test

