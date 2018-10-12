package api

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	soq "github.com/impactasaurus/soq-api"
	"github.com/impactasaurus/soq-api/status"
)

type v1 struct {
	status *status.Provider
}

type RouteHandler struct {
	Route   string
	Handler http.Handler
}

type QuestionnaireFetcher interface {
	Questionnaire(id string) (soq.Questionnaire, error)
	Questionnaires(page, limit int) (soq.QuestionnaireList, error)
}

// NewV1 returns a set of RouteHandler which serve V1 of the API
func NewV1(qf QuestionnaireFetcher) ([]RouteHandler, error) {
	v := &v1{
		status: status.New(),
	}
	return []RouteHandler{{
		Route:   "/v1/status",
		Handler: v.statusHandler(),
	}, {
		Route:   "/v1/playground",
		Handler: handler.Playground("GraphQL playground", "/v1/query"),
	}, {
		Route: "/v1/query",
		Handler: handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver{
			fetcher: qf,
		}})),
	}}, nil
}
