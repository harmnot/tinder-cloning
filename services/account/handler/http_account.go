package handler

import (
	"encoding/json"
	"fmt"
	netHttp "net/http"
	"tinder-cloning/pkg/util"
	"tinder-cloning/services/account/schema"
)

func (http *AccountHandler) RegisterHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	var payload schema.RequestRegister
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, "Invalid Payload")
		return
	}

	if err := http.accountService.SignUp(r.Context(), &payload); err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusCreated, map[string]bool{"success": true})
}

func (http *AccountHandler) LoginHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	var payload schema.RequestLogin
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println(err)
		util.RenderJSON(w, netHttp.StatusBadRequest, "Invalid Payload")
		return
	}

	token, err := http.accountService.SingIn(r.Context(), &payload)
	if err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusOK, map[string]string{"token": token})
}
