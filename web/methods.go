package web

import (
	"meme-sorter/internal"
)

type method struct {
	DB   *internal.DB
	Item *internal.Meme
}

func NewMethod(db *internal.DB, i *internal.Meme) *method {
	return &method{Item: i, DB: db}
}

func (r method) Create() internal.Response {
	err := r.DB.Create(r.Item)
	if err != nil {
		return internal.Response{
			Status:      internal.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return internal.Response{
		Status:      internal.Success,
		Description: "created!",
		Data:        []int{},
	}
}
func (r method) Update() internal.Response {
	err := r.DB.Update(r.Item)
	if err != nil {
		return internal.Response{
			Status:      internal.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return internal.Response{
		Status:      internal.Success,
		Description: "item has been updated",
		Data:        r.Item,
	}
}
func (r method) Take() internal.Response {
	var meme internal.Meme
	err := r.DB.Take(r.Item.ID, &meme)
	if err != nil {
		return internal.Response{
			Status:      internal.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return internal.Response{
		Status:      internal.Success,
		Description: "item received",
		Data:        meme,
	}
}
func (r method) Delete() internal.Response {
	err := r.DB.Delete(r.Item.ID)
	if err.Error != nil {
		return internal.Response{
			Status:      internal.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return internal.Response{
		Status:      internal.Success,
		Description: "item has been deleted",
		Data:        []int{},
	}
}
