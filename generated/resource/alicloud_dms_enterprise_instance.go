package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseInstance = `{
  "block": {
    "attributes": {
      "data_link_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "database_user": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dba_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dba_nick_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dba_uid": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "ddl_online": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ecs_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "env_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "export_timeout": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "host": {
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
      "instance_alias": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "query_timeout": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "safe_rule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "safe_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_test": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tid": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "use_dsql": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vpc_id": {
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

func AlicloudDmsEnterpriseInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseInstance), &result)
	return &result
}
