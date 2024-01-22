package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessionsSmpTableOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sessions_smp_table_oper`: Operational Status for the object smp-table\n\n__PLACEHOLDER__",
		ReadContext: resourceSessionsSmpTableOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dst4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dst6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"srcport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dstport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"payload": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSessionsSmpTableOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsSmpTableOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionsSmpTableOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SessionsSmpTableOperOper := setObjectSessionsSmpTableOperOper(res)
		d.Set("oper", SessionsSmpTableOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSessionsSmpTableOperOper(ret edpt.DataSessionsSmpTableOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceSessionsSmpTableOperOperEntryList(ret.DtSessionsSmpTableOper.Oper.EntryList),
		},
	}
}

func setSliceSessionsSmpTableOperOperEntryList(d []edpt.SessionsSmpTableOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src4"] = item.Src4
		in["src6"] = item.Src6
		in["dst4"] = item.Dst4
		in["dst6"] = item.Dst6
		in["srcport"] = item.Srcport
		in["dstport"] = item.Dstport
		in["ttl"] = item.Ttl
		in["type"] = item.Type
		in["payload"] = item.Payload
		result = append(result, in)
	}
	return result
}

func getObjectSessionsSmpTableOperOper(d []interface{}) edpt.SessionsSmpTableOperOper {

	count1 := len(d)
	var ret edpt.SessionsSmpTableOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceSessionsSmpTableOperOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceSessionsSmpTableOperOperEntryList(d []interface{}) []edpt.SessionsSmpTableOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.SessionsSmpTableOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsSmpTableOperOperEntryList
		oi.Src4 = in["src4"].(string)
		oi.Src6 = in["src6"].(string)
		oi.Dst4 = in["dst4"].(string)
		oi.Dst6 = in["dst6"].(string)
		oi.Srcport = in["srcport"].(int)
		oi.Dstport = in["dstport"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.Type = in["type"].(string)
		oi.Payload = in["payload"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSessionsSmpTableOper(d *schema.ResourceData) edpt.SessionsSmpTableOper {
	var ret edpt.SessionsSmpTableOper

	ret.Oper = getObjectSessionsSmpTableOperOper(d.Get("oper").([]interface{}))
	return ret
}
