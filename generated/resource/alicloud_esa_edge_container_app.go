package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaEdgeContainerApp = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "edge_container_app_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "health_check_fail_times": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_host": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_http_code": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_method": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_succ_times": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_uri": {
        "computed": true,
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
      "remarks": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_port": {
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

func AlicloudEsaEdgeContainerAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaEdgeContainerApp), &result)
	return &result
}
