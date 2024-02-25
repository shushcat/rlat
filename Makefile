SNOWBALL_SHA=
.PHONY=test

english_stemmer: snowball

snowball:
	cd stemmer && \
	git clone https://github.com/snowballstem/snowball.git && \
	cd snowball && \
	make

test:
	go test ./...
