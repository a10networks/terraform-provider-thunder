package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemClThreatCategoryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cl_threat_category_oper`: Operational Status for the object cl-threat-category\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemClThreatCategoryOperRead,

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
					},
				},
			},
		},
	}
}

func resourceSystemClThreatCategoryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemClThreatCategoryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemClThreatCategoryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemClThreatCategoryOperOper := setObjectSystemClThreatCategoryOperOper(res)
		d.Set("oper", SystemClThreatCategoryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemClThreatCategoryOperOper(ret edpt.DataSystemClThreatCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"category_list": setSliceSystemClThreatCategoryOperOperCategoryList(ret.DtSystemClThreatCategoryOper.Oper.CategoryList),
		},
	}
}

func setSliceSystemClThreatCategoryOperOperCategoryList(d []edpt.SystemClThreatCategoryOperOperCategoryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["category"] = item.Category
		result = append(result, in)
	}
	return result
}

func getObjectSystemClThreatCategoryOperOper(d []interface{}) edpt.SystemClThreatCategoryOperOper {

	count1 := len(d)
	var ret edpt.SystemClThreatCategoryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CategoryList = getSliceSystemClThreatCategoryOperOperCategoryList(in["category_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemClThreatCategoryOperOperCategoryList(d []interface{}) []edpt.SystemClThreatCategoryOperOperCategoryList {

	count1 := len(d)
	ret := make([]edpt.SystemClThreatCategoryOperOperCategoryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemClThreatCategoryOperOperCategoryList
		oi.Category = in["category"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemClThreatCategoryOper(d *schema.ResourceData) edpt.SystemClThreatCategoryOper {
	var ret edpt.SystemClThreatCategoryOper

	ret.Oper = getObjectSystemClThreatCategoryOperOper(d.Get("oper").([]interface{}))
	return ret
}
