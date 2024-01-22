package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamRdnsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_rdns_oper`: Operational Status for the object rdns\n\n__PLACEHOLDER__",
		ReadContext: resourceAamRdnsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamRdnsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamRdnsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamRdnsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamRdnsOperOper := setObjectAamRdnsOperOper(res)
		d.Set("oper", AamRdnsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamRdnsOperOper(ret edpt.DataAamRdnsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceAamRdnsOperOperEntryList(ret.DtAamRdnsOper.Oper.EntryList),
			"mode":       ret.DtAamRdnsOper.Oper.Mode,
			"addr":       ret.DtAamRdnsOper.Oper.Addr,
		},
	}
}

func setSliceAamRdnsOperOperEntryList(d []edpt.AamRdnsOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["address"] = item.Address
		in["domain"] = item.Domain
		in["ttl"] = item.Ttl
		result = append(result, in)
	}
	return result
}

func getObjectAamRdnsOperOper(d []interface{}) edpt.AamRdnsOperOper {

	count1 := len(d)
	var ret edpt.AamRdnsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceAamRdnsOperOperEntryList(in["entry_list"].([]interface{}))
		ret.Mode = in["mode"].(string)
		ret.Addr = in["addr"].(string)
	}
	return ret
}

func getSliceAamRdnsOperOperEntryList(d []interface{}) []edpt.AamRdnsOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.AamRdnsOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamRdnsOperOperEntryList
		oi.Type = in["type"].(string)
		oi.Address = in["address"].(string)
		oi.Domain = in["domain"].(string)
		oi.Ttl = in["ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamRdnsOper(d *schema.ResourceData) edpt.AamRdnsOper {
	var ret edpt.AamRdnsOper

	ret.Oper = getObjectAamRdnsOperOper(d.Get("oper").([]interface{}))
	return ret
}
