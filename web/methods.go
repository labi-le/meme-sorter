package web

import (
	"meme-sorter/internal"
	"meme-sorter/internal/structures"
)

type method struct {
	DB   *internal.DB
	Item *structures.Meme
}

func NewMethod(db *internal.DB, i *structures.Meme) *method {
	return &method{Item: i, DB: db}
}

func (r method) Create() structures.Response {
	err := r.DB.Create(r.Item)
	if err != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return structures.Response{
		Status:      structures.Success,
		Description: "created!",
		Data:        []int{},
	}
}
func (r method) Update() structures.Response {
	err := r.DB.Update(r.Item)
	if err != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return structures.Response{
		Status:      structures.Success,
		Description: "item has been updated",
		Data:        r.Item,
	}
}
func (r method) Read() structures.Response {
	var meme structures.Meme
	err := r.DB.Take(r.Item.ID, &meme)
	if err != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return structures.Response{
		Status:      structures.Success,
		Description: "item has been deleted",
		Data:        meme,
	}
}
func (r method) Delete() structures.Response {
	err := r.DB.Delete(r.Item.ID)
	if err.Error != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: err.Error(),
			Data:        []int{},
		}
	}

	return structures.Response{
		Status:      structures.Success,
		Description: "item has been deleted",
		Data:        []int{},
	}
}
