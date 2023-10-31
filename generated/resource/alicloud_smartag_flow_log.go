package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSmartagFlowLog = `{
  "block": {
    "attributes": {
      "active_aging": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_log_name": {
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
      "inactive_aging": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "logstore_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "netflow_server_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "netflow_server_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "netflow_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sls_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
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

func AlicloudSmartagFlowLogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSmartagFlowLog), &result)
	return &result
}
