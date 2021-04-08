package n26aas

import "github.com/nhatthm/n26api/pkg/transaction"

// TransactionsFinderProvider provides transaction.Finder.
type TransactionsFinderProvider interface {
	TransactionsFinder() transaction.Finder
}
