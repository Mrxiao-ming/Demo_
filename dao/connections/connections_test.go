package connections

import (
	"fmt"
	"testing"
	"time"
)

type EstateGroup struct {
	Id                    int       `gorm:"id;default:null" json:"id,omitempty"`                                           // 小区id
	Name                  string    `gorm:"name;default:''" json:"name,omitempty"`                                         // 名字
	BelongProvinceId      int       `gorm:"belong_province_id;default:0" json:"belong_province_id,omitempty"`              // 所属省份id
	BelongCityId          int       `gorm:"belong_city_id;default:0" json:"belong_city_id,omitempty"`                      // 所属城市id
	BelongDistrictId      int       `gorm:"belong_district_id;default:0" json:"belong_district_id,omitempty"`              // 所属区ID
	BelongCompanyId       int       `gorm:"belong_company_id;default:0" json:"belong_company_id,omitempty"`                // 所属公司id
	Lonlat                string    `gorm:"lonlat;default:''" json:"lonlat,omitempty"`                                     // 经纬度,用逗号分割
	Address               string    `gorm:"address;default:''" json:"address,omitempty"`                                   // 项目所在地
	KefuPhone             string    `gorm:"kefu_phone;default:''" json:"kefu_phone,omitempty"`                             // 客服电话,如果有多个客服电话,请用英文的逗号分割。
	Contact               string    `gorm:"contact;default:''" json:"contact,omitempty"`                                   // 联系人
	ContactPhone          string    `gorm:"contact_phone;default:''" json:"contact_phone,omitempty"`                       // 负责人电话
	DirectorStaffId       int       `gorm:"director_staff_id;default:null" json:"director_staff_id,omitempty"`             // 负责人员工id
	Plane                 string    `gorm:"plane;default:null" json:"plane,omitempty"`                                     // 平面图数据
	PlaneBase             string    `gorm:"plane_base;default:''" json:"plane_base,omitempty"`                             // 平面图底图
	CreateBy              int       `gorm:"create_by;default:null" json:"create_by,omitempty"`                             // 创建者
	UpdateBy              int       `gorm:"update_by;default:null" json:"update_by,omitempty"`                             // 更新者
	DeleteBy              int       `gorm:"delete_by;default:null" json:"delete_by,omitempty"`                             // 删除者
	DelStatus             int       `gorm:"del_status;default:0" json:"del_status,omitempty"`                              // 0正常 1删除
	CreateTime            time.Time `gorm:"create_time;default:CURRENT_TIMESTAMP" json:"create_time,omitempty"`            // 创建时间
	UpdateTime            time.Time `gorm:"update_time;default:CURRENT_TIMESTAMP" json:"update_time,omitempty"`            // 更新时间
	AreaTotal             float64   `gorm:"area_total;default:0.00" json:"area_total,omitempty"`                           // 土地面积
	AreaConstruction      float64   `gorm:"area_construction;default:0.00" json:"area_construction,omitempty"`             // 建筑面积
	PlanGraphUrl          string    `gorm:"plan_graph_url;default:''" json:"plan_graph_url,omitempty"`                     // 平面图地址
	AfforestedArea        float64   `gorm:"afforested_area;default:0.00" json:"afforested_area,omitempty"`                 // 绿化面积
	AfforestedRate        float64   `gorm:"afforested_rate;default:0.00" json:"afforested_rate,omitempty"`                 // 绿化率
	ChargeArea            float64   `gorm:"charge_area;default:0.00" json:"charge_area,omitempty"`                         // 计费面积
	PlotRate              float64   `gorm:"plot_rate;default:0.00" json:"plot_rate,omitempty"`                             // 容积率
	EstablishTime         time.Time `gorm:"establish_time;default:null" json:"establish_time,omitempty"`                   // 项目成立日期
	Stage                 int       `gorm:"stage;default:0" json:"stage,omitempty"`                                        // 项目阶段
	Format                string    `gorm:"format;default:''" json:"format,omitempty"`                                     // 项目业态
	FormatName            string    `gorm:"format_name;default:''" json:"format_name,omitempty"`                           // 项目业态名称
	TakeoverTime          time.Time `gorm:"takeover_time;default:null" json:"takeover_time,omitempty"`                     // 物业接管时间
	BusinessType          int       `gorm:"business_type;default:1" json:"business_type,omitempty"`                        // 业务类型
	ContractType          int       `gorm:"contract_type;default:1" json:"contract_type,omitempty"`                        // 合同类型
	ContractName          string    `gorm:"contract_name;default:''" json:"contract_name,omitempty"`                       // 合同名称
	ContractNo            string    `gorm:"contract_no;default:''" json:"contract_no,omitempty"`                           // 合同编号
	ContractStartTime     time.Time `gorm:"contract_start_time;default:null" json:"contract_start_time,omitempty"`         // 合同开始时间
	ContractEndTime       time.Time `gorm:"contract_end_time;default:null" json:"contract_end_time,omitempty"`             // 合同结束时间
	SignTime              time.Time `gorm:"sign_time;default:null" json:"sign_time,omitempty"`                             // 签约时间
	PartyA                string    `gorm:"party_a;default:''" json:"party_a,omitempty"`                                   // 甲方名称
	PartyB                string    `gorm:"party_b;default:''" json:"party_b,omitempty"`                                   // 乙方名称
	ContractAmount        float64   `gorm:"contract_amount;default:0.00" json:"contract_amount,omitempty"`                 // 合同金额
	CommissionRate        float64   `gorm:"commission_rate;default:0.00" json:"commission_rate,omitempty"`                 // 佣金比例
	Remarks               string    `gorm:"remarks;default:''" json:"remarks,omitempty"`                                   // 备注
	ChargeRate            float64   `gorm:"charge_rate;default:0" json:"charge_rate,omitempty"`                            // 收费率
	PraiseRate            float64   `gorm:"praise_rate;default:0" json:"praise_rate,omitempty"`                            // 好评率
	ObjectId              string    `gorm:"object_id;default:null" json:"object_id,omitempty"`                             // 对象id
	AccessManufactureId   int       `gorm:"access_manufacture_id;default:1" json:"access_manufacture_id,omitempty"`        // 物业门禁制造商 1.青岳门禁 2.大华科技 3.海康威视 4.其他 其他需要填写制造商名称
	AccessManufactureName string    `gorm:"access_manufacture_name;default:null" json:"access_manufacture_name,omitempty"` // 门禁制造商名字
}

type CompanyStaff struct {
	ID                           *int       `gorm:"primary_key;column:id" json:"id"`
	CreateTime                   *time.Time `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime                   *time.Time `gorm:"column:update_time;default:null;not null on update current_timestamp" json:"update_time"`
	CompanyID                    *int       `gorm:"company_id;default:0" json:"company_id"`
	No                           *string    `gorm:"no;default:''" json:"no"`
	Gender                       *int       `gorm:"gender;default:0" json:"gender"`
	Name                         *string    `gorm:"name;default:''" json:"name"`
	Nickname                     *string    `gorm:"nickname;default:''" json:"nickname"`
	Avatar                       *string    `gorm:"avatar;default:''" json:"avatar"`
	Wechat                       *string    `gorm:"wechat;default:''" json:"wechat"`
	Qq                           *string    `gorm:"qq;default:''" json:"qq"`
	Email                        *string    `gorm:"email;default:''" json:"email"`
	Position                     *string    `gorm:"position;default:''" json:"position"`
	Mobile                       *string    `gorm:"mobile;default:''" json:"mobile"`
	SsoAccount                   *string    `gorm:"sso_account;default:''" json:"sso_account"`
	Status                       *int       `gorm:"status;default:0" json:"status"`
	Password                     *string    `gorm:"password;default:''" json:"password"`
	Salt                         *string    `gorm:"salt;default:''" json:"salt"`
	EntryTime                    *time.Time `gorm:"entry_time;default:null" json:"entry_time"`
	DueTime                      *time.Time `gorm:"due_time;default:null" json:"due_time"`
	QuitTime                     *time.Time `gorm:"quit_time;default:null" json:"quit_time"`
	Comment                      *string    `gorm:"comment;default:''" json:"comment"`
	MinipStaffOpenid             *string    `gorm:"minip_staff_openid;default:''" json:"minip_staff_openid"`
	StaffType                    *int       `gorm:"staff_type;default:1" json:"staff_type,omitempty"`                                          // 员工类型
	EmergencyContactName         *string    `gorm:"emergency_contact_name;default:''" json:"emergency_contact_name,omitempty"`                 // 紧急联系人姓名
	EmergencyContactMobile       *string    `gorm:"emergency_contact_mobile;default:''" json:"emergency_contact_mobile,omitempty"`             // 紧急联系人电话
	EmergencyContactRelationship *string    `gorm:"emergency_contact_relationship;default:''" json:"emergency_contact_relationship,omitempty"` // 与紧急联系人关系
	Birthday                     *time.Time `gorm:"birthday;default:null" json:"birthday,omitempty"`                                           // 生日
	NationId                     *int       `gorm:"nation_id;default:0" json:"nation_id,omitempty"`                                            // 民族
	CardId                       *string    `gorm:"card_id;default:0" json:"card_id,omitempty"`                                                // 身份证
	Address                      *string    `gorm:"address;default:0" json:"address,omitempty"`                                                // 居住地址
	CensusRegisterType           *int       `gorm:"census_register_type;default:0" json:"census_register_type,omitempty"`                      // 户籍性质
	CensusRegisterLocation       *string    `gorm:"census_register_location;default:0" json:"census_register_location,omitempty"`              // 籍贯
	CensusRegisterAddress        *string    `gorm:"census_register_address;default:0" json:"census_register_address,omitempty"`                // 户籍地址
	PoliticCountenanceType       *int       `gorm:"politic_countenance_type;default:null" json:"politic_countenance_type,omitempty"`           // 政治面貌
	HealthyStatus                *int       `gorm:"healthy_status;default:null" json:"healthy_status,omitempty"`                               // 健康状态
	MaritalStatus                *int       `gorm:"marital_status;default:null" json:"marital_status,omitempty"`                               // 婚姻状态
	WorkWechat                   *string    `gorm:"work_wechat;default:''" json:"work_wechat,omitempty"`                                       // 工作微信号
	EducationId                  *int       `gorm:"education_id;default:1" json:"education_id,omitempty"`                                      // 最高学历
	University                   *string    `gorm:"university;default:''" json:"university,omitempty"`                                         // 毕业院校
	FormalTime                   *time.Time `gorm:"formal_time;default:null" json:"formal_time,omitempty"`                                     // 转正日期
	CardFrontUrl                 *string    `gorm:"card_front_url;default:''" json:"card_front_url,omitempty"`                                 // 身份证正面照
	CardBackUrl                  *string    `gorm:"card_back_url;default:''" json:"card_back_url,omitempty"`                                   // 身份证背面照
	EducationUrl                 *string    `gorm:"education_url;default:1" json:"education_url,omitempty"`                                    // 毕业证url
	DegreeUrl                    *string    `gorm:"degree_url;default:1" json:"degree_url,omitempty"`                                          // 学位证url
	OtherAttach                  *string    `gorm:"other_attach;default:1" json:"other_attach,omitempty"`                                      // 其他附件["", ""]
	OtherProperty                *string    `gorm:"other_property;default:1" json:"other_property,omitempty"`                                  // 其他属性[{"":""}]
	HealthExpiration             *time.Time `gorm:"health_expiration;default:null" json:"health_expiration,omitempty"`                         // 健康证有效期
	Height                       *float64   `gorm:"height;default:null" json:"height,omitempty"`                                               // 身高
	Weight                       *float64   `gorm:"weight;default:null" json:"weight,omitempty"`                                               // 体重
	CommunicationAddress         *string    `gorm:"communication_address;default:''" json:"communication_address,omitempty"`                   // 通讯地址
	RecruitmentSource            *string    `gorm:"recruitment_source;default:''" json:"recruitment_source,omitempty"`                         // 招聘来源
	Remark                       *string    `gorm:"remark;default:''" json:"remark,omitempty"`                                                 // 备注
	SsoId                        *string    `gorm:"sso_id;default:''" json:"sso_id,omitempty"`                                                 // SSOID
}

func TestNewMysqlConnection(t *testing.T) {
	var res EstateGroup
	db := NewMysqlConnection()
	//db.Model(&EstateGroup{}).Create(&EstateGroup{
	//	Name:                  "小明项目",
	//	BelongProvinceId:      1,
	//	BelongCityId:          1,
	//	BelongDistrictId:      1,
	//	BelongCompanyId:       1,
	//	Lonlat:                "1,1",
	//	Address:               "小明地址",
	//	KefuPhone:             "15586001233",
	//	Contact:               "nice",
	//	ContactPhone:          "15586001233",
	//	DirectorStaffId:       1,
	//	Plane:                 "15586001233",
	//	PlaneBase:             "1",
	//	CreateBy:              1,
	//	UpdateBy:              1,
	//	DeleteBy:              1,
	//	DelStatus:             1,
	//	CreateTime:            time.Now(),
	//	UpdateTime:            time.Now(),
	//	AreaTotal:             1,
	//	AreaConstruction:      1,
	//	PlanGraphUrl:          "nice",
	//	AfforestedArea:        1,
	//	AfforestedRate:        1,
	//	ChargeArea:            1,
	//	PlotRate:              1,
	//	EstablishTime:         time.Now(),
	//	Stage:                 1,
	//	Format:                "nice",
	//	FormatName:            "nice",
	//	TakeoverTime:          time.Now(),
	//	BusinessType:          1,
	//	ContractType:          1,
	//	ContractName:          "小明",
	//	ContractNo:            "nice",
	//	ContractStartTime:     time.Now(),
	//	ContractEndTime:       time.Now(),
	//	SignTime:              time.Now(),
	//	PartyA:                "a",
	//	PartyB:                "b",
	//	ContractAmount:        1,
	//	CommissionRate:        1,
	//	Remarks:               "nice",
	//	ChargeRate:            1,
	//	PraiseRate:            1,
	//	ObjectId:              "nice",
	//	AccessManufactureId:   1,
	//	AccessManufactureName: "nice",
	//})
	db.First(&res, 1).Scan(&res)
	fmt.Println(res)

	//自动创建表结构 并指定存储引擎为InnoDB
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&CompanyStaff{})
	if err != nil {
		fmt.Println(err)
	}
}
