package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseInstances = `{
  "block": {
    "attributes": {
      "env_type": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "instance_alias_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_link_name": "string",
              "database_password": "string",
              "database_user": "string",
              "dba_id": "string",
              "dba_nick_name": "string",
              "ddl_online": "number",
              "ecs_instance_id": "string",
              "ecs_region": "string",
              "env_type": "string",
              "export_timeout": "number",
              "host": "string",
              "id": "string",
              "instance_alias": "string",
              "instance_id": "string",
              "instance_name": "string",
              "instance_source": "string",
              "instance_type": "string",
              "port": "number",
              "query_timeout": "number",
              "safe_rule_id": "string",
              "sid": "string",
              "status": "string",
              "use_dsql": "number",
              "vpc_id": "string"
            }
          ]
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "net_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tid": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDmsEnterpriseInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseInstances), &result)
	return &result
}
