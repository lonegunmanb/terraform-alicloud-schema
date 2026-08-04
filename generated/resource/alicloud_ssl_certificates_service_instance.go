package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServiceInstance = `{
  "block": {
    "attributes": {
      "auto_reissue": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "average_waiting_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "brand": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "city": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "company_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "contact_id_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "country_code": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "csr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "full_domain_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "generate_csr_method": {
        "computed": true,
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
      "instance_end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instance_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_algorithm": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "order_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "pricing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "product_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "province": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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
      "upgrade_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "validation_method": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wildcard_domain_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "parameter": {
        "block": {
          "attributes": {
            "code": {
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

func AlicloudSslCertificatesServiceInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServiceInstance), &result)
	return &result
}
