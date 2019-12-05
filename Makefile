.PHONY: boilerplate
boilerplate:
	mkdir day${DAY}
	cp TemplateMakefile day${DAY}/Makefile
	cp template.go day${DAY}/part1.go
	curl https://adventofcode.com/2019/day/${DAY}/input --header "Cookie: session=" > day${DAY}/input.txt
	touch day${DAY}/test_input.txt
