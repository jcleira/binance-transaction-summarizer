package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/jcleira/binance-transaction-summarizer/internal/aggregates"
	"github.com/jcleira/binance-transaction-summarizer/internal/infra/repository/csv"
)

var parseCsvCmd = &cobra.Command{
	Use:   "parse-csv [file path]",
	Short: "Parse a Binance CSV file into a collection of transactions.",
	Args:  cobra.ExactArgs(1),
	Run:   Parse,
}

func Parse(cmd *cobra.Command, args []string) {
	csvRepository := csv.NewRepository(args)

	transactions, err := csvRepository.GetTransactions()
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	timestamps := []string{}
	transactionGroups := map[string]aggregates.Transactions{}

	for _, transaction := range transactions {
		if _, ok := transactionGroups[transaction.UTCTime]; !ok {
			transactionGroups[transaction.UTCTime] = aggregates.Transactions{}
			timestamps = append(timestamps, transaction.UTCTime)
		}

		transactionGroups[transaction.UTCTime] = append(
			transactionGroups[transaction.UTCTime], transaction)
	}

	for _, timestamp := range timestamps {
		if err := processTransactionGroup(transactionGroups[timestamp]); err != nil {
			/* TODO show this error message
			cmd.PrintErrln(
				fmt.Sprintf("error processing transaction group '%s': %s", timestamp, err),
			)
			*/
		}
	}

	for _, currency := range currencies {
		cmd.Println(currency)
	}
}

var (
	ErrNotSupportedTransactionGroup = errors.New("not supported transaction group")

	// currencies is a map of currencies with their respective amounts.
	// TODO this should not be a global variable.
	currencies = map[string]aggregates.Currency{}
)

func processTransactionGroup(transactions []aggregates.Transaction) error {
	// TODO use OperationTypes
	transactionTypes := []string{}

	for _, transaction := range transactions {
		transactionTypes = append(transactionTypes, transaction.Operation)
	}

	for _, transactionType := range transactionTypes {
		switch transactionType {
		case "Buy":
		case "Sell":
		case "Fee":
			// Do nothing, we currently support these transaction types.

		default:
			return ErrNotSupportedTransactionGroup
		}
	}

	for _, transaction := range transactions {
		currency := aggregates.Currency{}
		if _, ok := currencies[transaction.Coin]; !ok {
			currency = aggregates.Currency{
				Coin:   transaction.Coin,
				Amount: 0,
			}
		} else {
			currency = currencies[transaction.Coin]
		}

		// All operation can be added as they have symbol.
		currency.Amount += transaction.Change

		currencies[transaction.Coin] = currency
	}

	return nil
}
