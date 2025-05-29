package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv3Triggers = `{
  "block": {
    "attributes": {
      "function_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "triggers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "http_trigger": [
                "list",
                [
                  "object",
                  {
                    "url_internet": "string",
                    "url_intranet": "string"
                  }
                ]
              ],
              "id": "string",
              "invocation_role": "string",
              "last_modified_time": "string",
              "qualifier": "string",
              "source_arn": "string",
              "status": "string",
              "target_arn": "string",
              "trigger_config": "string",
              "trigger_id": "string",
              "trigger_name": "string",
              "trigger_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcv3TriggersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv3Triggers), &result)
	return &result
}
