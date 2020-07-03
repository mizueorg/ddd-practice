package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/trewanek-org/ddd-practice/usecase"
)

type User struct {
	UserApplicationService *usecase.User
}

func NewUser(userApplicationService *usecase.User) *User {
	return &User{UserApplicationService: userApplicationService}
}

func (u *User) List() interface{} {
	return nil
}

func (u *User) Get(id string) interface{} {
	return nil
}

func (u *User) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	str := struct {
		Name string `json:"name"`
	}{}
	b, err := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(b, &str); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	err = u.UserApplicationService.Register(ctx, str.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	log.Println("user created: ", str.Name)
}

func (u *User) Put() {

}

func (u *User) Delete(id string) {

}
