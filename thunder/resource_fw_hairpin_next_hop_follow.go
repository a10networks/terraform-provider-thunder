package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwHairpinNextHopFollow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_hairpin_next_hop_follow`: Configure IP/IPv6 address to lookup for next hop of firewall hairpin traffic\n\n__PLACEHOLDER__",
		CreateContext: resourceFwHairpinNextHopFollowCreate,
		UpdateContext: resourceFwHairpinNextHopFollowUpdate,
		ReadContext:   resourceFwHairpinNextHopFollowRead,
		DeleteContext: resourceFwHairpinNextHopFollowDelete,

		Schema: map[string]*schema.Schema{
			"disable_local_hairpin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable local hairpin",
			},
			"ipv4": {
				Type: schema.TypeString, Optional: true, Description: "IPV4 network",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 network",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwHairpinNextHopFollowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHairpinNextHopFollowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHairpinNextHopFollow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwHairpinNextHopFollowRead(ctx, d, meta)
	}
	return diags
}

func resourceFwHairpinNextHopFollowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHairpinNextHopFollowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHairpinNextHopFollow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwHairpinNextHopFollowRead(ctx, d, meta)
	}
	return diags
}
func resourceFwHairpinNextHopFollowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHairpinNextHopFollowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHairpinNextHopFollow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwHairpinNextHopFollowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHairpinNextHopFollowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHairpinNextHopFollow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwHairpinNextHopFollow(d *schema.ResourceData) edpt.FwHairpinNextHopFollow {
	var ret edpt.FwHairpinNextHopFollow
	ret.Inst.DisableLocalHairpin = d.Get("disable_local_hairpin").(int)
	ret.Inst.Ipv4 = d.Get("ipv4").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	return ret
}
