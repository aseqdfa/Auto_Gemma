package llm

type Mode string

type Router struct {
	mode      string
	model     string
	providers map[string]LLM
}

func NewRouter() *Router {
	return &Router{
		mode:  "single",
		model: "ollama",
		providers: map[string]LLM{
			"ollama": Ollama{},
			"gpt":    OpenAI{},
			"claude": Claude{},
		},
	}
}

func (r *Router) Get(prompt string) LLM {

	if r.mode == "single" || r.mode == "manual" {
		return r.providers[r.model]
	}

	// auto mode
	if len(prompt) < 100 {
		return r.providers["ollama"]
	}

	return r.providers["gpt"]
}
