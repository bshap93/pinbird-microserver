package handlers

import (
	"net/http"
)

type handler struct{}

type IEventHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	AddTag(w http.ResponseWriter, r *http.Request)

	MakePrivateOrPublic(w http.ResponseWriter, r *http.Request)
	ReadOrUnread(w http.ResponseWriter, r *http.Request)
}

func NewPinHander() IEventHandler {
	return &handler{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) AddTag(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Cancel(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Delete implements IEventHandler
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// MakePrivateOrPublic implements IEventHandler
func (h *handler) MakePrivateOrPublic(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// ReadOrUnread implements IEventHandler
func (h *handler) ReadOrUnread(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
