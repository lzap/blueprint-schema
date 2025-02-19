package main

import (
	"encoding/json"

	gojson "github.com/goccy/go-json"
	"github.com/kaptinlin/jsonschema"
)

var payload = `{
  "valid": false,
  "evaluationPath": "",
  "schemaLocation": "",
  "instanceLocation": "",
  "annotations": {
    "description": "Blueprint type prototype\n\nThis is just a brief example of a common blueprint structure. Just few fields\nwere selected to demonstrate the JSON schema."
  },
  "errors": {
    "properties": "Properties 'registration', 'network' do not match their schemas",
    "required": "Required properties 'registration', 'network' are missing"
  },
  "details": [
    {
      "valid": false,
      "evaluationPath": "/properties/registration",
      "schemaLocation": "https://github.com/lzap/common-blueprint-example/blueprint#/properties/registration",
      "instanceLocation": "/registration",
      "annotations": {
        "description": "Registration details"
      },
      "errors": {
        "type": "Value is null but should be object"
      }
    },
    {
      "valid": false,
      "evaluationPath": "/properties/network",
      "schemaLocation": "https://github.com/lzap/common-blueprint-example/blueprint#/properties/network",
      "instanceLocation": "/network",
      "annotations": {
        "description": "Networking details"
      },
      "errors": {
        "type": "Value is null but should be object"
      }
    },
    {
      "valid": true,
      "evaluationPath": "/properties/name",
      "schemaLocation": "https://github.com/lzap/common-blueprint-example/blueprint#/properties/name",
      "instanceLocation": "/name",
      "annotations": {
        "description": "Name of the blueprint"
      }
    }
  ]
}
`

func main() {
	var list jsonschema.List
	err := json.Unmarshal([]byte(payload), &list)
	if err != nil {
		panic(err)
	}

	_, err = gojson.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
}
