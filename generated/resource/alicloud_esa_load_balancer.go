package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaLoadBalancer = `{
  "block": {
    "attributes": {
      "default_pools": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "number"
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fallback_pool": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "load_balancer_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_pools": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "session_affinity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "steering_policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sub_region_pools": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ttl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "adaptive_routing": {
        "block": {
          "attributes": {
            "failover_across_pools": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "monitor": {
        "block": {
          "attributes": {
            "consecutive_down": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "consecutive_up": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "expected_codes": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "follow_redirects": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "header": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "method": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monitoring_region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "random_steering": {
        "block": {
          "attributes": {
            "default_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pool_weights": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rules": {
        "block": {
          "attributes": {
            "overrides": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sequence": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "terminates": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "fixed_response": {
              "block": {
                "attributes": {
                  "content_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_body": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
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

func AlicloudEsaLoadBalancerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaLoadBalancer), &result)
	return &result
}
