package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionCheckStructures = `{
  "block": {
    "attributes": {
      "current_page": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "structures": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "standard_type": "string",
              "standards": [
                "list",
                [
                  "object",
                  {
                    "id": "number",
                    "requirements": [
                      "list",
                      [
                        "object",
                        {
                          "id": "number",
                          "sections": [
                            "list",
                            [
                              "object",
                              {
                                "id": "number",
                                "show_name": "string"
                              }
                            ]
                          ],
                          "show_name": "string",
                          "total_check_count": "number"
                        }
                      ]
                    ],
                    "show_name": "string",
                    "type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "task_sources": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionCheckStructuresSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionCheckStructures), &result)
	return &result
}
