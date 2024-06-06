package utils

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type AgentType int

const (
	Forager_agent AgentType = iota
	Farmer_agent
	Bandit_Agent
)

type Action int

const (
	Stay Action = iota
	Moving
	Attacking
	Joining
	Allocation
)

type EnvironmentType int

const (
	Desert EnvironmentType = iota
	Pasture
	Forest
)

type SeasonType int

const (
	Spring SeasonType = iota
	Summer
	Autumn
	Winter
)

type Greediness int

const (
	Moderate = iota
	Greedy
)

var Greedinessstring = []string{
	"Moderate",
	"Greedy",
}

func (at Greediness) String() string {
	if int(at) >= 0 && int(at) < len(Greedinessstring) {
		return Greedinessstring[at]
	}
	return "Unknown"
}

func (s SeasonType) String() string {
	return [...]string{"Spring", "Summer", "Autumn", "Winter"}[s]
}

var agentTypeNames = []string{
	"Forager_agent",
	"Farmer_agent",
	"Bandit_Agent",
}

// String method returns the string name of the AgentType.
func (at AgentType) String() string {
	if int(at) >= 0 && int(at) < len(agentTypeNames) {
		return agentTypeNames[at]
	}
	return "Unknown"
}
