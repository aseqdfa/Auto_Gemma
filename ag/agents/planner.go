package agent

func (a *Agent) Plan(goal string) string {
	return a.router.Get(goal).Generate("Plan: " + goal)
}