test: go js py

go:
	@echo GOLANG
	@find . -name go.mod -execdir go test -v -shuffle on ./... \;
	@echo

js:
	@echo JAVASCRIPT
	@node js/test_money.js
	@echo

py:
	@echo PYTHON
	@python py/test_money.py -v
	@echo

.PHONY: go js py
