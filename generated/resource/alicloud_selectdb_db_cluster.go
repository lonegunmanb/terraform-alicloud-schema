package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSelectdbDbCluster = `{
  "block": {
    "attributes": {
      "cache_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "desired_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
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
      "memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "param_change_logs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_id": "number",
              "gmt_created": "string",
              "gmt_modified": "string",
              "is_applied": "bool",
              "name": "string",
              "new_value": "string",
              "old_value": "string"
            }
          ]
        ]
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "desired_params": {
        "block": {
          "attributes": {
            "name": {
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

func AlicloudSelectdbDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSelectdbDbCluster), &result)
	return &result
}
