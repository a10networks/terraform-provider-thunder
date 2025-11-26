package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthSourceNat() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_source_nat`: Define Source NAT for health monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthSourceNatCreate,
		UpdateContext: resourceHealthSourceNatUpdate,
		ReadContext:   resourceHealthSourceNatRead,
		DeleteContext: resourceHealthSourceNatDelete,

		Schema: map[string]*schema.Schema{
			"enable_vrrp_a_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward health check by active device only",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
			},
			"interface": {
				Type: schema.TypeString, Optional: true, Description: "'ethernet': ethernet; 'trunk': trunk; 've': ve;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'act_recv_from_sby': Packets received from standby; 'act_send_to_sby': Packets sent to standby; 'sby_recv_from_act': Packets received from active; 'sby_send_to_act': Packets sent to active; 'sby_recv_from_act_err': Packets received from active error; 'recv_from_kernel': Packets received from kernel; 'send_to_kernel': Packets sent to kernel; 'send_to_kernel_err': Packets sent to kernel error; 'sby_no_peer': Peer not found on standby; 'dcmsg_err': DCMSG error; 'no_slb_object': SLB object not found; 'smart_nat_init_port_err': Smart NAT port initialization error; 'smart_nat_init_inst_err': Smart NAT instance initialization error; 'smart_nat_rserver_route_err': Smart NAT rserver route update error; 'smart_nat_rserver_ip_err': Smart NAT rserver ip update error; 'nat_resource_err': NAT resource error; 'frag_err': Fragmentation error;",
						},
					},
				},
			},
			"smart_nat_precedence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use smart nat when resourece is presented in virtual port",
			},
			"smart_nat_vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Smart nat VRID",
			},
			"source_nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Use source nat for all health check (nat pool)",
			},
			"source_nat_pool_v6": {
				Type: schema.TypeString, Optional: true, Description: "Use ipv6 source nat for all health check (nat pool)",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk interface number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface number",
			},
		},
	}
}
func resourceHealthSourceNatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthSourceNatCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthSourceNat(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthSourceNatRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthSourceNatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthSourceNatUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthSourceNat(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthSourceNatRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthSourceNatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthSourceNatDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthSourceNat(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthSourceNatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthSourceNatRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthSourceNat(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceHealthSourceNatSamplingEnable(d []interface{}) []edpt.HealthSourceNatSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.HealthSourceNatSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.HealthSourceNatSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointHealthSourceNat(d *schema.ResourceData) edpt.HealthSourceNat {
	var ret edpt.HealthSourceNat
	ret.Inst.EnableVrrpAMode = d.Get("enable_vrrp_a_mode").(int)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Interface = d.Get("interface").(string)
	ret.Inst.SamplingEnable = getSliceHealthSourceNatSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SmartNatPrecedence = d.Get("smart_nat_precedence").(int)
	ret.Inst.SmartNatVrid = d.Get("smart_nat_vrid").(int)
	ret.Inst.SourceNatPool = d.Get("source_nat_pool").(string)
	ret.Inst.SourceNatPoolV6 = d.Get("source_nat_pool_v6").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	return ret
}
