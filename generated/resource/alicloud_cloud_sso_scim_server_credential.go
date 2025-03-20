package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoScimServerCredential = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "credential_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "credential_secret_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credential_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expire_time": {
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
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudSsoScimServerCredentialSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoScimServerCredential), &result)
	return &result
}
