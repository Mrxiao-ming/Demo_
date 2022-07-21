package estate_group_model

import (
	"time"
)

type EstateGroup struct {
	Id                    *int       `gorm:"id;default:null" json:"id,omitempty"`                                           // 小区id
	Name                  *string    `gorm:"name;default:''" json:"name,omitempty"`                                         // 名字
	BelongProvinceId      *int       `gorm:"belong_province_id;default:0" json:"belong_province_id,omitempty"`              // 所属省份id
	BelongCityId          *int       `gorm:"belong_city_id;default:0" json:"belong_city_id,omitempty"`                      // 所属城市id
	BelongDistrictId      *int       `gorm:"belong_district_id;default:0" json:"belong_district_id,omitempty"`              // 所属区ID
	BelongCompanyId       *int       `gorm:"belong_company_id;default:0" json:"belong_company_id,omitempty"`                // 所属公司id
	Lonlat                *string    `gorm:"lonlat;default:''" json:"lonlat,omitempty"`                                     // 经纬度,用逗号分割
	Address               *string    `gorm:"address;default:''" json:"address,omitempty"`                                   // 项目所在地
	KefuPhone             *string    `gorm:"kefu_phone;default:''" json:"kefu_phone,omitempty"`                             // 客服电话,如果有多个客服电话,请用英文的逗号分割。
	Contact               *string    `gorm:"contact;default:''" json:"contact,omitempty"`                                   // 联系人
	ContactPhone          *string    `gorm:"contact_phone;default:''" json:"contact_phone,omitempty"`                       // 负责人电话
	DirectorStaffId       *int       `gorm:"director_staff_id;default:null" json:"director_staff_id,omitempty"`             // 负责人员工id
	Plane                 *string    `gorm:"plane;default:null" json:"plane,omitempty"`                                     // 平面图数据
	PlaneBase             *string    `gorm:"plane_base;default:''" json:"plane_base,omitempty"`                             // 平面图底图
	CreateBy              *int       `gorm:"create_by;default:null" json:"create_by,omitempty"`                             // 创建者
	UpdateBy              *int       `gorm:"update_by;default:null" json:"update_by,omitempty"`                             // 更新者
	DeleteBy              *int       `gorm:"delete_by;default:null" json:"delete_by,omitempty"`                             // 删除者
	DelStatus             *int       `gorm:"del_status;default:0" json:"del_status,omitempty"`                              // 0正常 1删除
	CreateTime            *time.Time `gorm:"create_time;default:CURRENT_TIMESTAMP" json:"create_time,omitempty"`            // 创建时间
	UpdateTime            *time.Time `gorm:"update_time;default:CURRENT_TIMESTAMP" json:"update_time,omitempty"`            // 更新时间
	AreaTotal             *float64   `gorm:"area_total;default:0.00" json:"area_total,omitempty"`                           // 土地面积
	AreaConstruction      *float64   `gorm:"area_construction;default:0.00" json:"area_construction,omitempty"`             // 建筑面积
	PlanGraphUrl          *string    `gorm:"plan_graph_url;default:''" json:"plan_graph_url,omitempty"`                     // 平面图地址
	AfforestedArea        *float64   `gorm:"afforested_area;default:0.00" json:"afforested_area,omitempty"`                 // 绿化面积
	AfforestedRate        *float64   `gorm:"afforested_rate;default:0.00" json:"afforested_rate,omitempty"`                 // 绿化率
	ChargeArea            *float64   `gorm:"charge_area;default:0.00" json:"charge_area,omitempty"`                         // 计费面积
	PlotRate              *float64   `gorm:"plot_rate;default:0.00" json:"plot_rate,omitempty"`                             // 容积率
	EstablishTime         *time.Time `gorm:"establish_time;default:null" json:"establish_time,omitempty"`                   // 项目成立日期
	Stage                 *int       `gorm:"stage;default:0" json:"stage,omitempty"`                                        // 项目阶段
	Format                *string    `gorm:"format;default:''" json:"format,omitempty"`                                     // 项目业态
	FormatName            *string    `gorm:"format_name;default:''" json:"format_name,omitempty"`                           // 项目业态名称
	TakeoverTime          *time.Time `gorm:"takeover_time;default:null" json:"takeover_time,omitempty"`                     // 物业接管时间
	BusinessType          *int       `gorm:"business_type;default:1" json:"business_type,omitempty"`                        // 业务类型
	ContractType          *int       `gorm:"contract_type;default:1" json:"contract_type,omitempty"`                        // 合同类型
	ContractName          *string    `gorm:"contract_name;default:''" json:"contract_name,omitempty"`                       // 合同名称
	ContractNo            *string    `gorm:"contract_no;default:''" json:"contract_no,omitempty"`                           // 合同编号
	ContractStartTime     *time.Time `gorm:"contract_start_time;default:null" json:"contract_start_time,omitempty"`         // 合同开始时间
	ContractEndTime       *time.Time `gorm:"contract_end_time;default:null" json:"contract_end_time,omitempty"`             // 合同结束时间
	SignTime              *time.Time `gorm:"sign_time;default:null" json:"sign_time,omitempty"`                             // 签约时间
	PartyA                *string    `gorm:"party_a;default:''" json:"party_a,omitempty"`                                   // 甲方名称
	PartyB                *string    `gorm:"party_b;default:''" json:"party_b,omitempty"`                                   // 乙方名称
	ContractAmount        *float64   `gorm:"contract_amount;default:0.00" json:"contract_amount,omitempty"`                 // 合同金额
	CommissionRate        *float64   `gorm:"commission_rate;default:0.00" json:"commission_rate,omitempty"`                 // 佣金比例
	Remarks               *string    `gorm:"remarks;default:''" json:"remarks,omitempty"`                                   // 备注
	ChargeRate            *float64   `gorm:"charge_rate;default:0" json:"charge_rate,omitempty"`                            // 收费率
	PraiseRate            *float64   `gorm:"praise_rate;default:0" json:"praise_rate,omitempty"`                            // 好评率
	ObjectId              *string    `gorm:"object_id;default:null" json:"object_id,omitempty"`                             // 对象id
	AccessManufactureId   *int       `gorm:"access_manufacture_id;default:1" json:"access_manufacture_id,omitempty"`        // 物业门禁制造商 1.青岳门禁 2.大华科技 3.海康威视 4.其他 其他需要填写制造商名称
	AccessManufactureName *string    `gorm:"access_manufacture_name;default:null" json:"access_manufacture_name,omitempty"` // 门禁制造商名字
}

type GetCountGroupByCityIDResponse struct {
	CityCountList []CityCountModel
}

type CityCountModel struct {
	CityID int `gorm:"city_id"  json:"city_id"`
	Count  int `gorm:"count"    json:"count"`
}
