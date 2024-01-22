package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthorizationJwtCacheOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authorization_jwt_cache_oper`: Operational Status for the object cache\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthorizationJwtCacheOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token_cached": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"token_cache_hit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"max_token_cache": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cache_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cache_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"issuer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"audience": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"client_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"audience": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthorizationJwtCacheOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationJwtCacheOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationJwtCacheOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthorizationJwtCacheOperOper := setObjectAamAuthorizationJwtCacheOperOper(res)
		d.Set("oper", AamAuthorizationJwtCacheOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthorizationJwtCacheOperOper(ret edpt.DataAamAuthorizationJwtCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"token_cached":    ret.DtAamAuthorizationJwtCacheOper.Oper.Token_cached,
			"token_cache_hit": ret.DtAamAuthorizationJwtCacheOper.Oper.Token_cache_hit,
			"max_token_cache": ret.DtAamAuthorizationJwtCacheOper.Oper.Max_token_cache,
			"cache_list":      setSliceAamAuthorizationJwtCacheOperOperCacheList(ret.DtAamAuthorizationJwtCacheOper.Oper.CacheList),
			"audience":        ret.DtAamAuthorizationJwtCacheOper.Oper.Audience,
		},
	}
}

func setSliceAamAuthorizationJwtCacheOperOperCacheList(d []edpt.AamAuthorizationJwtCacheOperOperCacheList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cache_id"] = item.CacheId
		in["issuer"] = item.Issuer
		in["subject"] = item.Subject
		in["audience"] = item.Audience
		in["client_ip"] = item.ClientIp
		in["ttl"] = item.Ttl
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthorizationJwtCacheOperOper(d []interface{}) edpt.AamAuthorizationJwtCacheOperOper {

	count1 := len(d)
	var ret edpt.AamAuthorizationJwtCacheOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Token_cached = in["token_cached"].(int)
		ret.Token_cache_hit = in["token_cache_hit"].(int)
		ret.Max_token_cache = in["max_token_cache"].(int)
		ret.CacheList = getSliceAamAuthorizationJwtCacheOperOperCacheList(in["cache_list"].([]interface{}))
		ret.Audience = in["audience"].(string)
	}
	return ret
}

func getSliceAamAuthorizationJwtCacheOperOperCacheList(d []interface{}) []edpt.AamAuthorizationJwtCacheOperOperCacheList {

	count1 := len(d)
	ret := make([]edpt.AamAuthorizationJwtCacheOperOperCacheList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthorizationJwtCacheOperOperCacheList
		oi.CacheId = in["cache_id"].(int)
		oi.Issuer = in["issuer"].(string)
		oi.Subject = in["subject"].(string)
		oi.Audience = in["audience"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.Ttl = in["ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthorizationJwtCacheOper(d *schema.ResourceData) edpt.AamAuthorizationJwtCacheOper {
	var ret edpt.AamAuthorizationJwtCacheOper

	ret.Oper = getObjectAamAuthorizationJwtCacheOperOper(d.Get("oper").([]interface{}))
	return ret
}
