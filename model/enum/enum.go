package enum

type Time string

const (
	FormatTime Time = "2006-01-02 15:04:05"
	//删除状态
	StatusUndeleted = 0 //未删除
	StatusDeleted   = 1 //已删除
)