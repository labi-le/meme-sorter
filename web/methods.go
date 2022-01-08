package web

import (
	"meme-sorter/internal"
)

type method struct {
	Config *internal.Config
	Item   *internal.Meme
}

func NewMethod(config *internal.Config, i *internal.Meme) *method {
	return &method{Item: i, Config: config}
}

func (r method) Create() internal.Response {
	err := r.Config.DB.Create(r.Item)
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
	err := r.Config.DB.Update(r.Item)
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
	err := r.Config.DB.Take(r.Item.ID, &meme)
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
	err := r.Config.DB.Delete(r.Item.ID)
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
