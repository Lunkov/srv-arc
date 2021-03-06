package main

import (
  "github.com/golang/glog"
  "github.com/graphql-go/graphql"
  "github.com/SonicRoshan/straf"

  "github.com/Lunkov/lib-arc"
)

func RoleGQL(s *arc.Space) {
  RoleType, err := straf.GetGraphQLObject(arc.Role{})
  if err != nil {
    glog.Errorf("ERR: RoleGQL: %s", err)
  }
  
  AppendFields2GraphQL("role", &graphql.Field{
			Type: RoleType,
      Args: graphql.FieldConfigArgument{
                "code": &graphql.ArgumentConfig{
                  Description: "code of the Role",
                  Type:graphql.NewNonNull(graphql.String),
                },
              },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        id := p.Args["code"].(string)
				return s.Roles.GetByCODE(id), nil
			},
		})
    
	AppendFields2GraphQL("roles", &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return s.Roles.GetList(), nil
			},
    })
}

