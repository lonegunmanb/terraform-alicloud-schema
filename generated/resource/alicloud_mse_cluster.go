package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMseCluster = `{
  "block": {
    "attributes": {
      "acl_entry_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "app_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_alias_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_specification": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_type": {
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
      "instance_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "mse_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "net_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_slb_specification": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pub_network_flow": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pub_slb_specification": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_pars": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudMseClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMseCluster), &result)
	return &result
}
