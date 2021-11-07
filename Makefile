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
	@pytest -v --random-order
	@echo

.PHONY: go js py
