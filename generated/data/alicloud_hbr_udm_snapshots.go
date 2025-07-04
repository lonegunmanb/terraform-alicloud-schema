package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrUdmSnapshots = `{
  "block": {
    "attributes": {
      "disk_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_time": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshots": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "disk_id": "string",
              "id": "string",
              "instance_id": "string",
              "job_id": "string",
              "source_type": "string",
              "start_time": "number",
              "udm_snapshot_id": "string"
            }
          ]
        ]
      },
      "source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbrUdmSnapshotsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrUdmSnapshots), &result)
	return &result
}
