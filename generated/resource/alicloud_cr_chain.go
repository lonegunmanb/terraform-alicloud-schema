package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCrChain = `{
  "block": {
    "attributes": {
      "chain_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "chain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
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
    "block_types": {
      "chain_config": {
        "block": {
          "block_types": {
            "nodes": {
              "block": {
                "attributes": {
                  "enable": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "node_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "node_config": {
                    "block": {
                      "block_types": {
                        "deny_policy": {
                          "block": {
                            "attributes": {
                              "action": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "issue_count": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "issue_level": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "logic": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "routers": {
              "block": {
                "block_types": {
                  "from": {
                    "block": {
                      "attributes": {
                        "node_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "to": {
                    "block": {
                      "attributes": {
                        "node_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCrChainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCrChain), &result)
	return &result
}
