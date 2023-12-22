package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Operation struct {
	Tags        []string    `yaml:"tags"`
	Parameters  []Parameter `yaml:"parameters,omitempty"`
	Summary     string      `yaml:"summary,omitempty"`
	Description string      `yaml:"description,omitempty"`
	OperationId string      `yaml:"operationId,omitempty"`
	Responses   interface{} `yaml:"responses"`
}

type PathItem struct {
	Get     *Operation `yaml:"get,omitempty"`
	Put     *Operation `yaml:"put,omitempty"`
	Post    *Operation `yaml:"post,omitempty"`
	Delete  *Operation `yaml:"delete,omitempty"`
	Options *Operation `yaml:"options,omitempty"`
	Head    *Operation `yaml:"head,omitempty"`
	Patch   *Operation `yaml:"patch,omitempty"`
	Trace   *Operation `yaml:"trace,omitempty"`
}

type OpenAPISpec struct {
	Openapi    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Paths      map[string]*PathItem   `yaml:"paths"`
	Components map[string]interface{} `yaml:"components"`
}

type Schema struct {
	Type string `yaml:"type"`
}
type Parameter struct {
	Name     string `yaml:"name"`
	In       string `yaml:"in"`
	Required bool   `yaml:"required"`
	Schema   Schema `yaml:"schema"`
}

func formatFileName(name string) string {
	fileName := strings.ReplaceAll(name, " ", "_")
	fileName = strings.ReplaceAll(fileName, "-", "_")
	return fileName
}

func main() {
	specBytes, err := ioutil.ReadFile("api.json")
	checkErr(err)

	var spec OpenAPISpec
	err = yaml.Unmarshal(specBytes, &spec)
	checkErr(err)

	os.MkdirAll("./api", os.ModePerm)
	os.MkdirAll("./gosdk", os.ModePerm)

	tags := make(map[string]*OpenAPISpec, 0)

	for path, pathItem := range spec.Paths {
		if pathItem.Get != nil {
			appendTaggedSpecs(pathItem.Get, path, pathItem, &spec, tags)
		}
		if pathItem.Put != nil {
			appendTaggedSpecs(pathItem.Put, path, pathItem, &spec, tags)
		}
		if pathItem.Post != nil {
			appendTaggedSpecs(pathItem.Post, path, pathItem, &spec, tags)
		}
		if pathItem.Delete != nil {
			appendTaggedSpecs(pathItem.Delete, path, pathItem, &spec, tags)
		}
		if pathItem.Options != nil {
			appendTaggedSpecs(pathItem.Options, path, pathItem, &spec, tags)
		}
		if pathItem.Head != nil {
			appendTaggedSpecs(pathItem.Head, path, pathItem, &spec, tags)
		}
		if pathItem.Patch != nil {
			appendTaggedSpecs(pathItem.Patch, path, pathItem, &spec, tags)
		}
		if pathItem.Trace != nil {
			appendTaggedSpecs(pathItem.Trace, path, pathItem, &spec, tags)
		}
	}

	for tag, tagSpec := range tags {
		tagBytes, err := yaml.Marshal(tagSpec)
		checkErr(err)

		folderName := formatFileName(tag)
		err = os.MkdirAll("./api/"+folderName, os.ModePerm)
		checkErr(err)
		err = os.MkdirAll("./gosdk/"+folderName, os.ModePerm)
		checkErr(err)

		tagSpecPath := "./api/" + folderName + "/" + folderName + ".yaml"
		err = ioutil.WriteFile(tagSpecPath, tagBytes, 0644)
		checkErr(err)

		outPackage := folderName
		cmd := exec.Command("oapi-codegen", "-package", outPackage, "-generate", "types,client", tagSpecPath)
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to generate client code for the tag '%s'. oapi-codegen output: \n%s\n", tag, string(cmdOutput))
			continue
		}

		outputPath := "./gosdk/" + folderName + "/" + folderName + "_client.go"
		err = ioutil.WriteFile(outputPath, cmdOutput, 0644)
		checkErr(err)
	}
}

func appendTaggedSpecs(operation *Operation, path string, pathItem *PathItem, spec *OpenAPISpec, tags map[string]*OpenAPISpec) {
	if operation.Tags == nil || len(operation.Tags) == 0 {
		tags["Other"] = appendSpecToTag("Other", path, pathItem, spec, tags)
	} else {
		for _, tag := range operation.Tags {
			tags[tag] = appendSpecToTag(tag, path, pathItem, spec, tags)
		}
	}
}

// Modify appendSpecToTag function
func appendSpecToTag(tag string, path string, pathItem *PathItem, spec *OpenAPISpec, tags map[string]*OpenAPISpec) *OpenAPISpec {
	var tagSpec *OpenAPISpec
	if ts, ok := tags[tag]; ok {
		tagSpec = ts
	} else {
		tagSpec = &OpenAPISpec{
			Openapi:    spec.Openapi,
			Info:       spec.Info,
			Paths:      make(map[string]*PathItem),
			Components: spec.Components,
		}
	}

	tagSpec.Paths[path] = pathItem
	return tagSpec
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
