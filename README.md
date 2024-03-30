# Binance Transaction Summarizer
Binance Transaction Summarizer is designed to transform your Binance transaction CSV exports into accountant-friendly summaries.

> NOTE: This project is a work in progress. The current version is a prototype and will fail to run as intended. Please check back for updates. 

![Binance Transaction Summarizer](https://i.imgur.com/mmXiqNW.jpeg)

It aims to streamline the process of analyzing and reporting your cryptocurrency transactions for accounting and tax purposes.

## Features
- **CSV Import**: Easily import your transaction history from Binance by uploading the CSV file.
- **Summary Reports**: Generate concise and clear summary reports highlighting your trading activity in a format understandable by accountants.
- **Customizable Summaries**: Tailor your summary reports based on specific criteria, such as date ranges, transaction types (trades, deposits, withdrawals), and traded pairs.
- **Data Security**: Your transaction data remains on your machine, so there is no need to upload sensitive information to external servers.

## Installation
To get started with the Binance Transaction Summarizer, follow these steps:

1. Clone the repository to your local machine:
```bash
git clone git@github.com:jcleira/binance-transaction-summarizer.git
```

2. Navigate to the project directory:
```
cd Binance-transaction-summarizer
```

3. Execute the Go run command to run the script:
```
go run main.go --file=./your-binance-transactions.csv
```

## Usage
To generate a summary from your Binance CSV export, follow these steps:

1. Export your transaction history from Binance as a CSV file.
2. Place the CSV file in the designated directory within the project (e.g., `./data`).
3. Run the summarizer script with the path to your CSV file:
4. The script will process your transactions and generate a summary report as console output.

## Contributing
Contributions to the Binance Transaction Summarizer are welcome! Please feel free to fork the repository, make changes, and submit a pull request.

## License
This project is licensed under the MIT License.
