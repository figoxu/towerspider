package mydao

import (
	"figoxu/towerspider/common/db/model"
	"github.com/jinzhu/gorm"
)

type ActionLogDao struct {
	db *gorm.DB
}

func (p *ActionLogDao) Insert(actionLog *model.ActionLog) {
	p.db.Model(actionLog).Save(actionLog)
}

func ActionLog(db *gorm.DB) *ActionLogDao {
	return &ActionLogDao{
		db: db,
	}
}
