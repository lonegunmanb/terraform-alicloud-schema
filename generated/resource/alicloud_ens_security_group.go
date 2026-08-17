package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsSecurityGroup = `{
  "block": {
    "attributes": {
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
      "security_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "permissions": {
        "block": {
          "attributes": {
            "creation_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_cidr_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "direction": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_protocol": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv6_dest_cidr_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv6_source_cidr_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port_range": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "source_cidr_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "computed": true,
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

func AlicloudEnsSecurityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsSecurityGroup), &result)
	return &result
}
