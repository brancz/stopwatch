compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o main .
