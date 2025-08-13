build:
	sh scripts/build.sh

clean:
	sh scripts/clean.sh

coverage:
	sh scripts/coverage.sh

run:
	make clean
	sh scripts/run.sh

test:
	sh scripts/test.sh