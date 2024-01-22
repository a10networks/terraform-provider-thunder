package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_address`: Transparent mode IP Address\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAddressCreate,
		UpdateContext: resourceIpAddressUpdate,
		ReadContext:   resourceIpAddressRead,
		DeleteContext: resourceIpAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"ip_mask": {
				Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpAddress(d *schema.ResourceData) edpt.IpAddress {
	var ret edpt.IpAddress
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.IpMask = d.Get("ip_mask").(string)
	//omit uuid
	return ret
}
