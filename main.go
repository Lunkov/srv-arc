package main

import (
  "flag"
  "net/http"
  "github.com/gorilla/mux"
  
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "github.com/golang/glog"
  
  "github.com/Lunkov/lib-arc"
)

func main() {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  flag.Parse()
  
  space := arc.NewSpace("etc", true)
  space.LoadFromFiles()
  StageGQL(space)
  RoleGQL(space)
  ServiceGQL(space)
  SystemGQL(space)
  
  router := mux.NewRouter()
  
  schemaGQL, errQL := graphql.NewSchema(defineSchema())
	if errQL != nil {
		glog.Errorf("Error when creating the graphQL schema: %v", errQL)
    glog.Errorf("Error when creating the graphQL schema: %v", fieldsGQL)
    return
	}

	handleGraphQL := handler.New(&handler.Config{
		Schema:     &schemaGQL,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
  
  router.Handle("/graphql", handleGraphQL)
  
  glog.Infof("LOG: Start service")
  err := http.ListenAndServe(":3000", router)
  if err != nil {
    glog.Errorf("ERR: HTTP server: %s", err)
  }
}

