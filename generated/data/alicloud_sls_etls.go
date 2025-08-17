package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsEtls = `{
  "block": {
    "attributes": {
      "etls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "configuration": [
                "list",
                [
                  "object",
                  {
                    "from_time": "number",
                    "lang": "string",
                    "logstore": "string",
                    "parameters": [
                      "map",
                      "string"
                    ],
                    "role_arn": "string",
                    "script": "string",
                    "sink": [
                      "list",
                      [
                        "object",
                        {
                          "datasets": [
                            "list",
                            "string"
                          ],
                          "endpoint": "string",
                          "logstore": "string",
                          "name": "string",
                          "project": "string",
                          "role_arn": "string"
                        }
                      ]
                    ],
                    "to_time": "number"
                  }
                ]
              ],
              "create_time": "number",
              "description": "string",
              "display_name": "string",
              "id": "string",
              "job_name": "string",
              "last_modified_time": "number",
              "schedule_id": "string",
              "status": "string"
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
      "logstore": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "offset": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlsEtlsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsEtls), &result)
	return &result
}
