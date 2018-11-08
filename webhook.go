package go_rocket_webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// WebHook struct with rocket hook url
type WebHook struct {
	hookURL string
}

// WebHookPostPayload struct for create post payload
type WebHookPostPayload struct {
	Text        string        `json:"text,omitempty"`
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"username,omitempty"`
	EmojiIcon   string        `json:"icon_emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// NewWebHook create new webhook instance
func NewWebHook(hookURL string) *WebHook {
	return &WebHook{hookURL}
}

// PostMessage func send message to webhook
func (h *WebHook) PostMessage(payload *WebHookPostPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(h.hookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(t))
	}

	return nil
}
