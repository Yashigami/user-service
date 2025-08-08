package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"user-service/internal/handlers"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler { // Создаем структуру, но возвращаем интерфейс
	return &handler{}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)              // Получить список
	router.POST(usersURL, h.CreateUser)          // Создать нового
	router.GET(userURL, h.GetUserByUUID)         // Получить UUID
	router.PUT(userURL, h.UpdateUser)            // Обновить
	router.PATCH(userURL, h.PartiallyUpdateUser) // Частично обновить
	router.DELETE(userURL, h.DeleteUser)         // Удалить

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is list of users"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("this is create user"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this user by uuid user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
}
