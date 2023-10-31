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
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlikafkaSaslAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaSaslAcl), &result)
	return &result
}
