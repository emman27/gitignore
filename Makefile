build:
	mkdir -p bin
	for os in linux darwin windows; do \
		for arch in 386 amd64; do \
			GOOS=$$os GOARCH=$$arch go build -o bin/gitignore-$$os-$$arch . ; \
		done \
	done
