package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsPersistentCacheOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dns_persistent_cache_oper`: Operational Status for the object dns-persistent-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDnsPersistentCacheOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache_entry": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dnssec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cache_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cache_class": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"q_length": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"r_length": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_cache": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbDnsPersistentCacheOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsPersistentCacheOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsPersistentCacheOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDnsPersistentCacheOperOper := setObjectSlbDnsPersistentCacheOperOper(res)
		d.Set("oper", SlbDnsPersistentCacheOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDnsPersistentCacheOperOper(ret edpt.DataSlbDnsPersistentCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cache_entry": setSliceSlbDnsPersistentCacheOperOperCacheEntry(ret.DtSlbDnsPersistentCacheOper.Oper.CacheEntry),
			"total_cache": ret.DtSlbDnsPersistentCacheOper.Oper.Total_cache,
		},
	}
}

func setSliceSlbDnsPersistentCacheOperOperCacheEntry(d []edpt.SlbDnsPersistentCacheOperOperCacheEntry) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain"] = item.Domain
		in["dnssec"] = item.Dnssec
		in["cache_type"] = item.Cache_type
		in["cache_class"] = item.Cache_class
		in["q_length"] = item.Q_length
		in["r_length"] = item.R_length
		in["ttl"] = item.Ttl
		result = append(result, in)
	}
	return result
}

func getObjectSlbDnsPersistentCacheOperOper(d []interface{}) edpt.SlbDnsPersistentCacheOperOper {

	count1 := len(d)
	var ret edpt.SlbDnsPersistentCacheOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheEntry = getSliceSlbDnsPersistentCacheOperOperCacheEntry(in["cache_entry"].([]interface{}))
		ret.Total_cache = in["total_cache"].(int)
	}
	return ret
}

func getSliceSlbDnsPersistentCacheOperOperCacheEntry(d []interface{}) []edpt.SlbDnsPersistentCacheOperOperCacheEntry {

	count1 := len(d)
	ret := make([]edpt.SlbDnsPersistentCacheOperOperCacheEntry, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsPersistentCacheOperOperCacheEntry
		oi.Domain = in["domain"].(string)
		oi.Dnssec = in["dnssec"].(int)
		oi.Cache_type = in["cache_type"].(int)
		oi.Cache_class = in["cache_class"].(int)
		oi.Q_length = in["q_length"].(int)
		oi.R_length = in["r_length"].(int)
		oi.Ttl = in["ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDnsPersistentCacheOper(d *schema.ResourceData) edpt.SlbDnsPersistentCacheOper {
	var ret edpt.SlbDnsPersistentCacheOper

	ret.Oper = getObjectSlbDnsPersistentCacheOperOper(d.Get("oper").([]interface{}))
	return ret
}
