package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryUrlOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_url_oper`: Operational Status for the object url\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryUrlOperRead,

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
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"local_db_only": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryUrlOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryUrlOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryUrlOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryUrlOperOper := setObjectWebCategoryUrlOperOper(res)
		d.Set("oper", WebCategoryUrlOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryUrlOperOper(ret edpt.DataWebCategoryUrlOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"category_list": setSliceWebCategoryUrlOperOperCategoryList(ret.DtWebCategoryUrlOper.Oper.CategoryList),
			"name":          ret.DtWebCategoryUrlOper.Oper.Name,
			"local_db_only": ret.DtWebCategoryUrlOper.Oper.LocalDbOnly,
		},
	}
}

func setSliceWebCategoryUrlOperOperCategoryList(d []edpt.WebCategoryUrlOperOperCategoryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["category"] = item.Category
		result = append(result, in)
	}
	return result
}

func getObjectWebCategoryUrlOperOper(d []interface{}) edpt.WebCategoryUrlOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryUrlOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CategoryList = getSliceWebCategoryUrlOperOperCategoryList(in["category_list"].([]interface{}))
		ret.Name = in["name"].(string)
		ret.LocalDbOnly = in["local_db_only"].(int)
	}
	return ret
}

func getSliceWebCategoryUrlOperOperCategoryList(d []interface{}) []edpt.WebCategoryUrlOperOperCategoryList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryUrlOperOperCategoryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryUrlOperOperCategoryList
		oi.Category = in["category"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryUrlOper(d *schema.ResourceData) edpt.WebCategoryUrlOper {
	var ret edpt.WebCategoryUrlOper

	ret.Oper = getObjectWebCategoryUrlOperOper(d.Get("oper").([]interface{}))
	return ret
}
