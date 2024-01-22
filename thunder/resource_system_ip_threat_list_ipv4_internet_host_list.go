package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListIpv4InternetHostList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip_threat_list_ipv4_internet_host_list`: List of IPv4 Class Lists for Internet Host IPv4 Threat List\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpThreatListIpv4InternetHostListCreate,
		UpdateContext: resourceSystemIpThreatListIpv4InternetHostListUpdate,
		ReadContext:   resourceSystemIpThreatListIpv4InternetHostListRead,
		DeleteContext: resourceSystemIpThreatListIpv4InternetHostListDelete,

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
func resourceSystemIpThreatListIpv4InternetHostListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4InternetHostListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4InternetHostList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv4InternetHostListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpThreatListIpv4InternetHostListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4InternetHostListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4InternetHostList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListIpv4InternetHostListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpThreatListIpv4InternetHostListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4InternetHostListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4InternetHostList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpThreatListIpv4InternetHostListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListIpv4InternetHostListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListIpv4InternetHostList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIpThreatListIpv4InternetHostListClassListCfg(d []interface{}) []edpt.SystemIpThreatListIpv4InternetHostListClassListCfg {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4InternetHostListClassListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4InternetHostListClassListCfg
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatListIpv4InternetHostList(d *schema.ResourceData) edpt.SystemIpThreatListIpv4InternetHostList {
	var ret edpt.SystemIpThreatListIpv4InternetHostList
	ret.Inst.ClassListCfg = getSliceSystemIpThreatListIpv4InternetHostListClassListCfg(d.Get("class_list_cfg").([]interface{}))
	//omit uuid
	return ret
}
