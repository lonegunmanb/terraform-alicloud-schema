package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionHoneyPot = `{
  "block": {
    "attributes": {
      "honeypot_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "honeypot_image_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "honeypot_image_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "honeypot_name": {
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
      "node_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preset_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "delete": {
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

func AlicloudThreatDetectionHoneyPotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionHoneyPot), &result)
	return &result
}
