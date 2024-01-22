package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceObjectNetwork() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_object_network`: Configure Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceObjectNetworkCreate,
		UpdateContext: resourceObjectNetworkUpdate,
		ReadContext:   resourceObjectNetworkRead,
		DeleteContext: resourceObjectNetworkDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description of the object instance",
			},
			"ip_range_end": {
				Type: schema.TypeString, Optional: true, Description: "IPV4 Host address end",
			},
			"ip_range_start": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Host Address start",
			},
			"ipv6_range_end": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 Host address end",
			},
			"ipv6_range_start": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Host Address start",
			},
			"ipv6_subnet": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Network Address",
			},
			"net_name": {
				Type: schema.TypeString, Required: true, Description: "Network Object Name",
			},
			"subnet": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Network Address",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceObjectNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectNetworkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectNetwork(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceObjectNetworkRead(ctx, d, meta)
	}
	return diags
}

func resourceObjectNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectNetworkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectNetwork(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceObjectNetworkRead(ctx, d, meta)
	}
	return diags
}
func resourceObjectNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectNetworkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectNetwork(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceObjectNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectNetworkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectNetwork(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointObjectNetwork(d *schema.ResourceData) edpt.ObjectNetwork {
	var ret edpt.ObjectNetwork
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.IpRangeEnd = d.Get("ip_range_end").(string)
	ret.Inst.IpRangeStart = d.Get("ip_range_start").(string)
	ret.Inst.Ipv6RangeEnd = d.Get("ipv6_range_end").(string)
	ret.Inst.Ipv6RangeStart = d.Get("ipv6_range_start").(string)
	ret.Inst.Ipv6Subnet = d.Get("ipv6_subnet").(string)
	ret.Inst.NetName = d.Get("net_name").(string)
	ret.Inst.Subnet = d.Get("subnet").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
