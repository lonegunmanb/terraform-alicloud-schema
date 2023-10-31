package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdCommands = `{
  "block": {
    "attributes": {
      "command_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "commands": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "command_content": "string",
              "command_type": "string",
              "create_time": "string",
              "id": "string",
              "invoke_desktops": [
                "list",
                [
                  "object",
                  {
                    "desktop_id": "string",
                    "dropped": "number",
                    "error_code": "string",
                    "error_info": "string",
                    "exit_code": "string",
                    "finish_time": "string",
                    "invocation_status": "string",
                    "output": "string",
                    "repeats": "number",
                    "start_time": "string",
                    "stop_time": "string"
                  }
                ]
              ],
              "invoke_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "content_encoding": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desktop_id": {
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

func AlicloudEcdCommandsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdCommands), &result)
	return &result
}
