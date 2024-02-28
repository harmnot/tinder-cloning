package handler

import (
	"encoding/json"
	netHttp "net/http"
	"tinder-cloning/pkg/util"
	"tinder-cloning/services/membership/schema"
)

func (http *MembershipHandler) GetFeaturesHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	// get accountData from context
	accountData, ok := r.Context().Value("accountData").(util.AccountDataClaims)
	if !ok {
		util.RenderJSON(w, netHttp.StatusBadRequest, "Invalid Payload Token")
		return
	}

	features, err := http.membershipService.GetFeatureMembership(r.Context(), accountData.AccountID)
	if err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusOK, features)
}

func (http *MembershipHandler) UpgradeMembershipHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	accountData, ok := r.Context().Value("accountData").(util.AccountDataClaims)
	if !ok {
		util.RenderJSON(w, netHttp.StatusBadRequest, "Invalid Payload Token")
		return
	}

	var upgradeMembership schema.UpgradeMembership
	if err := json.NewDecoder(r.Body).Decode(&upgradeMembership); err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	upgradeMembership.AccountID = accountData.AccountID
	if err := http.membershipService.UpdateOne(r.Context(), &upgradeMembership); err != nil {
		util.RenderJSON(w, netHttp.StatusBadRequest, err.Error())
		return
	}

	util.RenderJSON(w, netHttp.StatusOK, map[string]bool{"success": true})
}
