package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDomainIpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_domain_ip_oper`: Operational Status for the object domain-ip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListDomainIpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_list": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceCgnv6LsnRuleListDomainIpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDomainIpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDomainIpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListDomainIpOperOper := setObjectCgnv6LsnRuleListDomainIpOperOper(res)
		d.Set("oper", Cgnv6LsnRuleListDomainIpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListDomainIpOperOper(ret edpt.DataCgnv6LsnRuleListDomainIpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_list": setSliceCgnv6LsnRuleListDomainIpOperOperIpList(ret.DtCgnv6LsnRuleListDomainIpOper.Oper.IpList),
		},
	}
}

func setSliceCgnv6LsnRuleListDomainIpOperOperIpList(d []edpt.Cgnv6LsnRuleListDomainIpOperOperIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain"] = item.Domain
		in["domain_list"] = item.DomainList
		in["ip_address"] = item.IpAddress
		in["ttl"] = item.Ttl
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6LsnRuleListDomainIpOperOper(d []interface{}) edpt.Cgnv6LsnRuleListDomainIpOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainIpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpList = getSliceCgnv6LsnRuleListDomainIpOperOperIpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6LsnRuleListDomainIpOperOperIpList(d []interface{}) []edpt.Cgnv6LsnRuleListDomainIpOperOperIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainIpOperOperIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainIpOperOperIpList
		oi.Domain = in["domain"].(string)
		oi.DomainList = in["domain_list"].(string)
		oi.IpAddress = in["ip_address"].(string)
		oi.Ttl = in["ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRuleListDomainIpOper(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDomainIpOper {
	var ret edpt.Cgnv6LsnRuleListDomainIpOper

	ret.Oper = getObjectCgnv6LsnRuleListDomainIpOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
