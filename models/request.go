package models

type Request struct {
	ID        string `gorm:"primary_key" json:"id"`
	NameUser  string `json:"name_user"`
	PhotoUser string `json:"photo_user"`
	ItemId    string `json:"item_id"`
	Status    string `json:"status"`
}

func (r *Request) Store() {
	db.Create(&r)
}

func (r *Request) Get() {
	db.First(r)
}

func GetAllRequests() []Request {
	var requests []Request
	db.Where("status = ?", "processing").Find(&requests)
	return requests
}

func (r *Request) Complete() {
	db.Model(r).Update("status", "completed")
}
