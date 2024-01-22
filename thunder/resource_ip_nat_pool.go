package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_pool`: Configure IP pool name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatPoolCreate,
		UpdateContext: resourceIpNatPoolUpdate,
		ReadContext:   resourceIpNatPoolRead,
		DeleteContext: resourceIpNatPoolDelete,

		Schema: map[string]*schema.Schema{
			"chunk_sharing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share NAT pool chunk across CPUs",
			},
			"end_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure end IP address of NAT pool",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"gateway": {
				Type: schema.TypeString, Optional: true, Description: "Configure gateway IP",
			},
			"ip_rr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use IP address round-robin behavior",
			},
			"netmask": {
				Type: schema.TypeString, Optional: true, Description: "Configure mask for pool",
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name or pool group",
			},
			"port_overload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Nat Pool Port overload",
			},
			"scaleout_device_id": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Scaleout device id to which this NAT pool is to be bound (Specify Scaleout device id)",
			},
			"start_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure start IP address of NAT pool",
			},
			"use_if_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Interface IP",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Configure VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceIpNatPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatPoolRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatPool(d *schema.ResourceData) edpt.IpNatPool {
	var ret edpt.IpNatPool
	ret.Inst.ChunkSharing = d.Get("chunk_sharing").(int)
	ret.Inst.EndAddress = d.Get("end_address").(string)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Gateway = d.Get("gateway").(string)
	ret.Inst.IpRr = d.Get("ip_rr").(int)
	ret.Inst.Netmask = d.Get("netmask").(string)
	ret.Inst.PoolName = d.Get("pool_name").(string)
	ret.Inst.PortOverload = d.Get("port_overload").(int)
	ret.Inst.ScaleoutDeviceId = d.Get("scaleout_device_id").(int)
	ret.Inst.StartAddress = d.Get("start_address").(string)
	ret.Inst.UseIfIp = d.Get("use_if_ip").(int)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
