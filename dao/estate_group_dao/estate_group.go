package estate_group_dao

import (
	"Demo_/model/enum"
	"Demo_/model/estate_group_model"
	"errors"
	"gorm.io/gorm"
)

//GORM 允许用户定义的钩子有 BeforeSave, BeforeCreate, AfterSave, AfterCreate 创建记录时将调用这些钩子方法，请参考 Hooks 中关于生命周期的详细信息
func (i *impl) BeforeCreate(tx *gorm.DB, eg estate_group_model.EstateGroup) (err error) {
	if eg.Name == "" {
		return errors.New("项目名称不可为空")
	}
	return nil
}

// 查询结束后会调用此Hook
func (i *impl) AfterFind(tx *gorm.DB) (err error) {
	return
}

// 增
func (i *impl) CreateEstateGroup(eg estate_group_model.EstateGroup) error {

	/*
		// 创建记录并更新给出的字段。
			db.Select("Name", "Age", "CreatedAt").Create(&user)

		// 创建记录并更新未给出的字段。
			db.Omit("Name", "Age", "CreatedAt").Create(&user)
	*/

	create := i.db.Model(&estate_group_model.EstateGroup{}).Create(&eg)
	return create.Error
}

// 批量插入
func (i *impl) BatchCreateEstateGroup(eg []estate_group_model.EstateGroup) error {
	/*
		// 指定插入数量为 100
			var users = []User{{name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}
			db.CreateInBatches(users, 100)
	*/
	create := i.db.Model(&estate_group_model.EstateGroup{}).Create(&eg)
	return create.Error
}

// 批量插入 限制单次插入数量 可以选择是否跳过hook
func (i *impl) LimitBatchCreateEstateGroup(eg []estate_group_model.EstateGroup, limit int, skipHooks bool) error {
	//限制新增个数
	i.db.Session(&gorm.Session{
		CreateBatchSize: limit,
		SkipHooks:       skipHooks})
	create := i.db.Model(&estate_group_model.EstateGroup{}).Create(&eg)
	return create.Error
}

// 查询第一条或最后一条数据
func (i *impl) SelectEstateGroupFirstOrLastOne(orderByFlag bool) (*estate_group_model.EstateGroup, error) {
	var (
		res *estate_group_model.EstateGroup
	)

	/*
		// SELECT * FROM users WHERE id = 10;
			db.First(&user, 10)

		// SELECT * FROM users WHERE id IN (1,2,3);
			db.Find(&users, []int{1,2,3})

		// 检查 ErrRecordNotFound 错误
			errors.Is(result.Error, gorm.ErrRecordNotFound)

		// 根据第一个字段排序
		// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
			type Language struct {
			  Code string
			  Name string
			}
			db.First(&Language{})
	*/

	if orderByFlag {
		// 根据主键(如果没有主键则根据第一个字段)`升序`获取
		// SELECT * FROM users ORDER BY id LIMIT 1;
		scan := i.db.Model(&estate_group_model.EstateGroup{}).First(res).Scan(res)
		if scan.Error != nil {
			return nil, scan.Error
		}
	} else {
		// 根据主键(如果没有主键则根据第一个字段)`降序`获取
		// SELECT * FROM users ORDER BY id DESC LIMIT 1
		scan := i.db.Last(res).Scan(res)
		if scan.Error != nil {
			return nil, scan.Error
		}
	}

	return res, nil
}

// 查询符合条件的所有数据
func (i *impl) SelectAllEstateGroupListByWhere(name string) ([]estate_group_model.EstateGroup, int64, error) {
	var (
		egList = make([]estate_group_model.EstateGroup, 0)
	)

	query := i.db.Model(&estate_group_model.EstateGroup{}).Find(&estate_group_model.EstateGroup{})

	if name != "" {
		/*
			// Struct
				// SELECT * FROM users WHERE name = "jinzhu" ORDER BY id LIMIT 1;
				db.Where(&User{Name: "jinzhu", Age: 0}).First(&user)
				// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
				db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)

			// Map - map 查询可以查询0值
				// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
				db.Where(map[string]interface{}{"name": "jinzhu", "age": 0}).Find(&users)
		*/
		query.Where("name = ?", name)
	}
	query.Not("del_status = ?", enum.StatusDeleted).Order("id asc").Scan(&egList)
	/*
		result.RowsAffected // 返回找到的记录数，相当于 `len(users)`
		result.Error        // returns error
	*/
	return egList, query.RowsAffected, query.Error
}

// 分页查询
func (i *impl) SelectEstateGroupPageList(page, pageSize int) ([]estate_group_model.EstateGroup, int64, error) {
	var (
		egList = make([]estate_group_model.EstateGroup, 0)
		total  int64
	)

	query := i.db.Model(&estate_group_model.EstateGroup{}).Select("*").Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&egList)

	return egList, total, query.Error
}

// 获取各城市的项目数量
func (i *impl) GetCountGroupByCityID() ([]estate_group_model.CityCountModel, error) {
	var (
		cityCountList = make([]estate_group_model.CityCountModel, 0)
	)

	query := i.db.Model(&estate_group_model.EstateGroup{}).Select("count,belong_city_id as city_id").
		Group("belong_city_id").Order("id").Find(&cityCountList)

	return cityCountList, query.Error
}

// FirstOrInit
func (i *impl) GetEstateGroupOrInit(name string) (estate_group_model.EstateGroup, error) {
	var res estate_group_model.EstateGroup
	queryOrInit := i.db.Model(&estate_group_model.EstateGroup{}).Where("name = ?", name).Attrs("remarks", "find_or_init").FirstOrInit(&res)
	return res, queryOrInit.Error
}

// FindInBatches
func (i *impl) GetEstateGroupBatches(name string, batchNum int) ([]estate_group_model.EstateGroup, error) {
	var res = make([]estate_group_model.EstateGroup, 0)
	query := i.db.Model(&estate_group_model.EstateGroup{}).Where("name = ?", name).FindInBatches(&res, batchNum, func(tx *gorm.DB, batch int) error {
		for _, r := range res {
			if r.Stage == 0 {

			}
			// 批量处理找到的记录
		}

		tx.Save(&res)

		//tx.RowsAffected // 本次批量操作影响的记录数

		//batch // Batch 1, 2, 3

		// 如果返回错误会终止后续批量操作
		return nil
	})
	return res, query.Error
}
