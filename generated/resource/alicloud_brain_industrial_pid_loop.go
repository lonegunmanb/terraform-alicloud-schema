package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudBrainIndustrialPidLoop = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pid_loop_configuration": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pid_loop_dcs_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pid_loop_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pid_loop_is_crucial": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "pid_loop_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pid_loop_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pid_project_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudBrainIndustrialPidLoopSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudBrainIndustrialPidLoop), &result)
	return &result
}
