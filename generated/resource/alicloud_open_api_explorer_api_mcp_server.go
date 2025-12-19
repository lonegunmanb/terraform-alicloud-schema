package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOpenApiExplorerApiMcpServer = `{
  "block": {
    "attributes": {
      "assume_role_extra_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "assume_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_assume_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_custom_vpc_whitelist": {
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
      "instructions": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oauth_client_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_tools": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_whitelists": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "additional_api_descriptions": {
        "block": {
          "attributes": {
            "api_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "api_override_json": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "api_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_output_schema": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "execute_cli_command": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "product": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "const_parameters": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
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
      },
      "apis": {
        "block": {
          "attributes": {
            "api_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "product": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "selectors": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "prompts": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "arguments": {
              "block": {
                "attributes": {
                  "description": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "required": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "terraform_tools": {
        "block": {
          "attributes": {
            "async": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destroy_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOpenApiExplorerApiMcpServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOpenApiExplorerApiMcpServer), &result)
	return &result
}
