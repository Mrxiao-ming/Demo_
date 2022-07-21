package estate_group_dao

import (
	"Demo_/model/estate_group_model"
	"gorm.io/gorm"
)

type Interface interface {
	BeforeCreate(tx *gorm.DB) (err error)
	CreateEstateGroup(isSkipHooks bool) error
	CreateOrUpdateEstateGroup() error
	BatchCreateEstateGroup(eg []estate_group_model.EstateGroup, isSkipHooks bool) error
	LimitBatchCreateEstateGroup(eg []estate_group_model.EstateGroup, limit int, skipHooks bool) error
	SelectEstateGroupFirstOrLastOne(orderByFlag bool) (*estate_group_model.EstateGroup, error)
	SelectEstateGroupPageList(page, pageSize int) ([]estate_group_model.EstateGroup, int64, error)
}

type impl struct {
	db *gorm.DB  // 可以带上context
	eg *estate_group_model.EstateGroup
}

var _ Interface = (*impl)(nil)

func NewInterface(mysql *gorm.DB) Interface {
	return &impl{
		db: mysql,
		eg: &estate_group_model.EstateGroup{},
	}
}
