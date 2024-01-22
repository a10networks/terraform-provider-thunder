package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryWebReputationInterceptedUrlsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_web_reputation_intercepted_urls_oper`: Operational Status for the object intercepted-urls\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryWebReputationInterceptedUrlsOperRead,

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

func resourceWebCategoryWebReputationInterceptedUrlsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationInterceptedUrlsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputationInterceptedUrlsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryWebReputationInterceptedUrlsOperOper := setObjectWebCategoryWebReputationInterceptedUrlsOperOper(res)
		d.Set("oper", WebCategoryWebReputationInterceptedUrlsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryWebReputationInterceptedUrlsOperOper(ret edpt.DataWebCategoryWebReputationInterceptedUrlsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"url_list":       setSliceWebCategoryWebReputationInterceptedUrlsOperOperUrlList(ret.DtWebCategoryWebReputationInterceptedUrlsOper.Oper.UrlList),
			"number_of_urls": ret.DtWebCategoryWebReputationInterceptedUrlsOper.Oper.NumberOfUrls,
			"all_urls":       ret.DtWebCategoryWebReputationInterceptedUrlsOper.Oper.AllUrls,
			"url_name":       ret.DtWebCategoryWebReputationInterceptedUrlsOper.Oper.UrlName,
		},
	}
}

func setSliceWebCategoryWebReputationInterceptedUrlsOperOperUrlList(d []edpt.WebCategoryWebReputationInterceptedUrlsOperOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func getObjectWebCategoryWebReputationInterceptedUrlsOperOper(d []interface{}) edpt.WebCategoryWebReputationInterceptedUrlsOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationInterceptedUrlsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryWebReputationInterceptedUrlsOperOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryWebReputationInterceptedUrlsOperOperUrlList(d []interface{}) []edpt.WebCategoryWebReputationInterceptedUrlsOperOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryWebReputationInterceptedUrlsOperOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryWebReputationInterceptedUrlsOperOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryWebReputationInterceptedUrlsOper(d *schema.ResourceData) edpt.WebCategoryWebReputationInterceptedUrlsOper {
	var ret edpt.WebCategoryWebReputationInterceptedUrlsOper

	ret.Oper = getObjectWebCategoryWebReputationInterceptedUrlsOperOper(d.Get("oper").([]interface{}))
	return ret
}
