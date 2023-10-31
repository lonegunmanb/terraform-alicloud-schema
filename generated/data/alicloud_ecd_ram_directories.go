package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdRamDirectories = `{
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
              "ad_connectors": [
                "list",
                [
                  "object",
                  {
                    "ad_connector_address": "string",
                    "connector_status": "string",
                    "network_interface_id": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "create_time": "string",
              "custom_security_group_id": "string",
              "desktop_access_type": "string",
              "desktop_vpc_endpoint": "string",
              "directory_type": "string",
              "dns_address": [
                "list",
                "string"
              ],
              "dns_user_name": "string",
              "domain_name": "string",
              "domain_password": "string",
              "domain_user_name": "string",
              "enable_admin_access": "bool",
              "enable_cross_desktop_access": "bool",
              "enable_internet_access": "bool",
              "file_system_ids": [
                "list",
                "string"
              ],
              "id": "string",
              "logs": [
                "list",
                [
                  "object",
                  {
                    "level": "string",
                    "message": "string",
                    "step": "string",
                    "time_stamp": "string"
                  }
                ]
              ],
              "mfa_enabled": "bool",
              "ram_directory_id": "string",
              "ram_directory_name": "string",
              "sso_enabled": "bool",
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

func AlicloudEcdRamDirectoriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdRamDirectories), &result)
	return &result
}
