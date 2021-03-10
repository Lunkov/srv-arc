package main

import (
  "net/http"
  "github.com/golang/glog"
  
  "github.com/Lunkov/lib-gql"
)

func webGraphQL(w http.ResponseWriter, r *http.Request)  {
  keys, ok := r.URL.Query()["query"]
  if !ok || len(keys[0]) < 1 {
    glog.Infof("Url Param 'key' is missing")
    return
  }
       
  query_str := keys[0]  
  w.Write(gql.Query(query_str))
}

