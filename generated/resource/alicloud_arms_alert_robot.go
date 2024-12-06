package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsAlertRobot = `{
  "block": {
    "attributes": {
      "alert_robot_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "daily_noc": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "daily_noc_time": {
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
      "robot_addr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "robot_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudArmsAlertRobotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsAlertRobot), &result)
	return &result
}
