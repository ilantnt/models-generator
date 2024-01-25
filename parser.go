package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// Read the YAML file
	yamlFile, err := ioutil.ReadFile("models.yaml")
	if err != nil {
		log.Fatalf("failed to read YAML file: %v", err)
	}
	log.Println("Loaded models yaml")

	var data map[string]interface{}
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal YAML data: %v", err)
	}
	modelsData, ok := data["models"].(map[interface{}]interface{})
	if !ok {
		log.Fatal("missing 'models' section in YAML")
	}

	// Open a new Go file for writing
	if _, err := os.Stat("generated-models"); os.IsNotExist(err) {
		err := os.Mkdir("generated-models", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	file, err := os.Create("generated-models/models.go")
	if err != nil {
		log.Fatalf("failed to create Go file: %v", err)
	}
	defer file.Close()

	// Write the package header
	fmt.Fprintln(file, "package generated_models\n")

	// Loop through each model and generate structs
	for modelName, modelFields := range modelsData {
		// Write the struct definition to the file
		log.Println(fmt.Sprintf("Load model- %s", modelName))
		fmt.Fprintf(file, "type %s struct {\n", modelName)
		generateFields(file, modelFields)
		fmt.Fprintln(file, "}\n")
	}
	log.Println("Successfully genereated models")
}

func generateFields(file *os.File, fields interface{}) {
	switch v := fields.(type) {
	case map[interface{}]interface{}:
		for fieldName, fieldType := range v {
			fmt.Fprintf(file, "\t%s %s\n", fieldName, getGoType(fieldType))
		}
	default:
		log.Fatal("unexpected type encountered while generating fields")
	}
}

func getGoType(fieldType interface{}) string {
	switch t := fieldType.(type) {
	case string:
		return mapYAMLTypeToGoType(t, nil)
	case map[interface{}]interface{}:
		if t["type"] != nil {
			if t["itemType"] != nil {
				return mapYAMLTypeToGoType(t["type"].(string), t["itemType"].(string))
			} else {
				return mapYAMLTypeToGoType(t["type"].(string), nil)
			}
		}
	}
	return ""
}

func mapYAMLTypeToGoType(yamlType string, itemType interface{}) string {
	switch yamlType {
	case "list":
		if itemType != nil {
			return fmt.Sprintf("[]%s", itemType.(string))
		} else {
			return "[]string"
		}
	case "number":
		return "float64"
	case "string":
		return "string"
	// Add more mappings as needed
	default:
		return "interface{}" // Default to interface{} for unknown types
	}
}
