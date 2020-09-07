build:
	go build -o virtcurse main.go

lint:
	goimports -w ./
