package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryBypassedUrlsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_bypassed_urls_oper`: Operational Status for the object bypassed-urls\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryBypassedUrlsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"number_of_urls": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_urls": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"url_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryBypassedUrlsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryBypassedUrlsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryBypassedUrlsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryBypassedUrlsOperOper := setObjectWebCategoryBypassedUrlsOperOper(res)
		d.Set("oper", WebCategoryBypassedUrlsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryBypassedUrlsOperOper(ret edpt.DataWebCategoryBypassedUrlsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"url_list":       setSliceWebCategoryBypassedUrlsOperOperUrlList(ret.DtWebCategoryBypassedUrlsOper.Oper.UrlList),
			"number_of_urls": ret.DtWebCategoryBypassedUrlsOper.Oper.NumberOfUrls,
			"all_urls":       ret.DtWebCategoryBypassedUrlsOper.Oper.AllUrls,
			"url_name":       ret.DtWebCategoryBypassedUrlsOper.Oper.UrlName,
		},
	}
}

func setSliceWebCategoryBypassedUrlsOperOperUrlList(d []edpt.WebCategoryBypassedUrlsOperOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func getObjectWebCategoryBypassedUrlsOperOper(d []interface{}) edpt.WebCategoryBypassedUrlsOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryBypassedUrlsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryBypassedUrlsOperOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryBypassedUrlsOperOperUrlList(d []interface{}) []edpt.WebCategoryBypassedUrlsOperOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryBypassedUrlsOperOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryBypassedUrlsOperOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryBypassedUrlsOper(d *schema.ResourceData) edpt.WebCategoryBypassedUrlsOper {
	var ret edpt.WebCategoryBypassedUrlsOper

	ret.Oper = getObjectWebCategoryBypassedUrlsOperOper(d.Get("oper").([]interface{}))
	return ret
}
