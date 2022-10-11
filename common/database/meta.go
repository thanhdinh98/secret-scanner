package database

type (
	Handler interface {
		Raw(sql string, values ...interface{}) HandlerResult
		Exec(sql string, values ...interface{}) HandlerResult
		Transaction(callback TransactionCallback) error
	}

	LightHandler interface {
		Raw(sql string, values ...interface{}) HandlerResult
		Exec(sql string, values ...interface{}) HandlerResult
	}

	HandlerResult interface {
		Error() error
		Scan(data interface{}) error
	}
)
