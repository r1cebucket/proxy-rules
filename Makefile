update:
	go run ./cmd/gen.go
	git add .
	git commit -m update
	git push