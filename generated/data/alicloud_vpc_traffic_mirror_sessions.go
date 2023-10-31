package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcTrafficMirrorSessions = `{
  "block": {
    "attributes": {
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sessions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "id": "string",
              "packet_length": "number",
              "priority": "number",
              "status": "string",
              "traffic_mirror_filter_id": "string",
              "traffic_mirror_session_business_status": "string",
              "traffic_mirror_session_description": "string",
              "traffic_mirror_session_id": "string",
              "traffic_mirror_session_name": "string",
              "traffic_mirror_source_ids": [
                "list",
                "string"
              ],
              "traffic_mirror_target_id": "string",
              "traffic_mirror_target_type": "string",
              "virtual_network_id": "number"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_filter_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_session_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_source_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "traffic_mirror_target_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcTrafficMirrorSessionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcTrafficMirrorSessions), &result)
	return &result
}
