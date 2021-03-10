package main

import (
  "github.com/golang/glog"
  "github.com/graphql-go/graphql"
  "github.com/SonicRoshan/straf"

  "github.com/Lunkov/lib-arc"
  "github.com/Lunkov/lib-gql"
)

func ServiceGQL(s *arc.Space) {
  ServiceType, err := straf.GetGraphQLObject(arc.Service{})
  if err != nil {
    glog.Errorf("ERR: ServiceGQL: %s", err)
  }

  gql.AppendQuery("service", &graphql.Field{
			Type: ServiceType,
      Args: graphql.FieldConfigArgument{
                "code": &graphql.ArgumentConfig{
                           Description: "code of the Service",
                           Type: graphql.NewNonNull(graphql.String),
                },
              },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        id := p.Args["code"].(string)
				return s.Services.GetByCODE(id), nil
			},
		})
    
	gql.AppendQuery("services", &graphql.Field{
			Type: graphql.NewList(ServiceType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return s.Services.GetList(), nil
			},
    })
}

