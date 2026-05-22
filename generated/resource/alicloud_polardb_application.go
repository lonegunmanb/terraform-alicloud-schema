package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbApplication = `{
  "block": {
    "attributes": {
      "ai_db_cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "application_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "architecture": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "component_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_id": {
        "description_kind": "plain",
        "optional": true,
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
      "model_api": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_api_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "model_base_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_from": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "modify_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pay_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ip_array_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ip_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "upgrade_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "components": {
        "block": {
          "attributes": {
            "component_class": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "component_replica": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "component_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "parameters": {
        "block": {
          "attributes": {
            "parameter_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_value": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
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

func AlicloudPolardbApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbApplication), &result)
	return &result
}
