package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/jcleira/binance-transaction-summarizer/internal/aggregates"
)

type Repository struct {
	files []string
}

func NewRepository(files []string) *Repository {
	return &Repository{
		files: files,
	}
}

func (r *Repository) GetTransactions() (aggregates.Transactions, error) {
	transactions := make(aggregates.Transactions, 0)

	for _, file := range r.files {
		f, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("could not open file %s: %w", file, err)
		}

		csvReader := csv.NewReader(f)
		if _, err := csvReader.Read(); err != nil {
			return nil, fmt.Errorf("could not read header row: %w", err)
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("could not read record: %w", err)
			}

			var transaction transaction
			transaction.UserID = record[0]
			transaction.UTCTime = record[1]
			transaction.Account = record[2]
			transaction.Operation = record[3]
			transaction.Coin = record[4]
			fmt.Sscanf(record[5], "%f", &transaction.Change)
			transaction.Remark = record[6]

			transactions = append(transactions, transaction.ToDomain())
		}
	}

	return transactions, nil
}
