package domain

import (
	"errors"
)

var (
	BalanceWasNotIncreasedErr               = errors.New("balance was not increased")
	BalanceWasNotDecreasedErr               = errors.New("balance was not decreased")
	ChangeBalanceByZeroAmountErr            = errors.New("changing balance by zero amount")
	TransactionSourceDestinationAreEqualErr = errors.New("source and destination of transaction are equal")
	TransactionNilSourceAndDestinationErr   = errors.New("source and destination of transaction are nil")
	TransactionNilSourceErr                 = errors.New("source of transaction are nil")
	TransactionNilDestinationErr            = errors.New("destination of transaction are nil")
	TransactionNilSourceOrDestinationErr    = errors.New("source or destination of transaction are nil")
	TransactionSourceDoesntExistErr         = errors.New("transaction source does not exist in database")
	ZeroAmountTransactionErr                = errors.New("transaction amount is zero")
	NegativeAmountTransactionErr            = errors.New("transaction amount is negative")
	UnknownTransactionTypeErr               = errors.New("unknown transaction type")
)
