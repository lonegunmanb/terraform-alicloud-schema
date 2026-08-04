package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServiceInstances = `{
  "block": {
    "attributes": {
      "brand": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_type": {
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
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_reissue": "string",
              "average_waiting_time": "string",
              "brand": "string",
              "certificate_type": "string",
              "city": "string",
              "company_id": "string",
              "contact_id_list": [
                "list",
                "number"
              ],
              "country_code": "string",
              "csr": "string",
              "domain": "string",
              "full_domain_count": "number",
              "generate_csr_method": "string",
              "id": "string",
              "instance_end_time": "number",
              "instance_id": "string",
              "instance_name": "string",
              "instance_start_time": "number",
              "instance_type": "string",
              "key_algorithm": "string",
              "order_end_time": "number",
              "order_start_time": "number",
              "province": "string",
              "resource_group_id": "string",
              "spec": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "upgrade_status": "string",
              "validation_method": "string",
              "wildcard_domain_count": "number"
            }
          ]
        ]
      },
      "keyword": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSslCertificatesServiceInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServiceInstances), &result)
	return &result
}
