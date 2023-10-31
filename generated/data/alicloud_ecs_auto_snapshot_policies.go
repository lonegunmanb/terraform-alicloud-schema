package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsAutoSnapshotPolicies = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_snapshot_policy_id": "string",
              "copied_snapshots_retention_days": "number",
              "disk_nums": "number",
              "enable_cross_region_copy": "bool",
              "id": "string",
              "name": "string",
              "repeat_weekdays": [
                "list",
                "string"
              ],
              "retention_days": "number",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "target_copy_regions": [
                "list",
                "string"
              ],
              "time_points": [
                "list",
                "string"
              ],
              "volume_nums": "number"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsAutoSnapshotPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsAutoSnapshotPolicies), &result)
	return &result
}
