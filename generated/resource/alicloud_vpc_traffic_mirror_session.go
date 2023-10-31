package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcTrafficMirrorSession = `{
  "block": {
    "attributes": {
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled": {
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
      "packet_length": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "traffic_mirror_filter_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_mirror_session_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_session_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_source_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "traffic_mirror_target_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_mirror_target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_network_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudVpcTrafficMirrorSessionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcTrafficMirrorSession), &result)
	return &result
}
