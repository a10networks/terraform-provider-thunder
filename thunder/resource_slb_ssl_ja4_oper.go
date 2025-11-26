package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslJa4Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_ja4_oper`: Operational Status for the object ssl-ja4\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslJa4OperRead,

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

func resourceSlbSslJa4OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslJa4OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslJa4Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslJa4OperOper := setObjectSlbSslJa4OperOper(res)
		d.Set("oper", SlbSslJa4OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslJa4OperOper(ret edpt.DataSlbSslJa4Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"record": setSliceSlbSslJa4OperOperRecord(ret.DtSlbSslJa4Oper.Oper.Record),
		},
	}
}

func setSliceSlbSslJa4OperOperRecord(d []edpt.SlbSslJa4OperOperRecord) []map[string]interface{} {
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

func getObjectSlbSslJa4OperOper(d []interface{}) edpt.SlbSslJa4OperOper {

	count1 := len(d)
	var ret edpt.SlbSslJa4OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Record = getSliceSlbSslJa4OperOperRecord(in["record"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslJa4OperOperRecord(d []interface{}) []edpt.SlbSslJa4OperOperRecord {

	count1 := len(d)
	ret := make([]edpt.SlbSslJa4OperOperRecord, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslJa4OperOperRecord
		oi.AddrV4 = in["addr_v4"].(string)
		oi.AddrV6 = in["addr_v6"].(string)
		oi.Amount = in["amount"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslJa4Oper(d *schema.ResourceData) edpt.SlbSslJa4Oper {
	var ret edpt.SlbSslJa4Oper

	ret.Oper = getObjectSlbSslJa4OperOper(d.Get("oper").([]interface{}))
	return ret
}
