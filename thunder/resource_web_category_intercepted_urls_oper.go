package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryInterceptedUrlsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_intercepted_urls_oper`: Operational Status for the object intercepted-urls\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryInterceptedUrlsOperRead,

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

func resourceWebCategoryInterceptedUrlsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryInterceptedUrlsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryInterceptedUrlsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryInterceptedUrlsOperOper := setObjectWebCategoryInterceptedUrlsOperOper(res)
		d.Set("oper", WebCategoryInterceptedUrlsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryInterceptedUrlsOperOper(ret edpt.DataWebCategoryInterceptedUrlsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"url_list":       setSliceWebCategoryInterceptedUrlsOperOperUrlList(ret.DtWebCategoryInterceptedUrlsOper.Oper.UrlList),
			"number_of_urls": ret.DtWebCategoryInterceptedUrlsOper.Oper.NumberOfUrls,
			"all_urls":       ret.DtWebCategoryInterceptedUrlsOper.Oper.AllUrls,
			"url_name":       ret.DtWebCategoryInterceptedUrlsOper.Oper.UrlName,
		},
	}
}

func setSliceWebCategoryInterceptedUrlsOperOperUrlList(d []edpt.WebCategoryInterceptedUrlsOperOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func getObjectWebCategoryInterceptedUrlsOperOper(d []interface{}) edpt.WebCategoryInterceptedUrlsOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryInterceptedUrlsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryInterceptedUrlsOperOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryInterceptedUrlsOperOperUrlList(d []interface{}) []edpt.WebCategoryInterceptedUrlsOperOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryInterceptedUrlsOperOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryInterceptedUrlsOperOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryInterceptedUrlsOper(d *schema.ResourceData) edpt.WebCategoryInterceptedUrlsOper {
	var ret edpt.WebCategoryInterceptedUrlsOper

	ret.Oper = getObjectWebCategoryInterceptedUrlsOperOper(d.Get("oper").([]interface{}))
	return ret
}
