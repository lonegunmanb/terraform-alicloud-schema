package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoDirectory = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_global_access_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_name": {
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
      "mfa_authentication_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scim_synchronization_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "login_preference": {
        "block": {
          "attributes": {
            "allow_user_to_get_credentials": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "login_network_masks": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mfa_authentication_setting_info": {
        "block": {
          "attributes": {
            "mfa_authentication_advance_settings": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "operation_for_risk_login": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "password_policy": {
        "block": {
          "attributes": {
            "hard_expire": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "max_login_attempts": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_password_age": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_password_length": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_password_different_chars": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_password_length": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "password_not_contain_username": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "password_reuse_prevention": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "require_lower_case_chars": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "require_numbers": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "require_symbols": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "require_upper_case_chars": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "saml_identity_provider_configuration": {
        "block": {
          "attributes": {
            "binding_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "certificate_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "create_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "encoded_metadata_document": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "entity_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "login_url": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sso_status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "want_request_signed": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "saml_service_provider": {
        "block": {
          "attributes": {
            "acs_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "authn_sign_algo": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "certificate_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encoded_metadata_document": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "entity_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "support_encrypted_assertion": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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
      },
      "user_provisioning_configuration": {
        "block": {
          "attributes": {
            "default_landing_page": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "session_duration": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudSsoDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoDirectory), &result)
	return &result
}
