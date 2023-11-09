package service

import (
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"log"
	"net/http"
	"time"
)

type WebhookNotifier struct{}

func NewWebhookNotifier() *WebhookNotifier {
	return &WebhookNotifier{}
}

func (n *WebhookNotifier) SendNotification(message, target string) error {
	payload, err := sonic.Marshal(map[string]string{"message": message})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", target, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook call failed: %s", resp.Status)
	}

	return nil
}
