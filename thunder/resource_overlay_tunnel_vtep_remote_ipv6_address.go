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
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Name of the class-list",
			},
			"encap": {
				Type: schema.TypeString, Optional: true, Description: "'vxlan': Tunnel Encapsulation Type is VXLAN;",
			},
			"gre_keepalive": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retry_time": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Keepalive retry interval in seconds",
						},
						"retry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Keepalive multiplier",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "IPv6 Address of the remote VTEP",
			},
			"use_gre_key": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gre_key": {
							Type: schema.TypeInt, Optional: true, Description: "key",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
			"vni_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"segment": {
							Type: schema.TypeInt, Required: true, Description: "VNI configured for the remote VTEP",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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

func getObjectOverlayTunnelVtepRemoteIpv6AddressGreKeepalive1164(d []interface{}) edpt.OverlayTunnelVtepRemoteIpv6AddressGreKeepalive1164 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressGreKeepalive1164
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetryTime = in["retry_time"].(int)
		ret.RetryCount = in["retry_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpv6AddressUseGreKey1165(d []interface{}) edpt.OverlayTunnelVtepRemoteIpv6AddressUseGreKey1165 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressUseGreKey1165
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreKey = in["gre_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpv6AddressUseLif1166(d []interface{}) edpt.OverlayTunnelVtepRemoteIpv6AddressUseLif1166 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressUseLif1166
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(string)
		ret.Lif = in["lif"].(string)
		//omit uuid
	}
	return ret
}

func getSliceOverlayTunnelVtepRemoteIpv6AddressVniList(d []interface{}) []edpt.OverlayTunnelVtepRemoteIpv6AddressVniList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepRemoteIpv6AddressVniList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepRemoteIpv6AddressVniList
		oi.Segment = in["segment"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointOverlayTunnelVtepRemoteIpv6Address(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpv6Address {
	var ret edpt.OverlayTunnelVtepRemoteIpv6Address
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.Encap = d.Get("encap").(string)
	ret.Inst.GreKeepalive = getObjectOverlayTunnelVtepRemoteIpv6AddressGreKeepalive1164(d.Get("gre_keepalive").([]interface{}))
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.UseGreKey = getObjectOverlayTunnelVtepRemoteIpv6AddressUseGreKey1165(d.Get("use_gre_key").([]interface{}))
	ret.Inst.UseLif = getObjectOverlayTunnelVtepRemoteIpv6AddressUseLif1166(d.Get("use_lif").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VniList = getSliceOverlayTunnelVtepRemoteIpv6AddressVniList(d.Get("vni_list").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
