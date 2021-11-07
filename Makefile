PY ?= python3
VENV = venv
PIP = $(VENV)/bin/pip

test: go js py

go:
	@echo GOLANG
	@find . -name go.mod -execdir go test -v -shuffle on ./... \;
	@echo

js:
	@echo JAVASCRIPT
	@node js/test_money.js
	@echo

py: $(PIP)
	@echo PYTHON
	@pytest -v --random-order
	@echo

$(PIP): requirements.txt
	@$(PY) -m venv $(VENV)
	@$(PIP) install --upgrade pip wheel
	@$(PIP) install -r requirements.txt

clean:
	@rm -rf $(VENV)

.PHONY: go js py clean
