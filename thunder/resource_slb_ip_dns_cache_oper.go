package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIpDnsCacheOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ip_dns_cache_oper`: Operational Status for the object ip-dns-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbIpDnsCacheOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interval": {
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

func resourceSlbIpDnsCacheOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIpDnsCacheOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIpDnsCacheOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbIpDnsCacheOperOper := setObjectSlbIpDnsCacheOperOper(res)
		d.Set("oper", SlbIpDnsCacheOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbIpDnsCacheOperOper(ret edpt.DataSlbIpDnsCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"domain_ip_list": setSliceSlbIpDnsCacheOperOperDomainIpList(ret.DtSlbIpDnsCacheOper.Oper.DomainIpList),
		},
	}
}

func setSliceSlbIpDnsCacheOperOperDomainIpList(d []edpt.SlbIpDnsCacheOperOperDomainIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain"] = item.Domain
		in["address"] = item.Address
		in["ttl"] = item.Ttl
		in["interval"] = item.Interval
		result = append(result, in)
	}
	return result
}

func getObjectSlbIpDnsCacheOperOper(d []interface{}) edpt.SlbIpDnsCacheOperOper {

	count1 := len(d)
	var ret edpt.SlbIpDnsCacheOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainIpList = getSliceSlbIpDnsCacheOperOperDomainIpList(in["domain_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbIpDnsCacheOperOperDomainIpList(d []interface{}) []edpt.SlbIpDnsCacheOperOperDomainIpList {

	count1 := len(d)
	ret := make([]edpt.SlbIpDnsCacheOperOperDomainIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbIpDnsCacheOperOperDomainIpList
		oi.Domain = in["domain"].(string)
		oi.Address = in["address"].(string)
		oi.Ttl = in["ttl"].(int)
		oi.Interval = in["interval"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbIpDnsCacheOper(d *schema.ResourceData) edpt.SlbIpDnsCacheOper {
	var ret edpt.SlbIpDnsCacheOper

	ret.Oper = getObjectSlbIpDnsCacheOperOper(d.Get("oper").([]interface{}))
	return ret
}
