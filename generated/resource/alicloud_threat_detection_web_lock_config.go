package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionWebLockConfig = `{
  "block": {
    "attributes": {
      "defence_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dir": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "exclusive_dir": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclusive_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclusive_file_type": {
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
      "inclusive_file_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_backup_dir": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uuid": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionWebLockConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionWebLockConfig), &result)
	return &result
}
