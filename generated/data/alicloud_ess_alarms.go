package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssAlarms = `{
  "block": {
    "attributes": {
      "alarms": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alarm_actions": [
                "set",
                "string"
              ],
              "cloud_monitor_group_id": "number",
              "comparison_operator": "string",
              "description": "string",
              "dimensions": [
                "map",
                "string"
              ],
              "enable": "bool",
              "evaluation_count": "number",
              "id": "string",
              "metric_name": "string",
              "metric_type": "string",
              "name": "string",
              "period": "number",
              "scaling_group_id": "string",
              "state": "string",
              "statistics": "string",
              "threshold": "string"
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
      "metric_type": {
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
      "scaling_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssAlarmsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssAlarms), &result)
	return &result
}
