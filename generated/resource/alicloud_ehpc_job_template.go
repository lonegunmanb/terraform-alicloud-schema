package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEhpcJobTemplate = `{
  "block": {
    "attributes": {
      "array_request": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "clock_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "command_line": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mem": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "package_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "queue": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "re_runable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "runas_user": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stderr_redirect_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stdout_redirect_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "thread": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "variables": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEhpcJobTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEhpcJobTemplate), &result)
	return &result
}
