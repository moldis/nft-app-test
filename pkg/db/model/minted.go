package model

type MintedCollections struct {
	ID         uint   `gorm:"primaryKey"`
	Collection string `gorm:"column:collection"`
	Recipient  string `gorm:"column:recipient"`
	TokenID    string `gorm:"column:token_id"`
	TokenURL   string `gorm:"column:token_url"`
	Updated    int64  `gorm:"autoUpdateTime:milli"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
}

func (MintedCollections) TableName() string {
	return "minted"
}
