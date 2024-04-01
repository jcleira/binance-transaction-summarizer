package aggregates

type OperationType string

const (
	Deposit                 OperationType = "DEPOSIT"
	TransactionRelated      OperationType = "TRANSACTION_RELATED"
	Withdraw                OperationType = "WITHDRAW"
	BinanceConvert          OperationType = "BINANCE_CONVERT"
	Buy                     OperationType = "BUY"
	Sell                    OperationType = "SELL"
	Fee                     OperationType = "FEE"
	MarginLoan              OperationType = "MARGIN_LOAN"
	MarginRepayment         OperationType = "MARGIN_REPAYMENT"
	IsolatedMarginLoan      OperationType = "ISOLATED_MARGIN_LOAN"
	IsolatedMarginRepayment OperationType = "ISOLATED_MARGIN_REPAYMENT"
	TransferInternal        OperationType = "TRANSFER_INTERNAL"
)

type BuyOperation struct {
	Time         string
	OriginCoin   string
	TargetCoin   string
	OriginAmount float64
	TargetAmount float64
	FeeAmount    float64
}
