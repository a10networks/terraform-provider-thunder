package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryWebReputationOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_web_reputation_oper`: Operational Status for the object web-reputation\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryWebReputationOperRead,

		Schema: map[string]*schema.Schema{
			"bypassed_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"intercepted_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
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
			"url": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"reputation_score": {
										Type: schema.TypeString, Optional: true, Description: "",
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
				},
			},
		},
	}
}

func resourceWebCategoryWebReputationOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryWebReputationOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryWebReputationOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryWebReputationOperBypassedUrls := setObjectWebCategoryWebReputationOperBypassedUrls(res)
		d.Set("bypassed_urls", WebCategoryWebReputationOperBypassedUrls)
		WebCategoryWebReputationOperInterceptedUrls := setObjectWebCategoryWebReputationOperInterceptedUrls(res)
		d.Set("intercepted_urls", WebCategoryWebReputationOperInterceptedUrls)
		WebCategoryWebReputationOperOper := setObjectWebCategoryWebReputationOperOper(res)
		d.Set("oper", WebCategoryWebReputationOperOper)
		WebCategoryWebReputationOperUrl := setObjectWebCategoryWebReputationOperUrl(res)
		d.Set("url", WebCategoryWebReputationOperUrl)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryWebReputationOperBypassedUrls(ret edpt.DataWebCategoryWebReputationOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryWebReputationOperBypassedUrlsOper(ret.DtWebCategoryWebReputationOper.BypassedUrls.Oper),
		},
	}
}

func setObjectWebCategoryWebReputationOperBypassedUrlsOper(d edpt.WebCategoryWebReputationOperBypassedUrlsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryWebReputationOperBypassedUrlsOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryWebReputationOperBypassedUrlsOperUrlList(d []edpt.WebCategoryWebReputationOperBypassedUrlsOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryWebReputationOperInterceptedUrls(ret edpt.DataWebCategoryWebReputationOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryWebReputationOperInterceptedUrlsOper(ret.DtWebCategoryWebReputationOper.InterceptedUrls.Oper),
		},
	}
}

func setObjectWebCategoryWebReputationOperInterceptedUrlsOper(d edpt.WebCategoryWebReputationOperInterceptedUrlsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryWebReputationOperInterceptedUrlsOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryWebReputationOperInterceptedUrlsOperUrlList(d []edpt.WebCategoryWebReputationOperInterceptedUrlsOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryWebReputationOperOper(ret edpt.DataWebCategoryWebReputationOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"url_list":       setSliceWebCategoryWebReputationOperOperUrlList(ret.DtWebCategoryWebReputationOper.Oper.UrlList),
			"number_of_urls": ret.DtWebCategoryWebReputationOper.Oper.NumberOfUrls,
			"all_urls":       ret.DtWebCategoryWebReputationOper.Oper.AllUrls,
			"url_name":       ret.DtWebCategoryWebReputationOper.Oper.UrlName,
		},
	}
}

func setSliceWebCategoryWebReputationOperOperUrlList(d []edpt.WebCategoryWebReputationOperOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryWebReputationOperUrl(ret edpt.DataWebCategoryWebReputationOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryWebReputationOperUrlOper(ret.DtWebCategoryWebReputationOper.Url.Oper),
		},
	}
}

func setObjectWebCategoryWebReputationOperUrlOper(d edpt.WebCategoryWebReputationOperUrlOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["reputation_score"] = d.ReputationScore

	in["name"] = d.Name

	in["local_db_only"] = d.LocalDbOnly
	result = append(result, in)
	return result
}

func getObjectWebCategoryWebReputationOperBypassedUrls(d []interface{}) edpt.WebCategoryWebReputationOperBypassedUrls {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperBypassedUrls
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryWebReputationOperBypassedUrlsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryWebReputationOperBypassedUrlsOper(d []interface{}) edpt.WebCategoryWebReputationOperBypassedUrlsOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperBypassedUrlsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryWebReputationOperBypassedUrlsOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryWebReputationOperBypassedUrlsOperUrlList(d []interface{}) []edpt.WebCategoryWebReputationOperBypassedUrlsOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryWebReputationOperBypassedUrlsOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryWebReputationOperBypassedUrlsOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryWebReputationOperInterceptedUrls(d []interface{}) edpt.WebCategoryWebReputationOperInterceptedUrls {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperInterceptedUrls
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryWebReputationOperInterceptedUrlsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryWebReputationOperInterceptedUrlsOper(d []interface{}) edpt.WebCategoryWebReputationOperInterceptedUrlsOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperInterceptedUrlsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryWebReputationOperInterceptedUrlsOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryWebReputationOperInterceptedUrlsOperUrlList(d []interface{}) []edpt.WebCategoryWebReputationOperInterceptedUrlsOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryWebReputationOperInterceptedUrlsOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryWebReputationOperInterceptedUrlsOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryWebReputationOperOper(d []interface{}) edpt.WebCategoryWebReputationOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryWebReputationOperOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryWebReputationOperOperUrlList(d []interface{}) []edpt.WebCategoryWebReputationOperOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryWebReputationOperOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryWebReputationOperOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryWebReputationOperUrl(d []interface{}) edpt.WebCategoryWebReputationOperUrl {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperUrl
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryWebReputationOperUrlOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryWebReputationOperUrlOper(d []interface{}) edpt.WebCategoryWebReputationOperUrlOper {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputationOperUrlOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ReputationScore = in["reputation_score"].(string)
		ret.Name = in["name"].(string)
		ret.LocalDbOnly = in["local_db_only"].(int)
	}
	return ret
}

func dataToEndpointWebCategoryWebReputationOper(d *schema.ResourceData) edpt.WebCategoryWebReputationOper {
	var ret edpt.WebCategoryWebReputationOper

	ret.BypassedUrls = getObjectWebCategoryWebReputationOperBypassedUrls(d.Get("bypassed_urls").([]interface{}))

	ret.InterceptedUrls = getObjectWebCategoryWebReputationOperInterceptedUrls(d.Get("intercepted_urls").([]interface{}))

	ret.Oper = getObjectWebCategoryWebReputationOperOper(d.Get("oper").([]interface{}))

	ret.Url = getObjectWebCategoryWebReputationOperUrl(d.Get("url").([]interface{}))
	return ret
}
