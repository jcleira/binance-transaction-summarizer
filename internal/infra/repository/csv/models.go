package csv

import "github.com/jcleira/binance-transaction-summarizer/internal/aggregates"

// transaction represents a transaction in the CSV file.
type transaction struct {
	UserID    string
	UTCTime   string
	Account   string
	Operation string
	Coin      string
	Change    float64
	Remark    string
}

// ToDomain converts a csv transaction to a domain transaction.
func (t transaction) ToDomain() aggregates.Transaction {
	return aggregates.Transaction{
		UserID:    t.UserID,
		UTCTime:   t.UTCTime,
		Account:   t.Account,
		Operation: t.Operation,
		Coin:      t.Coin,
		Change:    t.Change,
		Remark:    t.Remark,
	}
}
