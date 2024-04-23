package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudInstanceTypes = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu_core_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "eni_amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "gpu_amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "gpu_spec": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type_family": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zones": [
                "list",
                "string"
              ],
              "burstable_instance": [
                "map",
                "string"
              ],
              "cpu_core_count": "number",
              "eni_amount": "number",
              "family": "string",
              "gpu": [
                "map",
                "string"
              ],
              "id": "string",
              "local_storage": [
                "map",
                "string"
              ],
              "memory_size": "number",
              "nvme_support": "string",
              "price": "string"
            }
          ]
        ]
      },
      "is_outdated": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kubernetes_node_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_eni_ipv6_address_quantity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sorted_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudInstanceTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudInstanceTypes), &result)
	return &result
}
