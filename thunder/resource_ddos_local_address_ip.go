package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosLocalAddressIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_local_address_ip`: Configure DDoS ipv4 address\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosLocalAddressIpCreate,
		UpdateContext: resourceDdosLocalAddressIpUpdate,
		ReadContext:   resourceDdosLocalAddressIpRead,
		DeleteContext: resourceDdosLocalAddressIpDelete,

		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type: schema.TypeString, Required: true, Description: "DDoS IPv4 Address for syn cookie usage",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosLocalAddressIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosLocalAddressIpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosLocalAddressIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosLocalAddressIpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosLocalAddressIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosLocalAddressIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLocalAddressIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLocalAddressIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosLocalAddressIp(d *schema.ResourceData) edpt.DdosLocalAddressIp {
	var ret edpt.DdosLocalAddressIp
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	//omit uuid
	return ret
}
