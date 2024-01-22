package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTwampResponderIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_twamp_responder_ipv6`: Configure TWAMP responder\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkTwampResponderIpv6Create,
		UpdateContext: resourceNetworkTwampResponderIpv6Update,
		ReadContext:   resourceNetworkTwampResponderIpv6Read,
		DeleteContext: resourceNetworkTwampResponderIpv6Delete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v6_acl_name": {
				Type: schema.TypeString, Optional: true, Description: "Apply an access list (Named Access List)",
			},
		},
	}
}
func resourceNetworkTwampResponderIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkTwampResponderIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceNetworkTwampResponderIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkTwampResponderIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceNetworkTwampResponderIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkTwampResponderIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkTwampResponderIpv6(d *schema.ResourceData) edpt.NetworkTwampResponderIpv6 {
	var ret edpt.NetworkTwampResponderIpv6
	//omit uuid
	ret.Inst.V6AclName = d.Get("v6_acl_name").(string)
	return ret
}
