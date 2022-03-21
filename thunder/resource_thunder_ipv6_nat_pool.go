package thunder

import (
	"context"
	"util"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpv6NatPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpv6NatPoolCreate,
		UpdateContext: resourceIpv6NatPoolUpdate,
		ReadContext:   resourceIpv6NatPoolRead,
		DeleteContext: resourceIpv6NatPoolDelete,
		Schema: map[string]*schema.Schema{
			"end_address": {
				Type: schema.TypeString, ForceNew: true, Optional: true, Description: "Configure end IP address of NAT pool",
				ValidateFunc: validation.IsIPv6Address,
			},
			"gateway": {
				Type: schema.TypeString, ForceNew: true, Optional: true, Description: "Configure gateway IP",
				ValidateFunc: validation.IsIPv6Address,
			},
			"ip_rr": {
				Type: schema.TypeInt, ForceNew: true, Optional: true, Default: 0, Description: "Use IP address round-robin behavior",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"netmask": {
				Type: schema.TypeInt, ForceNew: true, Optional: true, Description: "Configure mask for pool",
				ValidateFunc: validation.IntBetween(64, 128),
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Specify pool name",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"port_overload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Nat Pool Port overload",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'Port-Usage': Port-Usage; 'Total-Used': Total-Used; 'Total-Freed': Total-Freed; 'Failed': Failed; ",
						},
					},
				},
			},
			"scaleout_device_id": {
				Type: schema.TypeInt, ForceNew: true, Optional: true, Description: "Configure Scaleout device id to which this NAT pool is to be bound (Specify Scaleout device id)",
				ValidateFunc: validation.IntBetween(1, 16),
			},
			"start_address": {
				Type: schema.TypeString, ForceNew: true, Optional: true, Description: "Configure start IP address of NAT pool",
				ValidateFunc: validation.IsIPv6Address,
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, ForceNew: true, Optional: true, Description: "Specify VRRP-A vrid (Specify ha VRRP-A vrid)",
				ValidateFunc: validation.IntBetween(1, 31),
			},
		},
	}
}

func resourceIpv6NatPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6NatPoolCreate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPool(d)
		d.SetId(obj.GetId())
		err := PostEndpointIpv6NatPool(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NatPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6NatPoolRead()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		_, err := GetEndpointIpv6NatPool(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6NatPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6NatPoolUpdate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPool(d)
		err := PutEndpointIpv6NatPool(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NatPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6NatPoolDelete()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		err := DeleteEndpointIpv6NatPool(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6NatPool(d *schema.ResourceData) EndpointIpv6NatPool {
	var ret EndpointIpv6NatPool
	count := 0
	ret.Inst.EndAddress = d.Get("end_address").(string)
	ret.Inst.Gateway = d.Get("gateway").(string)
	ret.Inst.IpRr = d.Get("ip_rr").(int)
	ret.Inst.Netmask = d.Get("netmask").(int)
	ret.Inst.PoolName = d.Get("pool_name").(string)
	ret.Inst.PortOverload = d.Get("port_overload").(int)
	count = d.Get("sampling_enable.#").(int)
	ret.Inst.SamplingEnable = make([]Ipv6NatPoolSamplingEnable, 0, count)
	for i := 0; i < count; i++ {
		var obj Ipv6NatPoolSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj.Counters1 = d.Get(prefix + "counters1").(string)
		ret.Inst.SamplingEnable = append(ret.Inst.SamplingEnable, obj)
	}
	ret.Inst.ScaleoutDeviceId = d.Get("scaleout_device_id").(int)
	ret.Inst.StartAddress = d.Get("start_address").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
