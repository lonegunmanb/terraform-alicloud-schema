package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenInstanceAttachment = `{
  "block": {
    "attributes": {
      "cen_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "child_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "child_instance_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "child_instance_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "child_instance_type": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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

func AlicloudCenInstanceAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenInstanceAttachment), &result)
	return &result
}
