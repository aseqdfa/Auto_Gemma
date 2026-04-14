package llm

import "os"

type OpenAI struct{}

func (o OpenAI) Generate(prompt string) string {
	key := os.Getenv("OPENAI_API_KEY")
	return "[GPT placeholder] " + key + " -> " + prompt
}
