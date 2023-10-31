package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsAutoSnapshotPolicyAttachment = `{
  "block": {
    "attributes": {
      "auto_snapshot_policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
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

func AlicloudEcsAutoSnapshotPolicyAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsAutoSnapshotPolicyAttachment), &result)
	return &result
}
