package models

import "time"

type Bug struct {
	Id          int64     `json:"id" xorm:"int(11) not null pk autoincr comment('主键ID')"`
	IterationId int64     `json:"iteration_id" xorm:"int(11) not null index(idx_iteration_id) comment('迭代ID')"`
	Name        string    `json:"name" xorm:"varchar(200) not null comment('bug名称')"`
	Description string    `json:"description" xorm:"text not null comment('bug描述')"`
	Level       int64     `json:"level" xorm:"int(11) not null comment('级别：1一般;2中等;3严重')"`
	Type        int64     `json:"type" xorm:"int(11) not null comment('BUG类型：1.代码错误;2.需求设计;3.界面优化;4.配置相关;5.性能问题;6.校验规范;7.测试脚本;8.其他')"`
	Status      int64     `json:"status" xorm:"int(11) not null comment('状态：0待解决;1进行中;2已解决;3已关闭;4延期处理;5不予处理;6重新打开;999已作废')"`
	StaffNo     string    `json:"staff_no" xorm:"varchar(10) not null comment('指派员工工号')"`
	StaffName   string    `json:"staff_name" xorm:"varchar(30) not null comment('指派员工姓名')"`
	EndTime     time.Time `json:"end_time" xorm:"datetime default null comment('解决日期')"`
	CreateBy    string    `json:"create_by" xorm:"varchar(10) not null comment('创建人工号')"`
	Creator     string    `json:"creator" xorm:"varchar(30) not null comment('创建人名字')"`
	DeleteTime  time.Time `json:"delete_time" xorm:"datetime default null comment('删除时间')"`
	UpdateTime  time.Time `json:"update_time" xorm:"datetime NOT NULL DEFAULT current_timestamp() comment('更新时间')"`
	CreateTime  time.Time `json:"create_time" xorm:"datetime NOT NULL DEFAULT current_timestamp() comment('创建时间')"`
	Ttl         int64     `json:"ttl" xorm:"int(11) not null default 0 comment('存活时长')"`
	Ttr         int64     `json:"ttr" xorm:"int(11) not null default 0 comment('响应时长')"`
	ReopenTimes int64     `json:"reopen_times" xorm:"int(11) not null default 0  comment('重新打开次数')"`
}
