package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfileThunderMgmtIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_profile_thunder_mgmt_ip`: thunder management ip address\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerProfileThunderMgmtIpCreate,
		UpdateContext: resourceHarmonyControllerProfileThunderMgmtIpUpdate,
		ReadContext:   resourceHarmonyControllerProfileThunderMgmtIpRead,
		DeleteContext: resourceHarmonyControllerProfileThunderMgmtIpDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type: schema.TypeString, Optional: true, Description: "IP address (IPv4 address)",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address for the host",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceHarmonyControllerProfileThunderMgmtIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileThunderMgmtIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileThunderMgmtIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileThunderMgmtIpRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerProfileThunderMgmtIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileThunderMgmtIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileThunderMgmtIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileThunderMgmtIpRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerProfileThunderMgmtIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileThunderMgmtIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileThunderMgmtIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerProfileThunderMgmtIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileThunderMgmtIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileThunderMgmtIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHarmonyControllerProfileThunderMgmtIp(d *schema.ResourceData) edpt.HarmonyControllerProfileThunderMgmtIp {
	var ret edpt.HarmonyControllerProfileThunderMgmtIp
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	//omit uuid
	return ret
}
