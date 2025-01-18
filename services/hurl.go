package services

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateHurlFile(url string)  *os.File{
	f,err:=os.Create("tests.hurl")
	if err != nil {
		log.Fatalf("Error creating hurl file")
	}
	_, err = f.WriteString(BuildFileContent(SwaggerToHurl(url)))
	if err != nil {
		fmt.Println(err)
        f.Close()
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func BuildRequest(hurl Hurl)string {
	var sb strings.Builder
	sb.WriteString(hurl.RequestType)
	sb.WriteString(" ")
	sb.WriteString(hurl.BaseUrl)
	sb.WriteString(hurl.Endpoint)
	sb.WriteString("\n")
	return sb.String()
}

func BuildFileContent(hurls []Hurl) string{
	var content strings.Builder
	for _,hurl:=range hurls{
		content.WriteString(BuildRequest(hurl))
	}
	return content.String()
}

func SwaggerToHurl(url string) []Hurl {
	swagger:=GetJsonFromUrl(url)
	var hurl []Hurl
	var sb strings.Builder
	sb.WriteString(swagger.Host)
	sb.WriteString(swagger.BasePath)
	for paths,methods:= range swagger.Paths{
		for method:= range methods{
			newHurl:=Hurl{}
			newHurl.BaseUrl = sb.String()
			newHurl.RequestType = strings.ToUpper(method)
			newHurl.Endpoint = paths
			hurl = append(hurl, newHurl)
		}
	}
	return hurl
}

type Hurl struct {
	BaseUrl string
	RequestType string
	Endpoint string

	// optional fields
	Body string
	Headers HurlHeader
	QueryParams HurlQueryParam
	PathParams HurlPathParam
}

type HurlHeader struct{
	HeaderParam map[string]string
}
type HurlQueryParam struct{
	QueryParam map[string]string
}
type HurlPathParam struct{
	PathParam map[string]string
}