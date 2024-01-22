package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslJa3Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_ja3_oper`: Operational Status for the object ssl-ja3\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslJa3OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"record": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr_v4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"addr_v6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"amount": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslJa3OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslJa3OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslJa3Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslJa3OperOper := setObjectSlbSslJa3OperOper(res)
		d.Set("oper", SlbSslJa3OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslJa3OperOper(ret edpt.DataSlbSslJa3Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"record": setSliceSlbSslJa3OperOperRecord(ret.DtSlbSslJa3Oper.Oper.Record),
		},
	}
}

func setSliceSlbSslJa3OperOperRecord(d []edpt.SlbSslJa3OperOperRecord) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr_v4"] = item.AddrV4
		in["addr_v6"] = item.AddrV6
		in["amount"] = item.Amount
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslJa3OperOper(d []interface{}) edpt.SlbSslJa3OperOper {

	count1 := len(d)
	var ret edpt.SlbSslJa3OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Record = getSliceSlbSslJa3OperOperRecord(in["record"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslJa3OperOperRecord(d []interface{}) []edpt.SlbSslJa3OperOperRecord {

	count1 := len(d)
	ret := make([]edpt.SlbSslJa3OperOperRecord, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslJa3OperOperRecord
		oi.AddrV4 = in["addr_v4"].(string)
		oi.AddrV6 = in["addr_v6"].(string)
		oi.Amount = in["amount"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslJa3Oper(d *schema.ResourceData) edpt.SlbSslJa3Oper {
	var ret edpt.SlbSslJa3Oper

	ret.Oper = getObjectSlbSslJa3OperOper(d.Get("oper").([]interface{}))
	return ret
}
