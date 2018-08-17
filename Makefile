default: build

deps:
	go get golang.org/x/oauth2
	go get golang.org/x/oauth2/facebook
	go get github.com/dgrijalva/jwt-go


bin/oauth-jwt-go: src/*.go
	go build -o $@ $^

build: bin/oauth-jwt-go

runnotify: build
	-killall oauth-jwt-go
	-terminal-notifier -title "oauth-jwt-go" -message "Built and running!" -remove
	bin/oauth-jwt-go

watch:
	supervisor --no-restart-on exit -e go -i bin --exec make -- runnotify

clean:
	rm -f bin/*

test:
	go test -v ./src

run: build
	bin/oauth-jwt-go
