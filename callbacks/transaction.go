package callbacks

import (
	"github.com/iznauy/gorm"
)

func BeginTransaction(db *gorm.DB) {
	if tx := db.Begin(); tx.Error == nil {
		db.Statement.ConnPool = tx.Statement.ConnPool
		db.InstanceSet("gorm:started_transaction", true)
	} else {
		tx.Error = nil
	}
}

func CommitOrRollbackTransaction(db *gorm.DB) {
	if _, ok := db.InstanceGet("gorm:started_transaction"); ok {
		if db.Error == nil {
			db.Commit()
		} else {
			db.Rollback()
		}
		db.Statement.ConnPool = db.ConnPool
	}
}
