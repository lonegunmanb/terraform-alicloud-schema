package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaRecord = `{
  "block": {
    "attributes": {
      "biz_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
      "host_policy": {
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
      "proxied": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "record_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "record_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "auth_conf": {
        "block": {
          "attributes": {
            "access_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "auth_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "version": {
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
      "data": {
        "block": {
          "attributes": {
            "algorithm": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "certificate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fingerprint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flag": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "key_tag": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "matching_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "selector": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tag": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "usage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AlicloudEsaRecordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaRecord), &result)
	return &result
}
