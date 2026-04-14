package agent

import "ag/llm"

type Agent struct {
	router *llm.Router
}

func New() *Agent {
	return &Agent{
		router: llm.NewRouter(),
	}
}
