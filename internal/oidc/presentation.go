package oidc

import (
	"encoding/json"
	"net/http"

	"github.com/example/vc-openid-idp/internal/session"
	"github.com/gorilla/mux"
)

func HandlePresentationRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]

	sess, ok := session.Get(sessionID)
	if !ok {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	issuer := "https://localhost:8080"

	response := map[string]interface{}{
		"client_id":     issuer,
		"scope":         "openid",
		"response_type": "id_token",
		"response_mode": "post",
		"nonce":         sess.Nonce,
		"state":         sess.State,
		"presentation_definition": map[string]interface{}{
			"id": "beta-id-request",
			"input_descriptors": []map[string]interface{}{
				{
					"id": "beta-id",
					"format": map[string]interface{}{
						"ldp_vc": map[string]interface{}{
							"proof_type": []string{"Ed25519Signature2018"},
						},
					},
					"constraints": map[string]interface{}{
						"fields": []map[string]interface{}{
							{"path": []string{"$.credentialSubject.family_name"}},
							{"path": []string{"$.credentialSubject.given_name"}},
						},
					},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
