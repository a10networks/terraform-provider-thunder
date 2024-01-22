package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDnsPrimary() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_dns_primary`: Primary DNS server\n\n__PLACEHOLDER__",
		CreateContext: resourceIpDnsPrimaryCreate,
		UpdateContext: resourceIpDnsPrimaryUpdate,
		ReadContext:   resourceIpDnsPrimaryRead,
		DeleteContext: resourceIpDnsPrimaryDelete,

		Schema: map[string]*schema.Schema{
			"ip_v4_addr": {
				Type: schema.TypeString, Optional: true, Description: "DNS server address",
			},
			"ip_v6_addr": {
				Type: schema.TypeString, Optional: true, Description: "DNS server address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpDnsPrimaryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsPrimaryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsPrimary(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDnsPrimaryRead(ctx, d, meta)
	}
	return diags
}

func resourceIpDnsPrimaryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsPrimaryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsPrimary(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDnsPrimaryRead(ctx, d, meta)
	}
	return diags
}
func resourceIpDnsPrimaryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsPrimaryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsPrimary(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpDnsPrimaryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDnsPrimaryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDnsPrimary(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpDnsPrimary(d *schema.ResourceData) edpt.IpDnsPrimary {
	var ret edpt.IpDnsPrimary
	ret.Inst.IpV4Addr = d.Get("ip_v4_addr").(string)
	ret.Inst.IpV6Addr = d.Get("ip_v6_addr").(string)
	//omit uuid
	return ret
}
