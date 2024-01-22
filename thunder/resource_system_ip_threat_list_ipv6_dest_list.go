package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListIpv6DestList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip_threat_list_ipv6_dest_list`: List of IPv6 Class Lists for Destination IPv6 Threat List\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpThreatListIpv6DestListCreate,
		UpdateContext: resourceSystemIpThreatListIpv6DestListUpdate,
		ReadContext:   resourceSystemIpThreatListIpv6DestListRead,
		DeleteContext: resourceSystemIpThreatListIpv6DestListDelete,

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
func resourceSystemIpThreatListIpv6DestListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6DestListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6DestList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv6DestListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpThreatListIpv6DestListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6DestListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6DestList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv6DestListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpThreatListIpv6DestListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6DestListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6DestList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpThreatListIpv6DestListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6DestListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6DestList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIpThreatListIpv6DestListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv6DestListClassListCfg {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6DestListClassListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6DestListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatListIpv6DestList(d *schema.ResourceData) edpt.SystemIpThreatListIpv6DestList {
	var ret edpt.SystemIpThreatListIpv6DestList
	ret.Inst.ClassListCfg = getSliceSystemIpThreatListIpv6DestListClassListCfg(d.Get("class_list_cfg").([]interface{}))
	//omit uuid
	return ret
}
