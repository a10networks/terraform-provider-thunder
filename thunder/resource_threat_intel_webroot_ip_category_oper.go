package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelWebrootIpCategoryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_threat_intel_webroot_ip_category_oper`: Operational Status for the object webroot-ip-category\n\n__PLACEHOLDER__",
		ReadContext: resourceThreatIntelWebrootIpCategoryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceThreatIntelWebrootIpCategoryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootIpCategoryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootIpCategoryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ThreatIntelWebrootIpCategoryOperOper := setObjectThreatIntelWebrootIpCategoryOperOper(res)
		d.Set("oper", ThreatIntelWebrootIpCategoryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectThreatIntelWebrootIpCategoryOperOper(ret edpt.DataThreatIntelWebrootIpCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"category_list": setSliceThreatIntelWebrootIpCategoryOperOperCategoryList(ret.DtThreatIntelWebrootIpCategoryOper.Oper.CategoryList),
			"ip":            ret.DtThreatIntelWebrootIpCategoryOper.Oper.Ip,
		},
	}
}

func setSliceThreatIntelWebrootIpCategoryOperOperCategoryList(d []edpt.ThreatIntelWebrootIpCategoryOperOperCategoryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["category"] = item.Category
		result = append(result, in)
	}
	return result
}

func getObjectThreatIntelWebrootIpCategoryOperOper(d []interface{}) edpt.ThreatIntelWebrootIpCategoryOperOper {

	count1 := len(d)
	var ret edpt.ThreatIntelWebrootIpCategoryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CategoryList = getSliceThreatIntelWebrootIpCategoryOperOperCategoryList(in["category_list"].([]interface{}))
		ret.Ip = in["ip"].(string)
	}
	return ret
}

func getSliceThreatIntelWebrootIpCategoryOperOperCategoryList(d []interface{}) []edpt.ThreatIntelWebrootIpCategoryOperOperCategoryList {

	count1 := len(d)
	ret := make([]edpt.ThreatIntelWebrootIpCategoryOperOperCategoryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ThreatIntelWebrootIpCategoryOperOperCategoryList
		oi.Category = in["category"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointThreatIntelWebrootIpCategoryOper(d *schema.ResourceData) edpt.ThreatIntelWebrootIpCategoryOper {
	var ret edpt.ThreatIntelWebrootIpCategoryOper

	ret.Oper = getObjectThreatIntelWebrootIpCategoryOperOper(d.Get("oper").([]interface{}))
	return ret
}
