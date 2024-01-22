package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryWebReputationBypassedUrlsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_web_reputation_bypassed_urls_oper`: Operational Status for the object bypassed-urls\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryWebReputationBypassedUrlsOperRead,

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

func resourceWebCategoryWebReputationBypassedUrlsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationBypassedUrlsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputationBypassedUrlsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryWebReputationBypassedUrlsOperOper := setObjectWebCategoryWebReputationBypassedUrlsOperOper(res)
		d.Set("oper", WebCategoryWebReputationBypassedUrlsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryWebReputationBypassedUrlsOperOper(ret edpt.DataWebCategoryWebReputationBypassedUrlsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"url_list":       setSliceWebCategoryWebReputationBypassedUrlsOperOperUrlList(ret.DtWebCategoryWebReputationBypassedUrlsOper.Oper.UrlList),
			"number_of_urls": ret.DtWebCategoryWebReputationBypassedUrlsOper.Oper.NumberOfUrls,
			"all_urls":       ret.DtWebCategoryWebReputationBypassedUrlsOper.Oper.AllUrls,
			"url_name":       ret.DtWebCategoryWebReputationBypassedUrlsOper.Oper.UrlName,
		},
	}
}

func setSliceWebCategoryWebReputationBypassedUrlsOperOperUrlList(d []edpt.WebCategoryWebReputationBypassedUrlsOperOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func getObjectWebCategoryWebReputationBypassedUrlsOperOper(d []interface{}) edpt.WebCategoryWebReputationBypassedUrlsOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationBypassedUrlsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryWebReputationBypassedUrlsOperOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryWebReputationBypassedUrlsOperOperUrlList(d []interface{}) []edpt.WebCategoryWebReputationBypassedUrlsOperOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryWebReputationBypassedUrlsOperOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryWebReputationBypassedUrlsOperOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryWebReputationBypassedUrlsOper(d *schema.ResourceData) edpt.WebCategoryWebReputationBypassedUrlsOper {
	var ret edpt.WebCategoryWebReputationBypassedUrlsOper

	ret.Oper = getObjectWebCategoryWebReputationBypassedUrlsOperOper(d.Get("oper").([]interface{}))
	return ret
}
