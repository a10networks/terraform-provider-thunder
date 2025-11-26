package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDomainListInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_domain_list_info_oper`: Operational Status for the object domain-list-info\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDomainListInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"resolved": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_number": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_ipv6": {
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

func resourceSlbDomainListInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDomainListInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDomainListInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDomainListInfoOperOper := setObjectSlbDomainListInfoOperOper(res)
		d.Set("oper", SlbDomainListInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDomainListInfoOperOper(ret edpt.DataSlbDomainListInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"domain_list": setSliceSlbDomainListInfoOperOperDomainList(ret.DtSlbDomainListInfoOper.Oper.DomainList),
		},
	}
}

func setSliceSlbDomainListInfoOperOperDomainList(d []edpt.SlbDomainListInfoOperOperDomainList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["domain_name"] = item.DomainName
		in["resolved"] = item.Resolved
		in["hit_count"] = item.HitCount
		in["ip_number"] = item.IpNumber
		in["ip"] = item.Ip
		in["is_ipv6"] = item.IsIpv6
		result = append(result, in)
	}
	return result
}

func getObjectSlbDomainListInfoOperOper(d []interface{}) edpt.SlbDomainListInfoOperOper {

	count1 := len(d)
	var ret edpt.SlbDomainListInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainList = getSliceSlbDomainListInfoOperOperDomainList(in["domain_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbDomainListInfoOperOperDomainList(d []interface{}) []edpt.SlbDomainListInfoOperOperDomainList {

	count1 := len(d)
	ret := make([]edpt.SlbDomainListInfoOperOperDomainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDomainListInfoOperOperDomainList
		oi.Name = in["name"].(string)
		oi.DomainName = in["domain_name"].(string)
		oi.Resolved = in["resolved"].(int)
		oi.HitCount = in["hit_count"].(int)
		oi.IpNumber = in["ip_number"].(int)
		oi.Ip = in["ip"].(string)
		oi.IsIpv6 = in["is_ipv6"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDomainListInfoOper(d *schema.ResourceData) edpt.SlbDomainListInfoOper {
	var ret edpt.SlbDomainListInfoOper

	ret.Oper = getObjectSlbDomainListInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
