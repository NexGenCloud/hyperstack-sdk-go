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
	RequestBody interface{} `yaml:"requestBody"`
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
	type nReq = []interface{}

	for apiPath, item := range spec.Paths {
		operations := map[string]*operation{
			"get":     item.Get,
			"put":     item.Put,
			"post":    item.Post,
			"delete":  item.Delete,
			"options": item.Options,
			"head":    item.Head,
			"patch":   item.Patch,
			"trace":   item.Trace,
		}

		for _, op := range operations {
			if op == nil {
				continue
			}

			appendTaggedSpecs(op, apiPath, item, &spec, tags)
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
				"Failed to generate client code for the tag '%s': %v. oapi-codegen output: \n%s\n",
				tag,
				err,
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

// TODO: remove?
// https://stackoverflow.com/questions/23057785/how-to-deep-copy-a-map-and-then-clear-the-original

type CopyableMap map[interface{}]interface{}
type CopyableSlice []interface{}

// DeepCopy will create a deep copy of this map. The depth of this
// copy is all inclusive. Both maps and slices will be considered when
// making the copy.
func (m CopyableMap) DeepCopy() map[interface{}]interface{} {
	result := map[interface{}]interface{}{}

	for k, v := range m {
		// Handle maps
		mapvalue, isMap := v.(map[interface{}]interface{})
		if isMap {
			result[k] = CopyableMap(mapvalue).DeepCopy()
			continue
		}

		// Handle slices
		slicevalue, isSlice := v.([]interface{})
		if isSlice {
			result[k] = CopyableSlice(slicevalue).DeepCopy()
			continue
		}

		result[k] = v
	}

	return result
}

// DeepCopy will create a deep copy of this slice. The depth of this
// copy is all inclusive. Both maps and slices will be considered when
// making the copy.
func (s CopyableSlice) DeepCopy() []interface{} {
	result := []interface{}{}

	for _, v := range s {
		// Handle maps
		mapvalue, isMap := v.(map[interface{}]interface{})
		if isMap {
			result = append(result, CopyableMap(mapvalue).DeepCopy())
			continue
		}

		// Handle slices
		slicevalue, isSlice := v.([]interface{})
		if isSlice {
			result = append(result, CopyableSlice(slicevalue).DeepCopy())
			continue
		}

		result = append(result, v)
	}

	return result
}
