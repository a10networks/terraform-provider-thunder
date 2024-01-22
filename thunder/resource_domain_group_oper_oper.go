package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainGroupOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_domain_group_oper_oper`: Operational Status for the object domain-group-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDomainGroupOperOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_groups": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_entry_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"binding_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"domain_group_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_match_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_list_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"domain_group_name_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"domain_name_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDomainGroupOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainGroupOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainGroupOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DomainGroupOperOperOper := setObjectDomainGroupOperOperOper(res)
		d.Set("oper", DomainGroupOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDomainGroupOperOperOper(ret edpt.DataDomainGroupOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"domain_groups":            setSliceDomainGroupOperOperOperDomainGroups(ret.DtDomainGroupOperOper.Oper.DomainGroups),
			"domain_group_entries":     setSliceDomainGroupOperOperOperDomainGroupEntries(ret.DtDomainGroupOperOper.Oper.DomainGroupEntries),
			"displayed_count":          ret.DtDomainGroupOperOper.Oper.DisplayedCount,
			"domain_group_name_filter": ret.DtDomainGroupOperOper.Oper.DomainGroupNameFilter,
			"domain_name_filter":       ret.DtDomainGroupOperOper.Oper.DomainNameFilter,
		},
	}
}

func setSliceDomainGroupOperOperOperDomainGroups(d []edpt.DomainGroupOperOperOperDomainGroups) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_group_name"] = item.DomainGroupName
		in["total_entry_num"] = item.TotalEntryNum
		in["binding_num"] = item.BindingNum
		result = append(result, in)
	}
	return result
}

func setSliceDomainGroupOperOperOperDomainGroupEntries(d []edpt.DomainGroupOperOperOperDomainGroupEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_match_type"] = item.DomainMatchType
		in["domain_name"] = item.DomainName
		in["domain_list_name"] = item.DomainListName
		in["action"] = item.Action
		in["hit_count"] = item.HitCount
		result = append(result, in)
	}
	return result
}

func getObjectDomainGroupOperOperOper(d []interface{}) edpt.DomainGroupOperOperOper {

	count1 := len(d)
	var ret edpt.DomainGroupOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainGroups = getSliceDomainGroupOperOperOperDomainGroups(in["domain_groups"].([]interface{}))
		ret.DomainGroupEntries = getSliceDomainGroupOperOperOperDomainGroupEntries(in["domain_group_entries"].([]interface{}))
		ret.DisplayedCount = in["displayed_count"].(int)
		ret.DomainGroupNameFilter = in["domain_group_name_filter"].(string)
		ret.DomainNameFilter = in["domain_name_filter"].(string)
	}
	return ret
}

func getSliceDomainGroupOperOperOperDomainGroups(d []interface{}) []edpt.DomainGroupOperOperOperDomainGroups {

	count1 := len(d)
	ret := make([]edpt.DomainGroupOperOperOperDomainGroups, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainGroupOperOperOperDomainGroups
		oi.DomainGroupName = in["domain_group_name"].(string)
		oi.TotalEntryNum = in["total_entry_num"].(int)
		oi.BindingNum = in["binding_num"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainGroupOperOperOperDomainGroupEntries(d []interface{}) []edpt.DomainGroupOperOperOperDomainGroupEntries {

	count1 := len(d)
	ret := make([]edpt.DomainGroupOperOperOperDomainGroupEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainGroupOperOperOperDomainGroupEntries
		oi.DomainMatchType = in["domain_match_type"].(string)
		oi.DomainName = in["domain_name"].(string)
		oi.DomainListName = in["domain_list_name"].(string)
		oi.Action = in["action"].(string)
		oi.HitCount = in["hit_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDomainGroupOperOper(d *schema.ResourceData) edpt.DomainGroupOperOper {
	var ret edpt.DomainGroupOperOper

	ret.Oper = getObjectDomainGroupOperOperOper(d.Get("oper").([]interface{}))
	return ret
}
