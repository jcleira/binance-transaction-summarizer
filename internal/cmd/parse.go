package cmd

import (
	"github.com/spf13/cobra"

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

	for _, transaction := range transactions {
		cmd.Println(transaction)
	}
}
