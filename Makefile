SNOWBALL_URL=https://github.com/snowballstem/snowball.git

.PHONY=test

english_stemmer: snowball

snowball:
	cd vendor; \
	./build_snowball.sh

test:
	go test ./...
