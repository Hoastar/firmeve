package contract

import "github.com/jinzhu/gorm"

type (
	GORMQueryMagic interface {
		Query(db *gorm.DB) *gorm.DB
	}
)
