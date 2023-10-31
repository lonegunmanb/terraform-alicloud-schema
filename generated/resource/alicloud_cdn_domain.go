package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCdnDomain = `{
  "block": {
    "attributes": {
      "block_ips": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cdn_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_name": {
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
      "optimize_enable": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_compress_enable": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "range_enable": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_port": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "source_type": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sources": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "video_seek_enable": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "auth_config": {
        "block": {
          "attributes": {
            "auth_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "master_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "slave_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "cache_config": {
        "block": {
          "attributes": {
            "cache_content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cache_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "cache_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ttl": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "certificate_config": {
        "block": {
          "attributes": {
            "private_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "server_certificate": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "server_certificate_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_header_config": {
        "block": {
          "attributes": {
            "header_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "header_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "header_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "set"
      },
      "page_404_config": {
        "block": {
          "attributes": {
            "custom_page_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "error_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "page_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "parameter_filter_config": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hash_key_args": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "refer_config": {
        "block": {
          "attributes": {
            "allow_empty": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "refer_list": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "refer_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCdnDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCdnDomain), &result)
	return &result
}
