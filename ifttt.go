package ifttt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const iftttUrlTemplate string = "https://maker.ifttt.com/trigger/%s/with/key/%s"

type IftttClient interface {
	// Triggers an event with the specified parameters
	Trigger(event string, values []string)
}

type iftttClientImpl struct {
	key string
}

func (ifttt *iftttClientImpl) Trigger(event string, values []string) {
	if len(values) > 3 {
		panic("IFTTT supports at most 3 values")
	}
	body := new(bytes.Buffer)
	m := make(map[string]string)
	for i, v := range values {
		m[fmt.Sprintf("value%d", i+1)] = v
	}
	json.NewEncoder(body).Encode(m)
	_, err := http.Post(fmt.Sprintf(
		iftttUrlTemplate, event, ifttt.key), "application/json", body)
	if err != nil {
		panic(err)
	}
}

// Creates a new IftttClient object with the given API key
func NewIftttClient(key string) IftttClient {
	return &iftttClientImpl{key}
}
