package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEhpcJobTemplates = `{
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "array_request": "string",
              "clock_time": "string",
              "command_line": "string",
              "gpu": "number",
              "id": "string",
              "job_template_id": "string",
              "job_template_name": "string",
              "mem": "string",
              "node": "number",
              "package_path": "string",
              "priority": "number",
              "queue": "string",
              "re_runable": "bool",
              "runas_user": "string",
              "stderr_redirect_path": "string",
              "stdout_redirect_path": "string",
              "task": "number",
              "thread": "number",
              "variables": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEhpcJobTemplatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEhpcJobTemplates), &result)
	return &result
}
