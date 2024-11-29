package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenFlowlogs = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_log_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_log_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_log_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flowlogs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cen_id": "string",
              "create_time": "string",
              "description": "string",
              "flow_log_id": "string",
              "flow_log_name": "string",
              "flow_log_version": "string",
              "id": "string",
              "interval": "number",
              "log_format_string": "string",
              "log_store_name": "string",
              "project_name": "string",
              "record_total": "string",
              "region_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "transit_router_attachment_id": "string",
              "transit_router_id": "string"
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
      "interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "log_store_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenFlowlogsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenFlowlogs), &result)
	return &result
}
