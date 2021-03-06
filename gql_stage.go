package main

import (
  "github.com/golang/glog"
  "github.com/graphql-go/graphql"
  "github.com/SonicRoshan/straf"

  "github.com/Lunkov/lib-arc"
  "github.com/Lunkov/lib-gql"
)

func StageGQL(s *arc.Space) {
  StageType, err := straf.GetGraphQLObject(arc.Stage{})
  if err != nil {
    glog.Errorf("ERR: StageGQL: %s", err)
  }

  gql.AppendQuery("stage", &graphql.Field{
			Type: StageType,
      Args: graphql.FieldConfigArgument{
                "code": &graphql.ArgumentConfig{
                  Description: "code of the Stage",
                  Type:graphql.NewNonNull(graphql.String),
                },
              },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        id := p.Args["code"].(string)
				return s.Stages.GetByCODE(id), nil
			},
		})
    
	gql.AppendQuery("stages", &graphql.Field{
			Type: graphql.NewList(StageType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return s.Stages.GetList(), nil
			},
    })
}

