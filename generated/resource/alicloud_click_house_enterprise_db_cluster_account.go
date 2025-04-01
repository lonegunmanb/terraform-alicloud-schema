package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudClickHouseEnterpriseDbClusterAccount = `{
  "block": {
    "attributes": {
      "account": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "account_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
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
      "password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      }
    },
    "block_types": {
      "dml_auth_setting": {
        "block": {
          "attributes": {
            "allow_databases": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_dictionaries": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ddl_authority": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "dml_authority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudClickHouseEnterpriseDbClusterAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudClickHouseEnterpriseDbClusterAccount), &result)
	return &result
}
