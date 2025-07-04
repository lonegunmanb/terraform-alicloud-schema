package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMaxcomputeProject = `{
  "block": {
    "attributes": {
      "comment": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_quota": {
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
      "is_logical": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "three_tier_model": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "ip_white_list": {
        "block": {
          "attributes": {
            "ip_list": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_ip_list": {
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
      "properties": {
        "block": {
          "attributes": {
            "allow_full_scan": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_decimal2": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_dr": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retention_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sql_metering_max": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timezone": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type_system": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "encryption": {
              "block": {
                "attributes": {
                  "algorithm": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "key": {
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
            "table_lifecycle": {
              "block": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
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
      "security_properties": {
        "block": {
          "attributes": {
            "enable_download_privilege": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "label_security": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "object_creator_has_access_permission": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "object_creator_has_grant_permission": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "using_acl": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "using_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "project_protection": {
              "block": {
                "attributes": {
                  "exception_policy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protected": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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

func AlicloudMaxcomputeProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMaxcomputeProject), &result)
	return &result
}
