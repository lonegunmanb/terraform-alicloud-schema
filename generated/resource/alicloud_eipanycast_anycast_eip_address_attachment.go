package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEipanycastAnycastEipAddressAttachment = `{
  "block": {
    "attributes": {
      "anycast_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "association_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bind_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bind_instance_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bind_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bind_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "pop_locations": {
        "block": {
          "attributes": {
            "pop_location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AlicloudEipanycastAnycastEipAddressAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEipanycastAnycastEipAddressAttachment), &result)
	return &result
}
