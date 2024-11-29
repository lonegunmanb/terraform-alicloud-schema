package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPaiWorkspaceDatasetversion = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dataset_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "options": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "property": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uri": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "labels": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPaiWorkspaceDatasetversionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPaiWorkspaceDatasetversion), &result)
	return &result
}
