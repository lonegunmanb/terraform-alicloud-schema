package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudClickHouseEnterpriseDbClusterComputingGroup = `{
  "block": {
    "attributes": {
      "computing_group_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "computing_group_endpoint_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "computing_group_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "computing_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "computing_group_public_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "computing_group_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_id": {
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
      "is_readonly": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "node_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "node_scale_max": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "node_scale_min": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudClickHouseEnterpriseDbClusterComputingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudClickHouseEnterpriseDbClusterComputingGroup), &result)
	return &result
}
