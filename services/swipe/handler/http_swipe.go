package handler

import (
	"encoding/json"
	netHttp "net/http"
	"tinder-cloning/pkg/util"
	"tinder-cloning/services/swipe/schema"
)

func (http *SwipeHandler) CreateReactionSwipesHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	// get accountData from context
	accountData, err := util.GetClaimsFromContext(r.Context())
	if err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	var payload schema.RequestSwipe
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, "Invalid Payload")
		return
	}

	payload.AccountID = accountData.AccountID
	if err := http.swipeService.CreateReactionSwipes(r.Context(), payload); err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusOK, map[string]bool{"success": true})
}

func (http *SwipeHandler) GetAllProfileHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	// get accountData from context
	accountData, err := util.GetClaimsFromContext(r.Context())
	if err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	// get query params per_page and page and is_verified
	perPage, page, isVerified := r.URL.Query().Get("per_page"), r.URL.Query().Get("page"), r.URL.Query().Get("is_verified")
	perPageInt, pageInt := util.ConvertStringToInt(perPage), util.ConvertStringToInt(page)
	gender := r.URL.Query().Get("gender")

	var isVerifiedBool *bool
	if isVerified != "" {
		isVerifiedBoolConverted := util.ConvertStringToBool(isVerified)
		isVerifiedBool = &isVerifiedBoolConverted
	}

	filter := schema.ProfileFilter{
		CurrentAccountID: accountData.AccountID,
		PerPage:          perPageInt,
		Page:             pageInt,
		IsVerified:       isVerifiedBool,
		Gender:           gender,
	}

	accounts, err := http.swipeService.GetAllProfile(r.Context(), filter)
	if err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusOK, accounts)
}
