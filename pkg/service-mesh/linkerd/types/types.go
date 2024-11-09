package types

// ADD THE ATTRIBUTES OF YOUR CHOICE HERE
// FEW MANDATORY ATTRIBUTES ARE ADDED BY DEFAULT

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	ExperimentName            string
	EngineName                string
	ChaosDuration             int
	ChaosInterval             int
	RampTime                  int
	AppNS                     string
	AppLabel                  string
	AppKind                   string
	ChaosNamespace            string
	ChaosPodName              string
	Timeout                   int
	Delay                     int
	IsTargetContainerProvided bool
	ErrorInjectionName        string
	ErrorInjectionNamespace   string
	HTTPRouteParentRefsName   string
	HTTPRouteParentRefsKind   string
	HTTPRouteParentRefsGroup  string
	HTTPRouteParentRefsPort   int
	ErrorRate                 float64
}
