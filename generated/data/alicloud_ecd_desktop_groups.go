package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdDesktopGroups = `{
  "block": {
    "attributes": {
      "desktop_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desktop_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_auto_setup": "number",
              "allow_buffer_count": "number",
              "bundle_id": "string",
              "comments": "string",
              "cpu": "number",
              "create_time": "string",
              "creator": "string",
              "data_disk_category": "string",
              "data_disk_size": "string",
              "desktop_group_id": "string",
              "desktop_group_name": "string",
              "directory_id": "string",
              "directory_type": "string",
              "end_user_count": "number",
              "end_user_ids": [
                "list",
                "string"
              ],
              "expired_time": "string",
              "gpu_count": "number",
              "gpu_spec": "string",
              "id": "string",
              "keep_duration": "number",
              "max_desktops_count": "number",
              "memory": "number",
              "min_desktops_count": "number",
              "office_site_id": "string",
              "office_site_name": "string",
              "office_site_type": "string",
              "own_bundle_name": "string",
              "pay_type": "string",
              "policy_group_id": "string",
              "policy_group_name": "string",
              "res_type": "number",
              "system_disk_category": "string",
              "system_disk_size": "number"
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
      "office_site_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcdDesktopGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdDesktopGroups), &result)
	return &result
}
