package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpmiIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipmi_ip`: Set IP address for IPMI interface\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpmiIpCreate,
		UpdateContext: resourceSystemIpmiIpUpdate,
		ReadContext:   resourceSystemIpmiIpRead,
		DeleteContext: resourceSystemIpmiIpDelete,

		Schema: map[string]*schema.Schema{
			"default_gateway": {
				Type: schema.TypeString, Optional: true, Description: "Default gateway address",
			},
			"ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"ipv4_netmask": {
				Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
			},
		},
	}
}
func resourceSystemIpmiIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiIpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpmiIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiIpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpmiIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpmiIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIpmiIp(d *schema.ResourceData) edpt.SystemIpmiIp {
	var ret edpt.SystemIpmiIp
	ret.Inst.DefaultGateway = d.Get("default_gateway").(string)
	ret.Inst.Ipv4Address = d.Get("ipv4_address").(string)
	ret.Inst.Ipv4Netmask = d.Get("ipv4_netmask").(string)
	return ret
}
