package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFwdPolicySnatPtOnlyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_fwd_policy_snat_pt_only_oper`: Operational Status for the object fwd-policy-snat-pt-only\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFwdPolicySnatPtOnlyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_usage": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"detail_info": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFwdPolicySnatPtOnlyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFwdPolicySnatPtOnlyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFwdPolicySnatPtOnlyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFwdPolicySnatPtOnlyOperOper := setObjectSlbFwdPolicySnatPtOnlyOperOper(res)
		d.Set("oper", SlbFwdPolicySnatPtOnlyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFwdPolicySnatPtOnlyOperOper(ret edpt.DataSlbFwdPolicySnatPtOnlyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"group_list":  setSliceSlbFwdPolicySnatPtOnlyOperOperGroupList(ret.DtSlbFwdPolicySnatPtOnlyOper.Oper.GroupList),
			"detail_info": ret.DtSlbFwdPolicySnatPtOnlyOper.Oper.DetailInfo,
		},
	}
}

func setSliceSlbFwdPolicySnatPtOnlyOperOperGroupList(d []edpt.SlbFwdPolicySnatPtOnlyOperOperGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["group"] = item.Group
		in["port_usage"] = item.PortUsage
		in["total_used"] = item.TotalUsed
		in["total_freed"] = item.TotalFreed
		in["failed"] = item.Failed
		result = append(result, in)
	}
	return result
}

func getObjectSlbFwdPolicySnatPtOnlyOperOper(d []interface{}) edpt.SlbFwdPolicySnatPtOnlyOperOper {

	count1 := len(d)
	var ret edpt.SlbFwdPolicySnatPtOnlyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupList = getSliceSlbFwdPolicySnatPtOnlyOperOperGroupList(in["group_list"].([]interface{}))
		ret.DetailInfo = in["detail_info"].(int)
	}
	return ret
}

func getSliceSlbFwdPolicySnatPtOnlyOperOperGroupList(d []interface{}) []edpt.SlbFwdPolicySnatPtOnlyOperOperGroupList {

	count1 := len(d)
	ret := make([]edpt.SlbFwdPolicySnatPtOnlyOperOperGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbFwdPolicySnatPtOnlyOperOperGroupList
		oi.Group = in["group"].(string)
		oi.PortUsage = in["port_usage"].(int)
		oi.TotalUsed = in["total_used"].(int)
		oi.TotalFreed = in["total_freed"].(int)
		oi.Failed = in["failed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbFwdPolicySnatPtOnlyOper(d *schema.ResourceData) edpt.SlbFwdPolicySnatPtOnlyOper {
	var ret edpt.SlbFwdPolicySnatPtOnlyOper

	ret.Oper = getObjectSlbFwdPolicySnatPtOnlyOperOper(d.Get("oper").([]interface{}))
	return ret
}
