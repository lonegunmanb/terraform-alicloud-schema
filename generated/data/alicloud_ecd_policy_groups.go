package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdPolicyGroups = `{
  "block": {
    "attributes": {
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authorize_access_policy_rules": [
                "list",
                [
                  "object",
                  {
                    "cidr_ip": "string",
                    "description": "string"
                  }
                ]
              ],
              "authorize_security_policy_rules": [
                "list",
                [
                  "object",
                  {
                    "cidr_ip": "string",
                    "description": "string",
                    "ip_protocol": "string",
                    "policy": "string",
                    "port_range": "string",
                    "priority": "string",
                    "type": "string"
                  }
                ]
              ],
              "camera_redirect": "string",
              "clipboard": "string",
              "domain_list": "string",
              "eds_count": "number",
              "html_access": "string",
              "html_file_transfer": "string",
              "id": "string",
              "local_drive": "string",
              "policy_group_id": "string",
              "policy_group_name": "string",
              "policy_group_type": "string",
              "recording": "string",
              "recording_end_time": "string",
              "recording_fps": "number",
              "recording_start_time": "string",
              "status": "string",
              "usb_redirect": "string",
              "visual_quality": "string",
              "watermark": "string",
              "watermark_transparency": "string",
              "watermark_type": "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcdPolicyGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdPolicyGroups), &result)
	return &result
}
