package csvutil_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cirius-go/csvutil"
)

func TestConfig_SetCols(t *testing.T) {
	c := &csvutil.Config{}

	cols := csvutil.ColFromKeys(
		"first_name",
		"last_name",
		"collaborator",
		"c.magic",
		"ic",
		"title",
		"addr1",
		"addr2",
		"addr3",
		"postal",
		"phone_number",
		"pdpamobile",
		"pdpasms",
		"email",
		"dob",
		"v.magic",
		"regn_date",
		"vehicle_no",
		"model",
		"chassis_no",
		"engine_no",
		"warr_exp",
		"ext_warr",
		"last_mileage",
		"last_svc",
		"last_work",
		"next_svc",
		"as.exec",
		"s.magic",
		"as.magic",
		"d.magic",
		"company",
		"job_title",
		"contact_owner",
		"lead_source",
		"priority",
		"subscriber",
		"last_channel",
		"date",
		"last_order_number",
		"last_order_items",
		"last_order_amount",
		"last_order_currency",
		"last_order_total_price",
		"last_order_date",
		"last_order_updated_date",
		"last_order_status",
		"last_order_status_url",
		"last_order_financial_status",
		"last_order_fulfillment_status",
		"last_order_remarks",
		"last_order_shipping_method",
		"last_order_tracking_company",
		"last_order_tracking_url",
		"last_order_tracking_number",
		"abadoned_cart_items",
		"abadoned_cart_amount",
		"abadoned_cart_currency",
		"abadoned_cart_total_amount",
		"abadoned_cart_cart_date",
		"abadoned_cart_cart_url",
		"payload",
	)

	c.SetCols(
		cols...,
	)

	csvData := csvutil.New(c)

	err := csvData.EncodeJSON(false, &[]any{
		map[string]any{
			"first_name":   "John",
			"last_name":    "Doe",
			"title":        "Mr",
			"phone_number": "1234567890",
		},
	})

	assert.NoError(t, err)

	err = csvData.Write(os.Stdout)
	assert.NoError(t, err)
}
