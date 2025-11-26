package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityZbarDestBadSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_zbar_dest_bad_sources_oper`: Operational Status for the object bad-sources\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityZbarDestBadSourcesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"multi_bad_src_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ind_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drop_cnt": {
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

func resourceVisibilityZbarDestBadSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDestBadSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarDestBadSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityZbarDestBadSourcesOperOper := setObjectVisibilityZbarDestBadSourcesOperOper(res)
		d.Set("oper", VisibilityZbarDestBadSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityZbarDestBadSourcesOperOper(ret edpt.DataVisibilityZbarDestBadSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipv4_addr":          ret.DtVisibilityZbarDestBadSourcesOper.Oper.Ipv4Addr,
			"ipv6_addr":          ret.DtVisibilityZbarDestBadSourcesOper.Oper.Ipv6Addr,
			"port":               ret.DtVisibilityZbarDestBadSourcesOper.Oper.Port,
			"protocol":           ret.DtVisibilityZbarDestBadSourcesOper.Oper.Protocol,
			"multi_bad_src_list": setSliceVisibilityZbarDestBadSourcesOperOperMultiBadSrcList(ret.DtVisibilityZbarDestBadSourcesOper.Oper.MultiBadSrcList),
		},
	}
}

func setSliceVisibilityZbarDestBadSourcesOperOperMultiBadSrcList(d []edpt.VisibilityZbarDestBadSourcesOperOperMultiBadSrcList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_ip"] = item.SrcIp
		in["ind_value"] = item.IndValue
		in["state"] = item.State
		in["drop_cnt"] = item.DropCnt
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityZbarDestBadSourcesOperOper(d []interface{}) edpt.VisibilityZbarDestBadSourcesOperOper {

	count1 := len(d)
	var ret edpt.VisibilityZbarDestBadSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.Port = in["port"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.MultiBadSrcList = getSliceVisibilityZbarDestBadSourcesOperOperMultiBadSrcList(in["multi_bad_src_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityZbarDestBadSourcesOperOperMultiBadSrcList(d []interface{}) []edpt.VisibilityZbarDestBadSourcesOperOperMultiBadSrcList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarDestBadSourcesOperOperMultiBadSrcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarDestBadSourcesOperOperMultiBadSrcList
		oi.SrcIp = in["src_ip"].(string)
		oi.IndValue = in["ind_value"].(int)
		oi.State = in["state"].(string)
		oi.DropCnt = in["drop_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityZbarDestBadSourcesOper(d *schema.ResourceData) edpt.VisibilityZbarDestBadSourcesOper {
	var ret edpt.VisibilityZbarDestBadSourcesOper

	ret.Oper = getObjectVisibilityZbarDestBadSourcesOperOper(d.Get("oper").([]interface{}))
	return ret
}
