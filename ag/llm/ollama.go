package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Ollama struct{}

func (o Ollama) Generate(prompt string) string {
	body := map[string]interface{}{
		"model":  "gemma:4b",
		"prompt": prompt,
		"stream": false,
	}

	b, _ := json.Marshal(body)

	resp, err := http.Post("http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(b))

	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var r map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&r)

	return r["response"].(string)
}
