package database

import (
	"gorm.io/gorm"
)

var (
	extension *Extension = nil
)

type Extension struct {
	db *gorm.DB
}

func getExtension() Handler {
	return extension
}

func (ext *Extension) Error() error {
	return ext.db.Error
}

func (ext *Extension) Scan(data interface{}) error {
	return ext.db.Scan(data).Error
}

func (ext *Extension) Raw(sql string, values ...interface{}) HandlerResult {
	ext.db = ext.db.Raw(sql, values...)
	return ext
}

func (ext *Extension) Exec(sql string, values ...interface{}) HandlerResult {
	ext.db = ext.db.Exec(sql, values...)
	return ext
}

func (ext *Extension) Transaction(callback TransactionCallback) error {
	return ext.db.Transaction(func(tx *gorm.DB) error {
		txn := &Extension{db: tx}
		if err := callback(txn); err != nil {
			return err
		}
		return nil
	})
}
