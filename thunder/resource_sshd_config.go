package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSshdConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sshd_config`: SSHD service config\n\n__PLACEHOLDER__",
		CreateContext: resourceSshdConfigCreate,
		UpdateContext: resourceSshdConfigUpdate,
		ReadContext:   resourceSshdConfigRead,
		DeleteContext: resourceSshdConfigDelete,

		Schema: map[string]*schema.Schema{
			"disable_agent_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable sshd agent forwarding",
			},
			"disable_tcp_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable sshd tcp forwarding",
			},
			"disable_x11_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable sshd x11 forwarding",
			},
			"tcp_port": {
				Type: schema.TypeInt, Optional: true, Default: 22, Description: "ssh port number (Available port is 22, 1025-64999)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSshdConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshdConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSshdConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceSshdConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshdConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSshdConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceSshdConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshdConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSshdConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshdConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshdConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSshdConfig(d *schema.ResourceData) edpt.SshdConfig {
	var ret edpt.SshdConfig
	ret.Inst.DisableAgentForwarding = d.Get("disable_agent_forwarding").(int)
	ret.Inst.DisableTcpForwarding = d.Get("disable_tcp_forwarding").(int)
	ret.Inst.DisableX11Forwarding = d.Get("disable_x11_forwarding").(int)
	ret.Inst.TcpPort = d.Get("tcp_port").(int)
	//omit uuid
	return ret
}
