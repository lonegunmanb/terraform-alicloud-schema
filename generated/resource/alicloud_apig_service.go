package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigService = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "dns_servers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "express_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "healthy_panic_threshold": {
        "computed": true,
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
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "outlier_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ports": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "port": "number",
              "protocol": "string"
            }
          ]
        ]
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "qualifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_detail_error_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_detail_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unhealthy_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "update_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "health_check_config": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "expected_statuses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "healthy_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_host": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "unhealthy_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "outlier_detection_config": {
        "block": {
          "attributes": {
            "base_ejection_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "failure_percentage_minimum_hosts": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "failure_percentage_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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

func AlicloudApigServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigService), &result)
	return &result
}
