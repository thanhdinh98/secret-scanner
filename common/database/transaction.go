package database

type (
	TransactionWrapper struct {
		ext               Handler
		onCommitCallbacks []TransactionOnCommitCallback
	}

	TransactionWrapperCallback  = func(txnWrapper *TransactionWrapper) TransactionCallback
	TransactionCallback         = func(Handler) error
	TransactionOnCommitCallback = func() error
)

func Transaction(callback TransactionWrapperCallback) error {
	txnWrapper := newTransactionWrapper()
	return txnWrapper.execute(callback(txnWrapper))
}

func newTransactionWrapper() *TransactionWrapper {
	return &TransactionWrapper{
		ext:               getExtension(),
		onCommitCallbacks: make([]TransactionOnCommitCallback, 0),
	}
}

func (txn *TransactionWrapper) execute(callback TransactionCallback) error {
	if err := txn.ext.Transaction(callback); err != nil {
		return err
	}

	if len(txn.onCommitCallbacks) > 0 {
		for _, onCommitCallback := range txn.onCommitCallbacks {
			if err := onCommitCallback(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (txn *TransactionWrapper) RegisterOnCommitCallback(callback TransactionOnCommitCallback) {
	txn.onCommitCallbacks = append(txn.onCommitCallbacks, callback)
}
