package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdAdConnectorDirectories = `{
  "block": {
    "attributes": {
      "directories": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ad_connector_directory_id": "string",
              "ad_connectors": [
                "list",
                [
                  "object",
                  {
                    "ad_connector_address": "string",
                    "connector_status": "string",
                    "network_interface_id": "string",
                    "specification": "string",
                    "trust_key": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "create_time": "string",
              "custom_security_group_id": "string",
              "directory_name": "string",
              "directory_type": "string",
              "dns_address": [
                "list",
                "string"
              ],
              "dns_user_name": "string",
              "domain_name": "string",
              "domain_user_name": "string",
              "enable_admin_access": "bool",
              "id": "string",
              "mfa_enabled": "bool",
              "status": "string",
              "sub_dns_address": [
                "list",
                "string"
              ],
              "sub_domain_name": "string",
              "trust_password": "string",
              "vpc_id": "string",
              "vswitch_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
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

func AlicloudEcdAdConnectorDirectoriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdAdConnectorDirectories), &result)
	return &result
}
