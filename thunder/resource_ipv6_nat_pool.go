package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatPool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_nat_pool`: IPv6 pool name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6NatPoolCreate,
		UpdateContext: resourceIpv6NatPoolUpdate,
		ReadContext:   resourceIpv6NatPoolRead,
		DeleteContext: resourceIpv6NatPoolDelete,

		Schema: map[string]*schema.Schema{
			"chunk_sharing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share NAT pool chunk across CPUs",
			},
			"end_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure end IP address of NAT pool",
			},
			"gateway": {
				Type: schema.TypeString, Optional: true, Description: "Configure gateway IP",
			},
			"ip_rr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use IP address round-robin behavior",
			},
			"netmask": {
				Type: schema.TypeInt, Optional: true, Description: "Configure mask for pool",
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name",
			},
			"port_overload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Nat Pool Port overload",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'Port-Usage': Port-Usage; 'Total-Used': Total-Used; 'Total-Freed': Total-Freed; 'Failed': Failed;",
						},
					},
				},
			},
			"scaleout_device_id": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Scaleout device id to which this NAT pool is to be bound (Specify Scaleout device id)",
			},
			"start_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure start IP address of NAT pool",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceIpv6NatPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NatPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6NatPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6NatPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpv6NatPoolSamplingEnable(d []interface{}) []edpt.Ipv6NatPoolSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Ipv6NatPoolSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6NatPoolSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6NatPool(d *schema.ResourceData) edpt.Ipv6NatPool {
	var ret edpt.Ipv6NatPool
	ret.Inst.ChunkSharing = d.Get("chunk_sharing").(int)
	ret.Inst.EndAddress = d.Get("end_address").(string)
	ret.Inst.Gateway = d.Get("gateway").(string)
	ret.Inst.IpRr = d.Get("ip_rr").(int)
	ret.Inst.Netmask = d.Get("netmask").(int)
	ret.Inst.PoolName = d.Get("pool_name").(string)
	ret.Inst.PortOverload = d.Get("port_overload").(int)
	ret.Inst.SamplingEnable = getSliceIpv6NatPoolSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ScaleoutDeviceId = d.Get("scaleout_device_id").(int)
	ret.Inst.StartAddress = d.Get("start_address").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
