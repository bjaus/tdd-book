test:
	go test go/money_test.go -v
	echo
	node js/test_money.js
	echo
	python py/test_money.py -v
