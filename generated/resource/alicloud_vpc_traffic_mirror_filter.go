package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcTrafficMirrorFilter = `{
  "block": {
    "attributes": {
      "dry_run": {
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
      "traffic_mirror_filter_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_filter_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "egress_rules": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "destination_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_port_range": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "traffic_mirror_filter_rule_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "ingress_rules": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "destination_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_port_range": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_cidr_block": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "traffic_mirror_filter_rule_status": {
              "computed": true,
              "description_kind": "plain",
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

func AlicloudVpcTrafficMirrorFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcTrafficMirrorFilter), &result)
	return &result
}
