package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsInvocations = `{
  "block": {
    "attributes": {
      "command_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_encoding": {
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
      "invocations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "command_content": "string",
              "command_id": "string",
              "command_name": "string",
              "command_type": "string",
              "create_time": "string",
              "frequency": "string",
              "id": "string",
              "invocation_id": "string",
              "invocation_status": "string",
              "invoke_instances": [
                "list",
                [
                  "object",
                  {
                    "creation_time": "string",
                    "dropped": "number",
                    "error_code": "string",
                    "error_info": "string",
                    "exit_code": "number",
                    "finish_time": "string",
                    "instance_id": "string",
                    "instance_invoke_status": "string",
                    "invocation_status": "string",
                    "output": "string",
                    "repeats": "number",
                    "start_time": "string",
                    "stop_time": "string",
                    "timed": "bool",
                    "update_time": "string"
                  }
                ]
              ],
              "invoke_status": "string",
              "parameters": "string",
              "repeat_mode": "string",
              "timed": "bool",
              "username": "string"
            }
          ]
        ]
      },
      "invoke_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsInvocationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsInvocations), &result)
	return &result
}
