package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowAgentAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_agent_address`: Configure agent address used in sFlow datagram, default use management IP address\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowAgentAddressCreate,
		UpdateContext: resourceSflowAgentAddressUpdate,
		ReadContext:   resourceSflowAgentAddressRead,
		DeleteContext: resourceSflowAgentAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type: schema.TypeString, Optional: true, Description: "Configure sFlow agent IP address",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Configure sFlow agent IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowAgentAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowAgentAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowAgentAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowAgentAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowAgentAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowAgentAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowAgentAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowAgentAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowAgentAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowAgentAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowAgentAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowAgentAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowAgentAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowAgentAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowAgentAddress(d *schema.ResourceData) edpt.SflowAgentAddress {
	var ret edpt.SflowAgentAddress
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	return ret
}
