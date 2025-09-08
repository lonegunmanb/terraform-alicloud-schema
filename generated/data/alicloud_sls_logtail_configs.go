package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsLogtailConfigs = `{
  "block": {
    "attributes": {
      "configs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "logtail_config_name": "string"
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
      "logstore_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logtail_config_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "offset": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlsLogtailConfigsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsLogtailConfigs), &result)
	return &result
}
