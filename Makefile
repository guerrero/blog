JSHINT=./node_modules/jshint/bin/jshint
JS_FILES_PATH=./public/scripts/**/*.js
JSHINT_FLAGS=--reporter node_modules/jshint-stylish/stylish.js

SASSC=/usr/local/bin/sassc
SASS_FILES_PATH=./public/stylesheets/scss/**/*.scss
SASSC_FLAGS=--style compressed \
						--precision 10

GO_FILES_PATH={**/,}*.go

run: import
	@gin -p="3000"

import:
	@goimports -w $(GO_FILES_PATH)

jshint: $(JS_FILES_PATH)
	@$(JSHINT) $(JSHINT_FLAGS) $? ||:

%.css: %.scss
	@$(SASSC) $(SASSC_FLAGS) $< > ./public/stylesheets/$@ ||:

prod: main.css jshint

watch: 
	@fswatch -0 $(GO_FILES_PATH) | xargs -0 -n 1 -I {} make import
	@fswatch -0 $(SASS_FILES_PATH) | xargs -0 -n 1 -I {} make prod
	@fswatch -0 $(GO_FILES_PATH) | xargs -0 -n 1 -I {} make jshint

clean:
	@rm -f gin-bin ./public/stylesheets/main.css

.PHONY: run import jshint prod watch clean