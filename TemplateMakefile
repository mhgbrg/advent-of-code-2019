.PHONY: part1
part1:
	ag -l . ../ | entr -s 'go run part1.go < input.txt'

.PHONY: part1-test
part1-test:
	ag -l . ../ | entr -s 'go run part1.go < test_input.txt'

part2.go: part1.go
	cp -n part1.go part2.go

.PHONY: part2
part2:
	ag -l . ../ | entr -s 'go run part2.go < input.txt'

.PHONY: part2-test
part2-test:
	ag -l . ../ | entr -s 'go run part2.go < test_input.txt'

