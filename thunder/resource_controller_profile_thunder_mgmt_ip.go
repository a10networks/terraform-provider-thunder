package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProfileThunderMgmtIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_profile_thunder_mgmt_ip`: thunder management ip address\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerProfileThunderMgmtIpCreate,
		UpdateContext: resourceControllerProfileThunderMgmtIpUpdate,
		ReadContext:   resourceControllerProfileThunderMgmtIpRead,
		DeleteContext: resourceControllerProfileThunderMgmtIpDelete,

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
func resourceControllerProfileThunderMgmtIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileThunderMgmtIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileThunderMgmtIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileThunderMgmtIpRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerProfileThunderMgmtIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileThunderMgmtIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileThunderMgmtIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileThunderMgmtIpRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerProfileThunderMgmtIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileThunderMgmtIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileThunderMgmtIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerProfileThunderMgmtIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileThunderMgmtIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileThunderMgmtIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointControllerProfileThunderMgmtIp(d *schema.ResourceData) edpt.ControllerProfileThunderMgmtIp {
	var ret edpt.ControllerProfileThunderMgmtIp
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	//omit uuid
	return ret
}
