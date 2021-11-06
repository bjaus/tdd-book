test:
	@echo GOLANG
	@find . -name go.mod -execdir go test -v ./... \;
	@echo
	@echo JAVASCRIPT
	@node js/test_money.js
	@echo
	@echo PYTHON
	@python py/test_money.py -v
