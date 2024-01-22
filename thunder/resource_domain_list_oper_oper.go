package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainListOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_domain_list_oper_oper`: Operational Status for the object domain-list-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDomainListOperOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_list_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"config_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_entry_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"binding_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"binding_info_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain_group_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"domain_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_match_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_transfer_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_transfer_port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_transfer_interval": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"match_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"domain_list_name_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"domain_list_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"binding_info": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDomainListOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainListOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainListOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DomainListOperOperOper := setObjectDomainListOperOperOper(res)
		d.Set("oper", DomainListOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDomainListOperOperOper(ret edpt.DataDomainListOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"domain_list":             setSliceDomainListOperOperOperDomainList(ret.DtDomainListOperOper.Oper.DomainList),
			"domain_entries":          setSliceDomainListOperOperOperDomainEntries(ret.DtDomainListOperOper.Oper.DomainEntries),
			"displayed_count":         ret.DtDomainListOperOper.Oper.DisplayedCount,
			"match_type":              ret.DtDomainListOperOper.Oper.MatchType,
			"domain_list_name_filter": ret.DtDomainListOperOper.Oper.DomainListNameFilter,
			"domain_list_type":        ret.DtDomainListOperOper.Oper.DomainListType,
			"binding_info":            ret.DtDomainListOperOper.Oper.BindingInfo,
		},
	}
}

func setSliceDomainListOperOperOperDomainList(d []edpt.DomainListOperOperOperDomainList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_list_name"] = item.DomainListName
		in["config_type"] = item.ConfigType
		in["total_entry_num"] = item.TotalEntryNum
		in["binding_num"] = item.BindingNum
		in["binding_info_list"] = setSliceDomainListOperOperOperDomainListBindingInfoList(item.BindingInfoList)
		result = append(result, in)
	}
	return result
}

func setSliceDomainListOperOperOperDomainListBindingInfoList(d []edpt.DomainListOperOperOperDomainListBindingInfoList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_group_name"] = item.DomainGroupName
		result = append(result, in)
	}
	return result
}

func setSliceDomainListOperOperOperDomainEntries(d []edpt.DomainListOperOperOperDomainEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain_match_type"] = item.DomainMatchType
		in["domain_name"] = item.DomainName
		in["zone_transfer_ip"] = item.ZoneTransferIp
		in["zone_transfer_port"] = item.ZoneTransferPort
		in["zone_transfer_interval"] = item.ZoneTransferInterval
		result = append(result, in)
	}
	return result
}

func getObjectDomainListOperOperOper(d []interface{}) edpt.DomainListOperOperOper {

	count1 := len(d)
	var ret edpt.DomainListOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainList = getSliceDomainListOperOperOperDomainList(in["domain_list"].([]interface{}))
		ret.DomainEntries = getSliceDomainListOperOperOperDomainEntries(in["domain_entries"].([]interface{}))
		ret.DisplayedCount = in["displayed_count"].(int)
		ret.MatchType = in["match_type"].(string)
		ret.DomainListNameFilter = in["domain_list_name_filter"].(string)
		ret.DomainListType = in["domain_list_type"].(string)
		ret.BindingInfo = in["binding_info"].(int)
	}
	return ret
}

func getSliceDomainListOperOperOperDomainList(d []interface{}) []edpt.DomainListOperOperOperDomainList {

	count1 := len(d)
	ret := make([]edpt.DomainListOperOperOperDomainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListOperOperOperDomainList
		oi.DomainListName = in["domain_list_name"].(string)
		oi.ConfigType = in["config_type"].(string)
		oi.TotalEntryNum = in["total_entry_num"].(int)
		oi.BindingNum = in["binding_num"].(int)
		oi.BindingInfoList = getSliceDomainListOperOperOperDomainListBindingInfoList(in["binding_info_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainListOperOperOperDomainListBindingInfoList(d []interface{}) []edpt.DomainListOperOperOperDomainListBindingInfoList {

	count1 := len(d)
	ret := make([]edpt.DomainListOperOperOperDomainListBindingInfoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListOperOperOperDomainListBindingInfoList
		oi.DomainGroupName = in["domain_group_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainListOperOperOperDomainEntries(d []interface{}) []edpt.DomainListOperOperOperDomainEntries {

	count1 := len(d)
	ret := make([]edpt.DomainListOperOperOperDomainEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListOperOperOperDomainEntries
		oi.DomainMatchType = in["domain_match_type"].(string)
		oi.DomainName = in["domain_name"].(string)
		oi.ZoneTransferIp = in["zone_transfer_ip"].(string)
		oi.ZoneTransferPort = in["zone_transfer_port"].(string)
		oi.ZoneTransferInterval = in["zone_transfer_interval"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDomainListOperOper(d *schema.ResourceData) edpt.DomainListOperOper {
	var ret edpt.DomainListOperOper

	ret.Oper = getObjectDomainListOperOperOper(d.Get("oper").([]interface{}))
	return ret
}
