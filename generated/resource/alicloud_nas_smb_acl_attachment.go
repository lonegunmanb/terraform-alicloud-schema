package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNasSmbAclAttachment = `{
  "block": {
    "attributes": {
      "auth_method": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_anonymous_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encrypt_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "file_system_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_dir_path": {
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
      "keytab": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "keytab_md5": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reject_unencrypted_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "super_admin_sid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudNasSmbAclAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNasSmbAclAttachment), &result)
	return &result
}
