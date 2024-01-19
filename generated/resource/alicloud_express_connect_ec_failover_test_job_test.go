package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-alicloud-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAlicloudExpressConnectEcFailoverTestJobSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AlicloudExpressConnectEcFailoverTestJobSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
