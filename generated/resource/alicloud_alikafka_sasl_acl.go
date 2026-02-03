package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlikafkaSaslAcl = `{
  "block": {
    "attributes": {
      "acl_operation_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acl_operation_types": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acl_permission_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acl_resource_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acl_resource_pattern_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acl_resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "username": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudAlikafkaSaslAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaSaslAcl), &result)
	return &result
}
