package oidc

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/example/vc-openid-idp/internal/session"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type authorizeData struct {
	SessionID string
	QRDataURL template.URL
}

func HandleAuthorize(w http.ResponseWriter, r *http.Request) {
	// Pflichtparameter abfragen
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	state := r.URL.Query().Get("state")
	nonce := r.URL.Query().Get("nonce")

	if clientID == "" || redirectURI == "" || state == "" || nonce == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	// Neue Session-ID
	sessionID := uuid.New().String()

	// Session speichern
	session.Save(sessionID, &session.Session{
		ClientID:    clientID,
		RedirectURI: redirectURI,
		State:       state,
		Nonce:       nonce,
		Verified:    false,
	})

	// QR-Link für Wallet (z. B. als `openid-vc://...`)
	host := "192.168.2.115" // Deine PC-IP
	requestURI := fmt.Sprintf("https://%s:8080/presentation-request/%s", host, sessionID)
	qrLink := fmt.Sprintf("openid-vc://?request_uri=%s", requestURI)

	// QR-Code als base64
	png, err := qrcode.Encode(qrLink, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "QR Code generation failed", http.StatusInternalServerError)
		return
	}

	dataURL := "data:image/png;base64," + encodeToBase64(png)

	// HTML rendern
	tmpl := template.Must(template.ParseFiles("web/templates/authorize.html"))
	tmpl.Execute(w, authorizeData{
		SessionID: sessionID,
		QRDataURL: template.URL(dataURL),
	})
}
