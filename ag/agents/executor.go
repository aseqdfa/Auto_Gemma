package agent

import (
	"ag/tools"
)

func (a *Agent) Execute(goal string) string {

	out := ""

	for i := 0; i < 5; i++ {

		step := a.router.Get(goal).Generate("Next action: " + goal)

		if step == "" {
			out += "[ERROR empty step]\n"
			continue
		}

		res := tools.Run(step)

		if res == "" {
			res = "[NO OUTPUT]"
		}

		out += res + "\n"

		if a.Critic(goal, res) {
			out += "[DONE]\n"
			break
		}
	}

	return out
}
