package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcFunctions = `{
  "block": {
    "attributes": {
      "functions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ca_port": "number",
              "code_checksum": "string",
              "code_size": "number",
              "creation_time": "string",
              "custom_container_config": [
                "list",
                [
                  "object",
                  {
                    "args": "string",
                    "command": "string",
                    "image": "string"
                  }
                ]
              ],
              "description": "string",
              "environment_variables": [
                "map",
                "string"
              ],
              "handler": "string",
              "id": "string",
              "initialization_timeout": "number",
              "initializer": "string",
              "instance_concurrency": "number",
              "instance_type": "string",
              "last_modification_time": "string",
              "memory_size": "number",
              "name": "string",
              "runtime": "string",
              "timeout": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcFunctionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcFunctions), &result)
	return &result
}
