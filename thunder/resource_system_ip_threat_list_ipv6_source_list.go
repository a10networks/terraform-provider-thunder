package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListIpv6SourceList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip_threat_list_ipv6_source_list`: List of IPv6 Class Lists for Source IPv6 Threat List\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpThreatListIpv6SourceListCreate,
		UpdateContext: resourceSystemIpThreatListIpv6SourceListUpdate,
		ReadContext:   resourceSystemIpThreatListIpv6SourceListRead,
		DeleteContext: resourceSystemIpThreatListIpv6SourceListDelete,

		Schema: map[string]*schema.Schema{
			"class_list_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
						},
						"ip_threat_action_tmpl": {
							Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemIpThreatListIpv6SourceListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6SourceListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6SourceList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv6SourceListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpThreatListIpv6SourceListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6SourceListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6SourceList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv6SourceListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpThreatListIpv6SourceListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6SourceListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6SourceList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpThreatListIpv6SourceListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6SourceListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6SourceList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIpThreatListIpv6SourceListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv6SourceListClassListCfg {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6SourceListClassListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6SourceListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatListIpv6SourceList(d *schema.ResourceData) edpt.SystemIpThreatListIpv6SourceList {
	var ret edpt.SystemIpThreatListIpv6SourceList
	ret.Inst.ClassListCfg = getSliceSystemIpThreatListIpv6SourceListClassListCfg(d.Get("class_list_cfg").([]interface{}))
	//omit uuid
	return ret
}
