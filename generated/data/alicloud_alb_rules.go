package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbRules = `{
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
      "listener_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "load_balancer_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "rule_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "listener_id": "string",
              "load_balancer_id": "string",
              "priority": "number",
              "rule_actions": [
                "list",
                [
                  "object",
                  {
                    "fixed_response_config": [
                      "list",
                      [
                        "object",
                        {
                          "content": "string",
                          "content_type": "string",
                          "http_code": "string"
                        }
                      ]
                    ],
                    "forward_group_config": [
                      "list",
                      [
                        "object",
                        {
                          "server_group_tuples": [
                            "list",
                            [
                              "object",
                              {
                                "server_group_id": "string",
                                "weight": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "insert_header_config": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string",
                          "value_type": "string"
                        }
                      ]
                    ],
                    "order": "number",
                    "redirect_config": [
                      "list",
                      [
                        "object",
                        {
                          "host": "string",
                          "http_code": "string",
                          "path": "string",
                          "port": "string",
                          "protocol": "string",
                          "query": "string"
                        }
                      ]
                    ],
                    "rewrite_config": [
                      "list",
                      [
                        "object",
                        {
                          "host": "string",
                          "path": "string",
                          "query": "string"
                        }
                      ]
                    ],
                    "traffic_limit_config": [
                      "list",
                      [
                        "object",
                        {
                          "qps": "number"
                        }
                      ]
                    ],
                    "traffic_mirror_config": [
                      "list",
                      [
                        "object",
                        {
                          "mirror_group_config": [
                            "list",
                            [
                              "object",
                              {
                                "server_group_tuples": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "server_group_id": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "target_type": "string"
                        }
                      ]
                    ],
                    "type": "string"
                  }
                ]
              ],
              "rule_conditions": [
                "list",
                [
                  "object",
                  {
                    "cookie_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "value": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "header_config": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "host_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "method_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "path_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "query_string_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "value": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "source_ip_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "type": "string"
                  }
                ]
              ],
              "rule_id": "string",
              "rule_name": "string",
              "status": "string"
            }
          ]
        ]
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

func AlicloudAlbRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbRules), &result)
	return &result
}
