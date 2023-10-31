package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrChains = `{
  "block": {
    "attributes": {
      "chains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "chain_config": [
                "list",
                [
                  "object",
                  {
                    "nodes": [
                      "list",
                      [
                        "object",
                        {
                          "enable": "bool",
                          "node_config": [
                            "list",
                            [
                              "object",
                              {
                                "deny_policy": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "action": "string",
                                      "issue_count": "string",
                                      "issue_level": "string",
                                      "logic": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "node_name": "string"
                        }
                      ]
                    ],
                    "routers": [
                      "list",
                      [
                        "object",
                        {
                          "from": [
                            "list",
                            [
                              "object",
                              {
                                "node_name": "string"
                              }
                            ]
                          ],
                          "to": [
                            "list",
                            [
                              "object",
                              {
                                "node_name": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "chain_id": "string",
              "chain_name": "string",
              "create_time": "string",
              "description": "string",
              "id": "string",
              "instance_id": "string",
              "modified_time": "string",
              "scope_id": "string",
              "scope_type": "string"
            }
          ]
        ]
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
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
      "repo_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repo_namespace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCrChainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrChains), &result)
	return &result
}
