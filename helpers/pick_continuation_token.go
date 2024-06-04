package helpers

import "github.com/ghoshRitesh12/brooktube/models/search"

func PickContinuationToken(continuations []search.APIRespContinuation) string {
	if len(continuations) > 0 {
		return continuations[0].NextContinuationData.Continuation
	}
	return ""
}
