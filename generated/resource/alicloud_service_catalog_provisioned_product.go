package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudServiceCatalogProvisionedProduct = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_provisioning_task_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_successful_provisioning_task_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_task_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "outputs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "description": "string",
              "output_key": "string",
              "output_value": "string"
            }
          ]
        ]
      },
      "owner_principal_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner_principal_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "portfolio_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_version_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_version_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_product_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_product_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_product_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioned_product_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
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
      }
    },
    "block_types": {
      "parameters": {
        "block": {
          "attributes": {
            "parameter_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_value": {
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

func AlicloudServiceCatalogProvisionedProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudServiceCatalogProvisionedProduct), &result)
	return &result
}
