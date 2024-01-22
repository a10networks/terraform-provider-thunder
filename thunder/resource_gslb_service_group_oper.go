package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_group_oper`: Operational Status for the object service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceGroupOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"best": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_second_hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"update": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aging": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"matched": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"service_group_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Service Group name",
			},
		},
	}
}

func resourceGslbServiceGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceGroupOperOper := setObjectGslbServiceGroupOperOper(res)
		d.Set("oper", GslbServiceGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServiceGroupOperOper(ret edpt.DataGslbServiceGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":   setSliceGslbServiceGroupOperOperSessionList(ret.DtGslbServiceGroupOper.Oper.SessionList),
			"matched":        ret.DtGslbServiceGroupOper.Oper.Matched,
			"total_sessions": ret.DtGslbServiceGroupOper.Oper.TotalSessions,
		},
	}
}

func setSliceGslbServiceGroupOperOperSessionList(d []edpt.GslbServiceGroupOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["client"] = item.Client
		in["best"] = item.Best
		in["mode"] = item.Mode
		in["hits"] = item.Hits
		in["last_second_hits"] = item.LastSecondHits
		in["ttl"] = item.Ttl
		in["update"] = item.Update
		in["aging"] = item.Aging
		result = append(result, in)
	}
	return result
}

func getObjectGslbServiceGroupOperOper(d []interface{}) edpt.GslbServiceGroupOperOper {

	count1 := len(d)
	var ret edpt.GslbServiceGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceGslbServiceGroupOperOperSessionList(in["session_list"].([]interface{}))
		ret.Matched = in["matched"].(int)
		ret.TotalSessions = in["total_sessions"].(int)
	}
	return ret
}

func getSliceGslbServiceGroupOperOperSessionList(d []interface{}) []edpt.GslbServiceGroupOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.GslbServiceGroupOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceGroupOperOperSessionList
		oi.Client = in["client"].(string)
		oi.Best = in["best"].(string)
		oi.Mode = in["mode"].(string)
		oi.Hits = in["hits"].(int)
		oi.LastSecondHits = in["last_second_hits"].(int)
		oi.Ttl = in["ttl"].(string)
		oi.Update = in["update"].(int)
		oi.Aging = in["aging"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbServiceGroupOper(d *schema.ResourceData) edpt.GslbServiceGroupOper {
	var ret edpt.GslbServiceGroupOper

	ret.Oper = getObjectGslbServiceGroupOperOper(d.Get("oper").([]interface{}))

	ret.ServiceGroupName = d.Get("service_group_name").(string)
	return ret
}
