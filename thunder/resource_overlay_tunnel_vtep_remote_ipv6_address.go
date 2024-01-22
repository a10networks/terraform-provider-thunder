package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpv6Address() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ipv6_address`: Configure remote tunnel end point parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpv6AddressCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpv6AddressUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpv6AddressRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpv6AddressDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "IPv6 Address of the remote VTEP",
			},
			"use_lif": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
						},
						"lif": {
							Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
		},
	}
}
func resourceOverlayTunnelVtepRemoteIpv6AddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6Address(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6Address(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpv6AddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6Address(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6Address(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectOverlayTunnelVtepRemoteIpv6AddressUseLif1084(d []interface{}) edpt.OverlayTunnelVtepRemoteIpv6AddressUseLif1084 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressUseLif1084
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(string)
		ret.Lif = in["lif"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointOverlayTunnelVtepRemoteIpv6Address(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpv6Address {
	var ret edpt.OverlayTunnelVtepRemoteIpv6Address
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.UseLif = getObjectOverlayTunnelVtepRemoteIpv6AddressUseLif1084(d.Get("use_lif").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
