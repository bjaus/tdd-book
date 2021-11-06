test:
	@echo GOLANG
	@find . -name go.mod -execdir go test ./... \;
	@echo
	@echo JAVASCRIPT
	@node js/test_money.js
	@echo
	@echo PYTHON
	@python py/test_money.py
