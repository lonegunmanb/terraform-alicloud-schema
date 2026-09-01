package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudExpressConnectRouterGrantAssociationsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AlicloudExpressConnectRouterGrantAssociationsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
