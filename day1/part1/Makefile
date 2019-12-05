.PHONY: all
all: run

.PHONY: build
build:
	go build

.PHONY: run
run: build
	./part1 < ../input.txt

.PHONY: test
test: build
	./part1 < ../testInput.txt

.PHONY: watch
watch:
	ag -l . ../ | entr make run

.PHONY: watch-test
watch-test:
	ag -l . ../ | entr make test
