package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafDomains = `{
  "block": {
    "attributes": {
      "domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_type": "string",
              "cname": "string",
              "connection_time": "number",
              "domain": "string",
              "domain_name": "string",
              "http2_port": [
                "list",
                "string"
              ],
              "http_port": [
                "list",
                "string"
              ],
              "http_to_user_ip": "string",
              "https_port": [
                "list",
                "string"
              ],
              "https_redirect": "string",
              "id": "string",
              "is_access_product": "string",
              "load_balancing": "string",
              "log_headers": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "read_time": "number",
              "resource_group_id": "string",
              "source_ips": [
                "list",
                "string"
              ],
              "version": "number",
              "write_time": "number"
            }
          ]
        ]
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudWafDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafDomains), &result)
	return &result
}
