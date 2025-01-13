package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudBastionhostInstance = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_public_access": {
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
      "license_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "plan_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_white_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_period_unit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renewal_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "storage": {
        "description_kind": "plain",
        "required": true,
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
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "ad_auth_server": {
        "block": {
          "attributes": {
            "account": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "base_dn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "domain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "email_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_ssl": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "mobile_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "server": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "standby_server": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ldap_auth_server": {
        "block": {
          "attributes": {
            "account": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "base_dn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "email_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_ssl": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "login_name_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mobile_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name_mapping": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "server": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "standby_server": {
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

func AlicloudBastionhostInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudBastionhostInstance), &result)
	return &result
}
