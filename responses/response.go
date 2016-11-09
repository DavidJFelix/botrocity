package responses

import "net/http"

// MattermostWebhookResponse is an interface for responding to mattermost
type MattermostWebhookResponse interface {
	Write(w http.ResponseWriter)
}
