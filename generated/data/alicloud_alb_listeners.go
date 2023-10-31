package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbListeners = `{
  "block": {
    "attributes": {
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
      "listener_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "listener_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listeners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_log_record_customized_headers_enabled": "bool",
              "access_log_tracing_config": [
                "list",
                [
                  "object",
                  {
                    "tracing_enabled": "bool",
                    "tracing_sample": "number",
                    "tracing_type": "string"
                  }
                ]
              ],
              "acl_config": [
                "list",
                [
                  "object",
                  {
                    "acl_relations": [
                      "list",
                      [
                        "object",
                        {
                          "acl_id": "string",
                          "status": "string"
                        }
                      ]
                    ],
                    "acl_type": "string"
                  }
                ]
              ],
              "certificates": [
                "list",
                [
                  "object",
                  {
                    "certificate_id": "string"
                  }
                ]
              ],
              "default_actions": [
                "list",
                [
                  "object",
                  {
                    "forward_group_config": [
                      "list",
                      [
                        "object",
                        {
                          "server_group_tuples": [
                            "list",
                            [
                              "object",
                              {
                                "server_group_id": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "type": "string"
                  }
                ]
              ],
              "gzip_enabled": "bool",
              "http2_enabled": "bool",
              "id": "string",
              "idle_timeout": "number",
              "listener_description": "string",
              "listener_id": "string",
              "listener_port": "number",
              "listener_protocol": "string",
              "load_balancer_id": "string",
              "max_results": "string",
              "next_token": "string",
              "quic_config": [
                "list",
                [
                  "object",
                  {
                    "quic_listener_id": "string",
                    "quic_upgrade_enabled": "bool"
                  }
                ]
              ],
              "request_timeout": "number",
              "security_policy_id": "string",
              "status": "string",
              "xforwarded_for_config": [
                "list",
                [
                  "object",
                  {
                    "xforwardedforclientcert_issuerdnalias": "string",
                    "xforwardedforclientcert_issuerdnenabled": "bool",
                    "xforwardedforclientcertclientverifyalias": "string",
                    "xforwardedforclientcertclientverifyenabled": "bool",
                    "xforwardedforclientcertfingerprintalias": "string",
                    "xforwardedforclientcertfingerprintenabled": "bool",
                    "xforwardedforclientcertsubjectdnalias": "string",
                    "xforwardedforclientcertsubjectdnenabled": "bool",
                    "xforwardedforclientsrcportenabled": "bool",
                    "xforwardedforenabled": "bool",
                    "xforwardedforprotoenabled": "bool",
                    "xforwardedforslbidenabled": "bool",
                    "xforwardedforslbportenabled": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "load_balancer_ids": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudAlbListenersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbListeners), &result)
	return &result
}
