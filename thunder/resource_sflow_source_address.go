package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowSourceAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_source_address`: Specify source address of sFlow packet\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowSourceAddressCreate,
		UpdateContext: resourceSflowSourceAddressUpdate,
		ReadContext:   resourceSflowSourceAddressRead,
		DeleteContext: resourceSflowSourceAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type: schema.TypeString, Optional: true, Description: "Source IPv4 address",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Source IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowSourceAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSourceAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSourceAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowSourceAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowSourceAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSourceAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSourceAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowSourceAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowSourceAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSourceAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSourceAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowSourceAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSourceAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSourceAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowSourceAddress(d *schema.ResourceData) edpt.SflowSourceAddress {
	var ret edpt.SflowSourceAddress
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	return ret
}
