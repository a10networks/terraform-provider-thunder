package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRouteRib() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_route_rib`: Establish static routes\n\n__PLACEHOLDER__",
		CreateContext: resourceIpRouteRibCreate,
		UpdateContext: resourceIpRouteRibUpdate,
		ReadContext:   resourceIpRouteRibRead,
		DeleteContext: resourceIpRouteRibDelete,

		Schema: map[string]*schema.Schema{
			"ip_dest_addr": {
				Type: schema.TypeString, Required: true, Description: "Destination prefix",
			},
			"ip_mask": {
				Type: schema.TypeString, Required: true, Description: "Destination prefix mask",
			},
			"ip_nexthop_ipv4": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_next_hop": {
							Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
						},
						"distance_nexthop_ip": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Distance value for this route",
						},
						"description_nexthop_ip": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
						},
					},
				},
			},
			"ip_nexthop_lif": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lif": {
							Type: schema.TypeString, Optional: true, Description: "LIF Interface (Logical tunnel interface name)",
						},
						"description_nexthop_lif": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
						},
					},
				},
			},
			"ip_nexthop_partition": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition_name": {
							Type: schema.TypeString, Optional: true, Description: "Name of network partition",
						},
						"vrid_num_in_partition": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP-A vrid",
						},
						"description_nexthop_partition": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
						},
						"description_partition_vrid": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
						},
					},
				},
			},
			"ip_nexthop_tunnel": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
						"ip_next_hop_tunnel": {
							Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
						},
						"distance_nexthop_tunnel": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Distance value for this route",
						},
						"description_nexthop_tunnel": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
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
func resourceIpRouteRibCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteRibCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteRib(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteRibRead(ctx, d, meta)
	}
	return diags
}

func resourceIpRouteRibUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteRibUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteRib(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteRibRead(ctx, d, meta)
	}
	return diags
}
func resourceIpRouteRibDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteRibDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteRib(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpRouteRibRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteRibRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteRib(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpRouteRibIpNexthopIpv4(d []interface{}) []edpt.IpRouteRibIpNexthopIpv4 {

	count1 := len(d)
	ret := make([]edpt.IpRouteRibIpNexthopIpv4, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpRouteRibIpNexthopIpv4
		oi.IpNextHop = in["ip_next_hop"].(string)
		oi.DistanceNexthopIp = in["distance_nexthop_ip"].(int)
		oi.DescriptionNexthopIp = in["description_nexthop_ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpRouteRibIpNexthopLif(d []interface{}) []edpt.IpRouteRibIpNexthopLif {

	count1 := len(d)
	ret := make([]edpt.IpRouteRibIpNexthopLif, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpRouteRibIpNexthopLif
		oi.Lif = in["lif"].(string)
		oi.DescriptionNexthopLif = in["description_nexthop_lif"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpRouteRibIpNexthopPartition(d []interface{}) []edpt.IpRouteRibIpNexthopPartition {

	count1 := len(d)
	ret := make([]edpt.IpRouteRibIpNexthopPartition, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpRouteRibIpNexthopPartition
		oi.PartitionName = in["partition_name"].(string)
		oi.VridNumInPartition = in["vrid_num_in_partition"].(int)
		oi.DescriptionNexthopPartition = in["description_nexthop_partition"].(string)
		oi.DescriptionPartitionVrid = in["description_partition_vrid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpRouteRibIpNexthopTunnel(d []interface{}) []edpt.IpRouteRibIpNexthopTunnel {

	count1 := len(d)
	ret := make([]edpt.IpRouteRibIpNexthopTunnel, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpRouteRibIpNexthopTunnel
		oi.Tunnel = in["tunnel"].(int)
		oi.IpNextHopTunnel = in["ip_next_hop_tunnel"].(string)
		oi.DistanceNexthopTunnel = in["distance_nexthop_tunnel"].(int)
		oi.DescriptionNexthopTunnel = in["description_nexthop_tunnel"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpRouteRib(d *schema.ResourceData) edpt.IpRouteRib {
	var ret edpt.IpRouteRib
	ret.Inst.IpDestAddr = d.Get("ip_dest_addr").(string)
	ret.Inst.IpMask = d.Get("ip_mask").(string)
	ret.Inst.IpNexthopIpv4 = getSliceIpRouteRibIpNexthopIpv4(d.Get("ip_nexthop_ipv4").([]interface{}))
	ret.Inst.IpNexthopLif = getSliceIpRouteRibIpNexthopLif(d.Get("ip_nexthop_lif").([]interface{}))
	ret.Inst.IpNexthopPartition = getSliceIpRouteRibIpNexthopPartition(d.Get("ip_nexthop_partition").([]interface{}))
	ret.Inst.IpNexthopTunnel = getSliceIpRouteRibIpNexthopTunnel(d.Get("ip_nexthop_tunnel").([]interface{}))
	//omit uuid
	return ret
}
