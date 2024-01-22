package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAppOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_app_oper`: Operational Status for the object app\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAppOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"contains": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"related": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_desc": {
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

func resourceFwAppOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAppOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAppOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAppOperOper := setObjectFwAppOperOper(res)
		d.Set("oper", FwAppOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAppOperOper(ret edpt.DataFwAppOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"category":   ret.DtFwAppOper.Oper.Category,
			"contains":   ret.DtFwAppOper.Oper.Contains,
			"related":    ret.DtFwAppOper.Oper.Related,
			"group_list": setSliceFwAppOperOperGroupList(ret.DtFwAppOper.Oper.GroupList),
		},
	}
}

func setSliceFwAppOperOperGroupList(d []edpt.FwAppOperOperGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["category_name"] = item.CategoryName
		in["app_name"] = item.AppName
		in["app_desc"] = item.AppDesc
		result = append(result, in)
	}
	return result
}

func getObjectFwAppOperOper(d []interface{}) edpt.FwAppOperOper {

	count1 := len(d)
	var ret edpt.FwAppOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Category = in["category"].(string)
		ret.Contains = in["contains"].(string)
		ret.Related = in["related"].(string)
		ret.GroupList = getSliceFwAppOperOperGroupList(in["group_list"].([]interface{}))
	}
	return ret
}

func getSliceFwAppOperOperGroupList(d []interface{}) []edpt.FwAppOperOperGroupList {

	count1 := len(d)
	ret := make([]edpt.FwAppOperOperGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAppOperOperGroupList
		oi.CategoryName = in["category_name"].(string)
		oi.AppName = in["app_name"].(string)
		oi.AppDesc = in["app_desc"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwAppOper(d *schema.ResourceData) edpt.FwAppOper {
	var ret edpt.FwAppOper

	ret.Oper = getObjectFwAppOperOper(d.Get("oper").([]interface{}))
	return ret
}
