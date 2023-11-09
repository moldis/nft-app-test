package model

type Collections struct {
	ID        uint   `gorm:"primaryKey"`
	Address   string `gorm:"column:address"`
	Symbol    string `gorm:"column:symbol"`
	Name      string `gorm:"column:name"`
	Updated   int64  `gorm:"autoUpdateTime:milli"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func (Collections) TableName() string {
	return "collections"
}
