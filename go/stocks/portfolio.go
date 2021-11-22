package stocks

import (
	"errors"
	"strings"
)

type Portfolio []Money

func (p Portfolio) Add(moneys ...Money) Portfolio {
	for _, m := range moneys {
		p = append(p, m)
	}
	return p
}

func buildErrMessage(failures []string) string {
	var b strings.Builder
	b.WriteString("[")
	for _, f := range failures {
		b.WriteString(f + ",")
	}
	b.WriteString("]")
	return b.String()
}

func (p Portfolio) Evaluate(bank Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, m := range p {
		if converted, err := bank.Convert(m, currency); err == nil {
			total += converted.amount
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}
	if len(failedConversions) == 0 {
		totalMoney := NewMoney(total, currency)
		return &totalMoney, nil
	}
	failures := buildErrMessage(failedConversions)
	return nil, errors.New("Missing exchange rate(s):" + failures)
}
