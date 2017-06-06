package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	pgql "github.com/playlyfe/go-graphql"
)

// see https://echo.labstack.com/guide
func main() {
	e := echo.New()
	e.GET("/graphql", graphql1)
	e.Logger.Fatal(e.Start(":3003"))
}

//-----------------------------------------------------------------------------
// use https://github.com/playlyfe/go-graphql
var schema1 = `
	type QueryRoot {
	    hello: String
	}
	`
var resolvers1 = map[string]interface{}{
	"QueryRoot/hello": func(params *pgql.ResolveParams) (interface{}, error) {
		return "world", nil
	},
}
var executor1, _ = pgql.NewExecutor(schema1, "QueryRoot", "", resolvers1)

func graphql1(context echo.Context) error {
	// executor1.ResolveType = func(value interface{}) string {
	// 	if object, ok := value.(map[string]interface{}); ok {
	// 		return object["__typename"].(string)
	// 	}
	// 	return ""
	// }
	query := context.QueryParam("query")
	// body, err := ioutil.ReadAll(context.Request().Body())
	// if err != nil {
	// 	panic(err)
	// }
	// var data map[string]interface{}
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	panic(err)
	// }
	variables := map[string]interface{}{}
	result, err := executor1.Execute(context, query, variables, "")
	if err != nil {
		panic(err)
	}
	rJSON, _ := json.Marshal(result)
	return context.String(http.StatusOK, fmt.Sprintf("%s", rJSON))
}
