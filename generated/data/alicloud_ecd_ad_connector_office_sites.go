package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdAdConnectorOfficeSites = `{
  "block": {
    "attributes": {
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
      "sites": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ad_connector_office_site_name": "string",
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
              "bandwidth": "number",
              "cen_id": "string",
              "cidr_block": "string",
              "create_time": "string",
              "custom_security_group_id": "string",
              "desktop_access_type": "string",
              "desktop_vpc_endpoint": "string",
              "dns_address": [
                "list",
                "string"
              ],
              "dns_user_name": "string",
              "domain_name": "string",
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
              "network_package_id": "string",
              "office_site_id": "string",
              "office_site_type": "string",
              "sso_enabled": "bool",
              "status": "string",
              "sub_domain_dns_address": [
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

func AlicloudEcdAdConnectorOfficeSitesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdAdConnectorOfficeSites), &result)
	return &result
}
