package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListIpv4DestList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip_threat_list_ipv4_dest_list`: List of IPv4 Class Lists for Destination IPv4 Threat List\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpThreatListIpv4DestListCreate,
		UpdateContext: resourceSystemIpThreatListIpv4DestListUpdate,
		ReadContext:   resourceSystemIpThreatListIpv4DestListRead,
		DeleteContext: resourceSystemIpThreatListIpv4DestListDelete,

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
func resourceSystemIpThreatListIpv4DestListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4DestListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4DestList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv4DestListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpThreatListIpv4DestListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4DestListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4DestList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv4DestListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpThreatListIpv4DestListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4DestListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4DestList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpThreatListIpv4DestListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4DestListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4DestList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIpThreatListIpv4DestListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv4DestListClassListCfg {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4DestListClassListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4DestListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatListIpv4DestList(d *schema.ResourceData) edpt.SystemIpThreatListIpv4DestList {
	var ret edpt.SystemIpThreatListIpv4DestList
	ret.Inst.ClassListCfg = getSliceSystemIpThreatListIpv4DestListClassListCfg(d.Get("class_list_cfg").([]interface{}))
	//omit uuid
	return ret
}
