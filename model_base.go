package main

type ModelBase struct {
	Id string `gorm:"primary_key;not null;type:varchar(32)"`
}

func (model ModelBase) GetId() string {
	return model.Id
}
