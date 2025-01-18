package services

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

func  GetJsonFromUrl(jsonUrl string) *Swagger{
	resp,err:=http.Get(jsonUrl)
	if err != nil {
		log.Printf("error: ",err)
	}
	defer resp.Body.Close()
	if resp.StatusCode!=http.StatusOK {
		log.Fatal("Swagger response not 200. Url: ",jsonUrl)
	}
	body,err:=io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Can't parse json response")
	}

	var swagger Swagger
	err = json.Unmarshal(body, &swagger)
	if err != nil {
		fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &swagger
	/* 
	for paths,methods:= range swagger.Paths{
		for method:= range methods{
			fmt.Printf("My path is: %s, my method is %s",paths,method)
		}
	} */

}

type Swagger struct {
	Paths  map[string]map[string]Endpoint `json:"paths"`
}


type Endpoint struct {
	Tags        []string `json:"tags"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	OperationID string   `json:"operationId"`
	Consumes    []string `json:"consumes"`
	Produces    []string `json:"produces"`
	Parameters  []struct {
		In          string `json:"in"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Required    bool   `json:"required"`
		Schema      struct {
			Ref string `json:"$ref"`
		} `json:"schema"`
	} `json:"parameters"`
}