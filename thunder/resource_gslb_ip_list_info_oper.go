package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbIpListInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_ip_list_info_oper`: Operational Status for the object ip-list-info\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbIpListInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_lists": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listname": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_entries_in_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filename": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"file_lines": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"errors": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipaddr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipmask": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"group_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipaddr_filter": {
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

func resourceGslbIpListInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbIpListInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbIpListInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbIpListInfoOperOper := setObjectGslbIpListInfoOperOper(res)
		d.Set("oper", GslbIpListInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbIpListInfoOperOper(ret edpt.DataGslbIpListInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_lists": setSliceGslbIpListInfoOperOperIpLists(ret.DtGslbIpListInfoOper.Oper.IpLists),
		},
	}
}

func setSliceGslbIpListInfoOperOperIpLists(d []edpt.GslbIpListInfoOperOperIpLists) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["listname"] = item.Listname
		in["total_entries_in_list"] = item.TotalEntriesInList
		in["filename"] = item.Filename
		in["file_lines"] = item.FileLines
		in["errors"] = item.Errors
		in["ipaddr"] = item.Ipaddr
		in["ipmask"] = item.Ipmask
		in["group_id"] = item.GroupId
		in["last"] = item.Last
		in["hits"] = item.Hits
		in["type"] = item.Type
		in["ipaddr_filter"] = item.IpaddrFilter
		result = append(result, in)
	}
	return result
}

func getObjectGslbIpListInfoOperOper(d []interface{}) edpt.GslbIpListInfoOperOper {

	count1 := len(d)
	var ret edpt.GslbIpListInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpLists = getSliceGslbIpListInfoOperOperIpLists(in["ip_lists"].([]interface{}))
	}
	return ret
}

func getSliceGslbIpListInfoOperOperIpLists(d []interface{}) []edpt.GslbIpListInfoOperOperIpLists {

	count1 := len(d)
	ret := make([]edpt.GslbIpListInfoOperOperIpLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbIpListInfoOperOperIpLists
		oi.Listname = in["listname"].(string)
		oi.TotalEntriesInList = in["total_entries_in_list"].(int)
		oi.Filename = in["filename"].(string)
		oi.FileLines = in["file_lines"].(int)
		oi.Errors = in["errors"].(int)
		oi.Ipaddr = in["ipaddr"].(string)
		oi.Ipmask = in["ipmask"].(string)
		oi.GroupId = in["group_id"].(int)
		oi.Last = in["last"].(string)
		oi.Hits = in["hits"].(int)
		oi.Type = in["type"].(string)
		oi.IpaddrFilter = in["ipaddr_filter"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbIpListInfoOper(d *schema.ResourceData) edpt.GslbIpListInfoOper {
	var ret edpt.GslbIpListInfoOper

	ret.Oper = getObjectGslbIpListInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
