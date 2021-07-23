package thunder

//Thunder resource Vrrp peer group

import (
	"context"
	"fmt"
	"util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpPeerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVrrpPeerGroupCreate,
		UpdateContext: resourceVrrpPeerGroupUpdate,
		ReadContext:   resourceVrrpPeerGroupRead,
		DeleteContext: resourceVrrpPeerGroupDelete,

		Schema: map[string]*schema.Schema{
			"peer": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_peer_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_peer_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_peer_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_peer_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceVrrpPeerGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating vrrp PeerGroup (Inside resourceVrrpPeerGroupCreate)")

	if client.Host != "" {
		vc := dataToVrrpPeerGroup(d)
		d.SetId("1")
		err := go_thunder.PostVrrpPeerGroup(client.Token, vc, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpPeerGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpPeerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading vrrp PeerGroup (Inside resourceVrrpPeerGroupRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetVrrpPeerGroup(client.Token, client.Host)
if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No vrrp PeerGroup found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceVrrpPeerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceVrrpPeerGroupRead(ctx, d, meta)
}

func resourceVrrpPeerGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceVrrpPeerGroupRead(ctx, d, meta)
}

//Utility method to instantiate Vrrp PeerGroup Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpPeerGroup(d *schema.ResourceData) go_thunder.PeerGroup {
	var vc go_thunder.PeerGroup

	var c go_thunder.PeerGroupInstance

	var obj1 go_thunder.Peer
	prefix := "peer.0."

	IpPeerAddressCfgCount := d.Get(prefix + "ip_peer_address_cfg.#").(int)
	obj1.IPPeerAddress = make([]go_thunder.IPPeerAddressCfg, 0, IpPeerAddressCfgCount)

	for i := 0; i < IpPeerAddressCfgCount; i++ {
		var obj1_1 go_thunder.IPPeerAddressCfg
		prefix1 := fmt.Sprintf(prefix+"ip_peer_address_cfg.%d.", i)
		obj1_1.IPPeerAddress = d.Get(prefix1 + "ip_peer_address").(string)
		obj1.IPPeerAddress = append(obj1.IPPeerAddress, obj1_1)
	}

	Ipv6PeerAddressCfgCount := d.Get(prefix + "ipv6_peer_address_cfg.#").(int)
	obj1.Ipv6PeerAddress = make([]go_thunder.Ipv6PeerAddressCfg, 0, Ipv6PeerAddressCfgCount)

	for i := 0; i < Ipv6PeerAddressCfgCount; i++ {
		var obj1_2 go_thunder.Ipv6PeerAddressCfg
		prefix2 := fmt.Sprintf(prefix+"ipv6_peer_address_cfg.%d.", i)
		obj1_2.Ipv6PeerAddress = d.Get(prefix2 + "ipv6_peer_address").(string)
		obj1.Ipv6PeerAddress = append(obj1.Ipv6PeerAddress, obj1_2)
	}

	c.IPPeerAddressCfg = obj1

	vc.UUID = c

	return vc
}
