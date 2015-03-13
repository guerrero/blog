build: import
	@gin -p="3000"

import:
	@goimports -w {**/,}*.go

clean:
	@rm -rf gin-bin
