package aggregates

// Transaction represents a transaction in the domain.
type Transaction struct {
	UserID    string
	UTCTime   string
	Account   string
	Operation string
	Coin      string
	Change    float64
	Remark    string
}

// Transactions represents a collection of transactions.
type Transactions []Transaction
