package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityZbarTruplesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_zbar_truples_oper`: Operational Status for the object truples\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityZbarTruplesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"phase": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"zbar_multi_ind_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ind_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ind_total_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ind_band_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zbar_ind_sec_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"second": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"value": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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

func resourceVisibilityZbarTruplesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarTruplesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarTruplesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityZbarTruplesOperOper := setObjectVisibilityZbarTruplesOperOper(res)
		d.Set("oper", VisibilityZbarTruplesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityZbarTruplesOperOper(ret edpt.DataVisibilityZbarTruplesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"src_ipv4_addr":       ret.DtVisibilityZbarTruplesOper.Oper.SrcIpv4Addr,
			"src_ipv6_addr":       ret.DtVisibilityZbarTruplesOper.Oper.SrcIpv6Addr,
			"phase":               ret.DtVisibilityZbarTruplesOper.Oper.Phase,
			"dest_ipv4_addr":      ret.DtVisibilityZbarTruplesOper.Oper.DestIpv4Addr,
			"dest_ipv6_addr":      ret.DtVisibilityZbarTruplesOper.Oper.DestIpv6Addr,
			"port":                ret.DtVisibilityZbarTruplesOper.Oper.Port,
			"protocol":            ret.DtVisibilityZbarTruplesOper.Oper.Protocol,
			"zbar_multi_ind_list": setSliceVisibilityZbarTruplesOperOperZbarMultiIndList(ret.DtVisibilityZbarTruplesOper.Oper.ZbarMultiIndList),
		},
	}
}

func setSliceVisibilityZbarTruplesOperOperZbarMultiIndList(d []edpt.VisibilityZbarTruplesOperOperZbarMultiIndList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ind_name"] = item.IndName
		in["ind_total_count"] = item.IndTotalCount
		in["ind_band_id"] = item.IndBandId
		in["zbar_ind_sec_list"] = setSliceVisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList(item.ZbarIndSecList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList(d []edpt.VisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["second"] = item.Second
		in["value"] = item.Value
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityZbarTruplesOperOper(d []interface{}) edpt.VisibilityZbarTruplesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityZbarTruplesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcIpv4Addr = in["src_ipv4_addr"].(string)
		ret.SrcIpv6Addr = in["src_ipv6_addr"].(string)
		ret.Phase = in["phase"].(string)
		ret.DestIpv4Addr = in["dest_ipv4_addr"].(string)
		ret.DestIpv6Addr = in["dest_ipv6_addr"].(string)
		ret.Port = in["port"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.ZbarMultiIndList = getSliceVisibilityZbarTruplesOperOperZbarMultiIndList(in["zbar_multi_ind_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityZbarTruplesOperOperZbarMultiIndList(d []interface{}) []edpt.VisibilityZbarTruplesOperOperZbarMultiIndList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarTruplesOperOperZbarMultiIndList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarTruplesOperOperZbarMultiIndList
		oi.IndName = in["ind_name"].(string)
		oi.IndTotalCount = in["ind_total_count"].(int)
		oi.IndBandId = in["ind_band_id"].(int)
		oi.ZbarIndSecList = getSliceVisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList(in["zbar_ind_sec_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList(d []interface{}) []edpt.VisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarTruplesOperOperZbarMultiIndListZbarIndSecList
		oi.Second = in["second"].(string)
		oi.Value = in["value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityZbarTruplesOper(d *schema.ResourceData) edpt.VisibilityZbarTruplesOper {
	var ret edpt.VisibilityZbarTruplesOper

	ret.Oper = getObjectVisibilityZbarTruplesOperOper(d.Get("oper").([]interface{}))
	return ret
}
