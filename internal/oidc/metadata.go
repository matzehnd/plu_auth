package oidc

import (
    "encoding/json"
    "net/http"
)

func HandleMetadata(w http.ResponseWriter, r *http.Request) {
    metadata := map[string]interface{}{
        "issuer": "http://localhost:8080",
        "authorization_endpoint": "http://localhost:8080/authorize",
        "token_endpoint": "http://localhost:8080/token",
        "jwks_uri": "http://localhost:8080/jwks.json",
        "response_types_supported": []string{"code", "id_token"},
        "subject_types_supported": []string{"public"},
        "id_token_signing_alg_values_supported": []string{"RS256"},
        "scopes_supported": []string{"openid", "profile"},
        "claims_supported": []string{"sub", "given_name", "family_name", "birthdate", "ol_number"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metadata)
}
