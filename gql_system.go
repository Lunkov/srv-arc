package main

import (
  "github.com/golang/glog"
  "github.com/graphql-go/graphql"
  "github.com/SonicRoshan/straf"

  "github.com/Lunkov/lib-arc"
  "github.com/Lunkov/lib-gql"
)

func SystemGQL(s *arc.Space) {
  SystemType, err := straf.GetGraphQLObject(arc.System{})
  if err != nil {
    glog.Errorf("ERR: SystemGQL: %s", err)
  }

  gql.AppendQuery("system", &graphql.Field{
			Type: SystemType,
      Args: graphql.FieldConfigArgument{
                "code": &graphql.ArgumentConfig{
                  Description: "code of the Service",
                  Type:graphql.NewNonNull(graphql.String),
                },
              },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        id := p.Args["code"].(string)
				return s.Services.GetByCODE(id), nil
			},
		})
    
	gql.AppendQuery("systems", &graphql.Field{
			Type: graphql.NewList(SystemType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return s.Services.GetList(), nil
			},
    })
}

