package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdAdConnectorOfficeSite = `{
  "block": {
    "attributes": {
      "ad_connector_office_site_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ad_hostname": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bandwidth": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cen_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_block": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "desktop_access_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_address": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "domain_user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_admin_access": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_internet_access": {
        "computed": true,
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
      "mfa_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "protocol_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "specification": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sub_domain_dns_address": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "sub_domain_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "verify_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudEcdAdConnectorOfficeSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdAdConnectorOfficeSite), &result)
	return &result
}
