package responses

import (
	"encoding/json"
	"net/http"
)

// MattermostTextResponse is the prefered JSON structure for responding
// to a mattermost webhook
type MattermostTextResponse struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	IconURL  string `json:"icon_url,omitempty"`
}

func (m *MattermostTextResponse) Write(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(m)
}
