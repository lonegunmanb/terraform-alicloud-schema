package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApiGatewayApi = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "force_nonce_check": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "constant_parameters": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "in": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "fc_service_config": {
        "block": {
          "attributes": {
            "arn_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "function_base_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "method": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "only_business_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "qualifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_service_config": {
        "block": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "aone_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type_category": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "method": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_vpc_service_config": {
        "block": {
          "attributes": {
            "aone_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type_category": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_type_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "method": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "vpc_scheme": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mock_service_config": {
        "block": {
          "attributes": {
            "aone_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "result": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "request_config": {
        "block": {
          "attributes": {
            "body_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "method": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "request_parameters": {
        "block": {
          "attributes": {
            "default_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "in": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "in_service": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name_service": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "required": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "system_parameters": {
        "block": {
          "attributes": {
            "in": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name_service": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AlicloudApiGatewayApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApiGatewayApi), &result)
	return &result
}
