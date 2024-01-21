package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type operation struct {
	Tags        []string    `yaml:"tags"`
	Parameters  []parameter `yaml:"parameters,omitempty"`
	Summary     string      `yaml:"summary,omitempty"`
	Description string      `yaml:"description,omitempty"`
	OperationId string      `yaml:"operationId,omitempty"`
	Responses   interface{} `yaml:"responses"`
}

type pathItem struct {
	Get     *operation `yaml:"get,omitempty"`
	Put     *operation `yaml:"put,omitempty"`
	Post    *operation `yaml:"post,omitempty"`
	Delete  *operation `yaml:"delete,omitempty"`
	Options *operation `yaml:"options,omitempty"`
	Head    *operation `yaml:"head,omitempty"`
	Patch   *operation `yaml:"patch,omitempty"`
	Trace   *operation `yaml:"trace,omitempty"`
}

type openAPISpec struct {
	Openapi    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Paths      map[string]*pathItem   `yaml:"paths"`
	Components map[string]interface{} `yaml:"components"`
}

type schema struct {
	Type string `yaml:"type"`
}
type parameter struct {
	Name     string `yaml:"name"`
	In       string `yaml:"in"`
	Required bool   `yaml:"required"`
	Schema   schema `yaml:"schema"`
}

func formatFileName(name string) string {
	fileName := strings.ReplaceAll(name, " ", "_")
	fileName = strings.ReplaceAll(fileName, "-", "_")
	return fileName
}
func main() {
	var err error

	dirArtifacts := os.Getenv("DIR_ARTIFACTS")
	if dirArtifacts == "" {
		dirArtifacts = "."
	}

	dirApi := path.Join(dirArtifacts, "api")
	err = os.RemoveAll(dirApi)
	checkErr(err)
	err = os.MkdirAll(dirApi, os.ModePerm)
	checkErr(err)

	dirSdk := os.Getenv("DIR_PACKAGE")
	if dirSdk == "" {
		checkErr(fmt.Errorf("DIR_PACKAGE must be specified"))
	}
	err = os.RemoveAll(dirSdk)
	checkErr(err)
	err = os.MkdirAll(dirSdk, os.ModePerm)
	checkErr(err)

	specBytes, err := ioutil.ReadFile(path.Join(dirArtifacts, "api.json"))
	checkErr(err)

	var spec openAPISpec
	err = yaml.Unmarshal(specBytes, &spec)
	checkErr(err)

	tags := make(map[string]*openAPISpec, 0)

	type nMap = map[interface{}]interface{}

	for apiPath, pathItem := range spec.Paths {
		switch apiPath {
		case "/auth/roles/{id}":
			spec.Components["schemas"].(nMap)["RBAC Role"].(nMap)["properties"].(nMap)["roles"] = spec.Components["schemas"].(nMap)["RBAC Role"].(nMap)["properties"].(nMap)["role"]
			delete(spec.Components["schemas"].(nMap)["RBAC Role"].(nMap)["properties"].(nMap), "role")
		}

		if pathItem.Get != nil {
			appendTaggedSpecs(pathItem.Get, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Put != nil {
			appendTaggedSpecs(pathItem.Put, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Post != nil {
			appendTaggedSpecs(pathItem.Post, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Delete != nil {
			appendTaggedSpecs(pathItem.Delete, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Options != nil {
			appendTaggedSpecs(pathItem.Options, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Head != nil {
			appendTaggedSpecs(pathItem.Head, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Patch != nil {
			appendTaggedSpecs(pathItem.Patch, apiPath, pathItem, &spec, tags)
		}
		if pathItem.Trace != nil {
			appendTaggedSpecs(pathItem.Trace, apiPath, pathItem, &spec, tags)
		}
	}

	for tag, tagSpec := range tags {
		tagBytes, err := yaml.Marshal(tagSpec)
		checkErr(err)

		folderName := formatFileName(tag)
		err = os.MkdirAll(path.Join(dirApi, folderName), os.ModePerm)
		checkErr(err)
		err = os.MkdirAll(path.Join(dirSdk, folderName), os.ModePerm)
		checkErr(err)

		tagSpecPath := path.Join(dirApi, folderName, folderName+".yaml")
		err = ioutil.WriteFile(tagSpecPath, tagBytes, 0644)
		checkErr(err)

		outPackage := folderName
		cmd := exec.Command(
			"oapi-codegen",
			"-package",
			outPackage,
			"-generate",
			"types,client",
			tagSpecPath,
		)
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf(
				"Failed to generate client code for the tag '%s'. oapi-codegen output: \n%s\n",
				tag,
				string(cmdOutput),
			)
			continue
		}

		outputPath := path.Join(dirSdk, folderName, folderName+"_client.go")
		err = ioutil.WriteFile(outputPath, cmdOutput, 0644)
		checkErr(err)
	}
}

func appendTaggedSpecs(operation *operation, path string, pathItem *pathItem, spec *openAPISpec, tags map[string]*openAPISpec) {
	if operation.Tags == nil || len(operation.Tags) == 0 {
		tags["Other"] = appendSpecToTag("Other", path, pathItem, spec, tags)
	} else {
		for _, tag := range operation.Tags {
			tags[tag] = appendSpecToTag(tag, path, pathItem, spec, tags)
		}
	}
}

// Modify appendSpecToTag function
func appendSpecToTag(tag string, path string, item *pathItem, spec *openAPISpec, tags map[string]*openAPISpec) *openAPISpec {
	var tagSpec *openAPISpec
	if ts, ok := tags[tag]; ok {
		tagSpec = ts
	} else {
		tagSpec = &openAPISpec{
			Openapi:    spec.Openapi,
			Info:       spec.Info,
			Paths:      make(map[string]*pathItem),
			Components: spec.Components,
		}
	}

	tagSpec.Paths[path] = item
	return tagSpec
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
