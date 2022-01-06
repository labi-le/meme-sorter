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
	result := r.DB.Create(r.Item)
	if result.Error != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: result.Error.Error(),
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
	result := r.DB.Update(r.Item)
	if result.Error != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: result.Error.Error(),
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
	result := r.DB.Take(r.Item.ID)
	//if result.Error != nil {
	//	return structures.Response{
	//		Status:      structures.Failed,
	//		Description: result.Error.Error(),
	//		Data:        []int{},
	//	}
	//}

	return structures.Response{
		Status:      structures.Success,
		Description: "item has been deleted",
		Data:        result,
	}
}
func (r method) Delete() structures.Response {
	result := r.DB.Delete(r.Item.ID)
	if result.Error != nil {
		return structures.Response{
			Status:      structures.Failed,
			Description: result.Error.Error(),
			Data:        []int{},
		}
	}

	return structures.Response{
		Status:      structures.Success,
		Description: "item has been deleted",
		Data:        []int{},
	}
}
