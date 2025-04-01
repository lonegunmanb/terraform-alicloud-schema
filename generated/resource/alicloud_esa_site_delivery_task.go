package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaSiteDeliveryTask = `{
  "block": {
    "attributes": {
      "business_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_center": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delivery_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "discard_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "field_name": {
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
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "http_delivery": {
        "block": {
          "attributes": {
            "compress": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dest_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "header_param": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "log_body_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_body_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_batch_mb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_batch_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "query_param": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "standard_auth_on": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "transform_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "standard_auth_param": {
              "block": {
                "attributes": {
                  "expired_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "private_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "url_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "kafka_delivery": {
        "block": {
          "attributes": {
            "balancer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "brokers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "compress": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "machanism_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_auth": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "user_name": {
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
      "oss_delivery": {
        "block": {
          "attributes": {
            "aliuid": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
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
      "s3_delivery": {
        "block": {
          "attributes": {
            "access_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_cmpt": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "secret_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_side_encryption": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vertify_type": {
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
      "sls_delivery": {
        "block": {
          "attributes": {
            "sls_log_store": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sls_project": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sls_region": {
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

func AlicloudEsaSiteDeliveryTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaSiteDeliveryTask), &result)
	return &result
}
