package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdCgnOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_cgn_oper`: Operational Status for the object cgn\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdCgnOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"end_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cgn_data": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_user_quota_create": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_user_quota_delete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat64_user_quota_create": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat64_user_quota_delete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dslite_user_quota_create": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dslite_user_quota_delete": {
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

func resourceRrdCgnOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdCgnOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdCgnOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdCgnOperOper := setObjectRrdCgnOperOper(res)
		d.Set("oper", RrdCgnOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdCgnOperOper(ret edpt.DataRrdCgnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time": ret.DtRrdCgnOper.Oper.StartTime,
			"end_time":   ret.DtRrdCgnOper.Oper.EndTime,
			"cgn_data":   setSliceRrdCgnOperOperCgnData(ret.DtRrdCgnOper.Oper.CgnData),
		},
	}
}

func setSliceRrdCgnOperOperCgnData(d []edpt.RrdCgnOperOperCgnData) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["lsn_user_quota_create"] = item.Lsn_user_quota_create
		in["lsn_user_quota_delete"] = item.Lsn_user_quota_delete
		in["nat64_user_quota_create"] = item.Nat64_user_quota_create
		in["nat64_user_quota_delete"] = item.Nat64_user_quota_delete
		in["dslite_user_quota_create"] = item.Dslite_user_quota_create
		in["dslite_user_quota_delete"] = item.Dslite_user_quota_delete
		result = append(result, in)
	}
	return result
}

func getObjectRrdCgnOperOper(d []interface{}) edpt.RrdCgnOperOper {

	count1 := len(d)
	var ret edpt.RrdCgnOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.CgnData = getSliceRrdCgnOperOperCgnData(in["cgn_data"].([]interface{}))
	}
	return ret
}

func getSliceRrdCgnOperOperCgnData(d []interface{}) []edpt.RrdCgnOperOperCgnData {

	count1 := len(d)
	ret := make([]edpt.RrdCgnOperOperCgnData, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdCgnOperOperCgnData
		oi.Time = in["time"].(int)
		oi.Lsn_user_quota_create = in["lsn_user_quota_create"].(int)
		oi.Lsn_user_quota_delete = in["lsn_user_quota_delete"].(int)
		oi.Nat64_user_quota_create = in["nat64_user_quota_create"].(int)
		oi.Nat64_user_quota_delete = in["nat64_user_quota_delete"].(int)
		oi.Dslite_user_quota_create = in["dslite_user_quota_create"].(int)
		oi.Dslite_user_quota_delete = in["dslite_user_quota_delete"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdCgnOper(d *schema.ResourceData) edpt.RrdCgnOper {
	var ret edpt.RrdCgnOper

	ret.Oper = getObjectRrdCgnOperOper(d.Get("oper").([]interface{}))
	return ret
}
