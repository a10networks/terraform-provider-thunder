package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkArpTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_arp_timeout`: ARP entry timeout\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkArpTimeoutCreate,
		UpdateContext: resourceNetworkArpTimeoutUpdate,
		ReadContext:   resourceNetworkArpTimeoutRead,
		DeleteContext: resourceNetworkArpTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "ARP entry timeout",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkArpTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkArpTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkArpTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkArpTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkArpTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkArpTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkArpTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkArpTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkArpTimeout(d *schema.ResourceData) edpt.NetworkArpTimeout {
	var ret edpt.NetworkArpTimeout
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
