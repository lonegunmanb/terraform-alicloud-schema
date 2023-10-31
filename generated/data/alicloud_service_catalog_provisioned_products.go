package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudServiceCatalogProvisionedProducts = `{
  "block": {
    "attributes": {
      "access_level_filter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "products": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "id": "string",
              "last_provisioning_task_id": "string",
              "last_successful_provisioning_task_id": "string",
              "last_task_id": "string",
              "outputs": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "output_key": "string",
                    "output_value": "string"
                  }
                ]
              ],
              "owner_principal_id": "string",
              "owner_principal_type": "string",
              "parameters": [
                "list",
                [
                  "object",
                  {
                    "parameter_key": "string",
                    "parameter_value": "string"
                  }
                ]
              ],
              "portfolio_id": "string",
              "product_id": "string",
              "product_name": "string",
              "product_version_id": "string",
              "product_version_name": "string",
              "provisioned_product_arn": "string",
              "provisioned_product_id": "string",
              "provisioned_product_name": "string",
              "provisioned_product_type": "string",
              "stack_id": "string",
              "stack_region_id": "string",
              "status": "string",
              "status_message": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "provisioned_products": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "id": "string",
              "last_provisioning_task_id": "string",
              "last_successful_provisioning_task_id": "string",
              "last_task_id": "string",
              "outputs": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "output_key": "string",
                    "output_value": "string"
                  }
                ]
              ],
              "owner_principal_id": "string",
              "owner_principal_type": "string",
              "parameters": [
                "list",
                [
                  "object",
                  {
                    "parameter_key": "string",
                    "parameter_value": "string"
                  }
                ]
              ],
              "portfolio_id": "string",
              "product_id": "string",
              "product_name": "string",
              "product_version_id": "string",
              "product_version_name": "string",
              "provisioned_product_arn": "string",
              "provisioned_product_id": "string",
              "provisioned_product_name": "string",
              "provisioned_product_type": "string",
              "stack_id": "string",
              "stack_region_id": "string",
              "status": "string",
              "status_message": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "sort_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort_order": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudServiceCatalogProvisionedProductsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudServiceCatalogProvisionedProducts), &result)
	return &result
}
