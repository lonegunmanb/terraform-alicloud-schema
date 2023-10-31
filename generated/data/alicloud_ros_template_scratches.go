package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRosTemplateScratches = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scratches": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "id": "string",
              "logical_id_strategy": "string",
              "preference_parameters": [
                "list",
                [
                  "object",
                  {
                    "parameter_key": "string",
                    "parameter_value": "string"
                  }
                ]
              ],
              "source_resource_group": [
                "list",
                [
                  "object",
                  {
                    "resource_group_id": "string",
                    "resource_type_filter": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "source_resources": [
                "list",
                [
                  "object",
                  {
                    "resource_id": "string",
                    "resource_type": "string"
                  }
                ]
              ],
              "source_tag": [
                "list",
                [
                  "object",
                  {
                    "resource_tags": [
                      "map",
                      "string"
                    ],
                    "resource_type_filter": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "stacks": [
                "list",
                [
                  "object",
                  {
                    "stack_id": "string"
                  }
                ]
              ],
              "status": "string",
              "template_scratch_id": "string",
              "template_scratch_type": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_scratch_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRosTemplateScratchesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRosTemplateScratches), &result)
	return &result
}
