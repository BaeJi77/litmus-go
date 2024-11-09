package environment

import (
	"strconv"

	experimentTypes "github.com/litmuschaos/litmus-go/pkg/service-mesh/linkerd/types"
	"github.com/litmuschaos/litmus-go/pkg/types"
)

// GetENV fetches all the env variables from the runner pod
func GetENV(experimentDetails *experimentTypes.ExperimentDetails) {
	experimentDetails.ExperimentName = types.Getenv("EXPERIMENT_NAME", "")
	experimentDetails.ChaosNamespace = types.Getenv("CHAOS_NAMESPACE", "litmus")
	experimentDetails.EngineName = types.Getenv("CHAOSENGINE", "")
	experimentDetails.ChaosDuration, _ = strconv.Atoi(types.Getenv("TOTAL_CHAOS_DURATION", "30"))
	experimentDetails.ChaosInterval, _ = strconv.Atoi(types.Getenv("CHAOS_INTERVAL", "10"))
	experimentDetails.RampTime, _ = strconv.Atoi(types.Getenv("RAMP_TIME", "0"))
	experimentDetails.AppNS = types.Getenv("APP_NAMESPACE", "")
	experimentDetails.AppLabel = types.Getenv("APP_LABEL", "")
	experimentDetails.AppKind = types.Getenv("APP_KIND", "")
	experimentDetails.Delay, _ = strconv.Atoi(types.Getenv("STATUS_CHECK_DELAY", "2"))
	experimentDetails.Timeout, _ = strconv.Atoi(types.Getenv("STATUS_CHECK_TIMEOUT", "180"))
	experimentDetails.PodsAffectedPerc, _ = strconv.Atoi(types.Getenv("PODS_AFFECTED_PERC", "0"))
	experimentDetails.ErrorInjectionName = types.Getenv("ERROR_INJECTION_NAME", "")
	experimentDetails.ErrorInjectionNamespace = types.Getenv("ERROR_INJECTION_NAMESPACE", "")
	experimentDetails.HTTPRouteParentRefsName = types.Getenv("HTTP_ROUTE_PARENT_REFS_NAME", "")
	experimentDetails.HTTPRouteParentRefsKind = types.Getenv("HTTP_ROUTE_PARENT_REFS_KIND", "")
	experimentDetails.HTTPRouteParentRefsGroup = types.Getenv("HTTP_ROUTE_PARENT_REFS_GROUP", "")
	experimentDetails.HTTPRouteParentRefsPort, _ = strconv.Atoi(types.Getenv("HTTP_ROUTE_PARENT_REFS_PORT", "0"))
	experimentDetails.ErrorRate, _ = strconv.ParseFloat(types.Getenv("ERROR_RATE", "0"), 64) // (0, 1]
}

// weight 를 딱 맞게 100, 10을 딱 맞게 할 필요는 없음. 딱 맞지 않으면 어떻게 되지? -> 비율 맞춰줘?
// 유리수로 해도 되나? 0.1 이렇게? 만약 안된다면 내가 추가로 *100 해주던가 해서 정수로 바꾸는 작업.

// k8s service만 되게 할것인지? 아니면 Server에 대한 개념도 가능하게 할것인지.
// 일단 service에 대해서 개발해보자. -> linkerd Server에 대한 개념도 가능할수도 있으니까 그것도 고려해서.

// fault injection을 할때 configMap, Deployment, Service
// HTTPRoute라는 것을 만들어서 트래픽을 거기로 흘러들어가도록 함.
// HTTPRoute에서는 k8s Service에 대한 정보.
// HTTPRoute가 똑같은 Service에 대해서 여러개 있다면? 그러니까 기존에 있던게 있고 error injection에 대해서 추가한다면 어떻게 되는거지?

// 필드
// HTTPRoute kind -> k8s Service, Server
// group -> 2개 밖에 없음. k8s Service면 core이고 Server면 policy.linkerd.io
// port는 해당 Service나 Server에 대한 port
// weight -> 0~100 -> 생각해보니까 error rate만 제공해주면 그것을 뺀 나머지를 본 요청으로 보내면 되잖아.
