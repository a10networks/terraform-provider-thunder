package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListIpv6InternetHostList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip_threat_list_ipv6_internet_host_list`: List of IPv6 Class Lists for Internet Host IPv6 Threat List\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpThreatListIpv6InternetHostListCreate,
		UpdateContext: resourceSystemIpThreatListIpv6InternetHostListUpdate,
		ReadContext:   resourceSystemIpThreatListIpv6InternetHostListRead,
		DeleteContext: resourceSystemIpThreatListIpv6InternetHostListDelete,

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
func resourceSystemIpThreatListIpv6InternetHostListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6InternetHostListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6InternetHostList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv6InternetHostListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpThreatListIpv6InternetHostListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6InternetHostListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6InternetHostList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv6InternetHostListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpThreatListIpv6InternetHostListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6InternetHostListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6InternetHostList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpThreatListIpv6InternetHostListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv6InternetHostListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv6InternetHostList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIpThreatListIpv6InternetHostListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv6InternetHostListClassListCfg {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6InternetHostListClassListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6InternetHostListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatListIpv6InternetHostList(d *schema.ResourceData) edpt.SystemIpThreatListIpv6InternetHostList {
	var ret edpt.SystemIpThreatListIpv6InternetHostList
	ret.Inst.ClassListCfg = getSliceSystemIpThreatListIpv6InternetHostListClassListCfg(d.Get("class_list_cfg").([]interface{}))
	//omit uuid
	return ret
}
