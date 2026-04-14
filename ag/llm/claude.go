package llm

import "os"

type Claude struct{}

func (c Claude) Generate(prompt string) string {
	key := os.Getenv("ANTHROPIC_API_KEY")
	return "[Claude placeholder] " + key + " -> " + prompt
}
