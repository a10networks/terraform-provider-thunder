package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ip_address`: Configure remote tunnel end point parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpAddressCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpAddressUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpAddressRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpAddressDelete,

		Schema: map[string]*schema.Schema{
			"encap": {
				Type: schema.TypeString, Optional: true, Description: "'nvgre': Tunnel Encapsulation Type is NVGRE; 'vxlan': Tunnel Encapsulation Type is VXLAN;",
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
			"ip_address": {
				Type: schema.TypeString, Required: true, Description: "IP Address of the remote VTEP",
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
func resourceOverlayTunnelVtepRemoteIpAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectOverlayTunnelVtepRemoteIpAddressGreKeepalive1081(d []interface{}) edpt.OverlayTunnelVtepRemoteIpAddressGreKeepalive1081 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpAddressGreKeepalive1081
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetryTime = in["retry_time"].(int)
		ret.RetryCount = in["retry_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpAddressUseGreKey1082(d []interface{}) edpt.OverlayTunnelVtepRemoteIpAddressUseGreKey1082 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpAddressUseGreKey1082
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreKey = in["gre_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpAddressUseLif1083(d []interface{}) edpt.OverlayTunnelVtepRemoteIpAddressUseLif1083 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpAddressUseLif1083
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(string)
		ret.Lif = in["lif"].(string)
		//omit uuid
	}
	return ret
}

func getSliceOverlayTunnelVtepRemoteIpAddressVniList(d []interface{}) []edpt.OverlayTunnelVtepRemoteIpAddressVniList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepRemoteIpAddressVniList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepRemoteIpAddressVniList
		oi.Segment = in["segment"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointOverlayTunnelVtepRemoteIpAddress(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpAddress {
	var ret edpt.OverlayTunnelVtepRemoteIpAddress
	ret.Inst.Encap = d.Get("encap").(string)
	ret.Inst.GreKeepalive = getObjectOverlayTunnelVtepRemoteIpAddressGreKeepalive1081(d.Get("gre_keepalive").([]interface{}))
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.UseGreKey = getObjectOverlayTunnelVtepRemoteIpAddressUseGreKey1082(d.Get("use_gre_key").([]interface{}))
	ret.Inst.UseLif = getObjectOverlayTunnelVtepRemoteIpAddressUseLif1083(d.Get("use_lif").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VniList = getSliceOverlayTunnelVtepRemoteIpAddressVniList(d.Get("vni_list").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
