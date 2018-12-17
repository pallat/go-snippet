# for building parameters

> -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git rev-parse HEAD`"

# pipeline

test -z $(gofmt -l . | grep -v vendor)


go get github.com/alecthomas/gometalinter
gometalinter --install --update
test -z $(gometalinter -D ./... |grep -vE vendor)