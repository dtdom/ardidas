package models

type Item struct {
	ID         string `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	Modality   string `json:"modality"`
	Km         string `json:"km"`
	Gender     string `json:"gender"`
	Size       string `json:"size"`
	Hashtag    string `json:"hashtag"`
	ArrayInsta string `json:"arrayinsta"`
	Photo      string `json:"photo"`
	Views      int    `json:"views"`
	Sales      int    `json:"sales"`
}

func (i *Item) Store() {
	db.Create(&i)

}

func (i *Item) Get() {
	db.First(i)
	views := i.Views + 1
	db.Model(i).Update("views", views)
}

func (i *Item) Sell() {
	db.First(i)
	sales := i.Sales + 1
	db.Model(i).Update("sales", sales)
}

func GetAllItems() []Item {
	var items []Item
	db.Find(&items)
	return items
}

func (i *Item) Filter() []Item {
	var params []string
	query := ""
	first := true
	if i.Modality != "" {
		query = "modality = ?"
		first = false
		params = append(params, i.Modality)
	}
	if i.Km != "" {
		if !first {
			query += " AND "
		} else {
			first = true
		}
		query = query + "km = ?"
		params = append(params, i.Km)
	}
	if i.Gender != "" {
		if !first {
			query += " AND "
		} else {
			first = true
		}
		query = query + "gender = ?"
		params = append(params, i.Gender)

	}
	if i.Size != "" {
		if !first {
			query += " AND "
		} else {
			first = true
		}
		query = query + "size = ?"
		params = append(params, i.Size)

	}
	var items []Item

	db.Where(query, params).Find(&items)
	return items
}
