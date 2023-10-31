package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlikafkaSaslAcls = `{
  "block": {
    "attributes": {
      "acl_resource_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acl_resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl_operation_type": "string",
              "acl_resource_name": "string",
              "acl_resource_pattern_type": "string",
              "acl_resource_type": "string",
              "host": "string",
              "username": "string"
            }
          ]
        ]
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlikafkaSaslAclsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaSaslAcls), &result)
	return &result
}
