package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAPeerGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_peer_group`: VRRP-A peer group\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAPeerGroupCreate,
		UpdateContext: resourceVrrpAPeerGroupUpdate,
		ReadContext:   resourceVrrpAPeerGroupRead,
		DeleteContext: resourceVrrpAPeerGroupDelete,

		Schema: map[string]*schema.Schema{
			"peer": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_peer_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_peer_address": {
										Type: schema.TypeString, Optional: true, Description: "IP Address",
									},
								},
							},
						},
						"ipv6_peer_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_peer_address": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpAPeerGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPeerGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPeerGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAPeerGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAPeerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPeerGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPeerGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAPeerGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAPeerGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPeerGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPeerGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAPeerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPeerGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPeerGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAPeerGroupPeer(d []interface{}) edpt.VrrpAPeerGroupPeer {

	count1 := len(d)
	var ret edpt.VrrpAPeerGroupPeer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpPeerAddressCfg = getSliceVrrpAPeerGroupPeerIpPeerAddressCfg(in["ip_peer_address_cfg"].([]interface{}))
		ret.Ipv6PeerAddressCfg = getSliceVrrpAPeerGroupPeerIpv6PeerAddressCfg(in["ipv6_peer_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAPeerGroupPeerIpPeerAddressCfg(d []interface{}) []edpt.VrrpAPeerGroupPeerIpPeerAddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAPeerGroupPeerIpPeerAddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAPeerGroupPeerIpPeerAddressCfg
		oi.IpPeerAddress = in["ip_peer_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAPeerGroupPeerIpv6PeerAddressCfg(d []interface{}) []edpt.VrrpAPeerGroupPeerIpv6PeerAddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAPeerGroupPeerIpv6PeerAddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAPeerGroupPeerIpv6PeerAddressCfg
		oi.Ipv6PeerAddress = in["ipv6_peer_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAPeerGroup(d *schema.ResourceData) edpt.VrrpAPeerGroup {
	var ret edpt.VrrpAPeerGroup
	ret.Inst.Peer = getObjectVrrpAPeerGroupPeer(d.Get("peer").([]interface{}))
	//omit uuid
	return ret
}
