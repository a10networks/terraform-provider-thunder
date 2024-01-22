package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatPool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_pool`: Configure CGNv6 NAT pool\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatPoolCreate,
		UpdateContext: resourceCgnv6NatPoolUpdate,
		ReadContext:   resourceCgnv6NatPoolRead,
		DeleteContext: resourceCgnv6NatPoolDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share with all partitions",
			},
			"end_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure end IP address of NAT pool",
			},
			"exclude_ip": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_ip_start": {
							Type: schema.TypeString, Optional: true, Description: "Single IP address or IP address range start",
						},
						"exclude_ip_end": {
							Type: schema.TypeString, Optional: true, Description: "Address range end",
						},
					},
				},
			},
			"group": {
				Type: schema.TypeString, Optional: true, Description: "Share with a partition group (Partition Group Name)",
			},
			"max_users_per_ip": {
				Type: schema.TypeInt, Optional: true, Description: "Number of users that can be assigned to a NAT IP",
			},
			"netmask": {
				Type: schema.TypeString, Optional: true, Description: "Configure mask for pool",
			},
			"partition": {
				Type: schema.TypeString, Optional: true, Description: "Share with a single partition (Partition Name)",
			},
			"per_batch_port_usage_warning_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Configure warning log threshold for per batch port usage (default: disabled) (Number of ports)",
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name",
			},
			"port_batch_v2_size": {
				Type: schema.TypeString, Optional: true, Description: "'64': Allocate 64 ports at a time; '128': Allocate 128 ports at a time; '256': Allocate 256 ports at a time; '512': Allocate 512 ports at a time; '1024': Allocate 1024 ports at a time; '2048': Allocate 2048 ports at a time; '4096': Allocate 4096 ports at a time;",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share this pool with other partitions (default: not shared)",
			},
			"simultaneous_batch_allocation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allocate same TCP and UDP batches at once",
			},
			"start_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure start IP address of NAT pool",
			},
			"tcp_time_wait_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Minutes before TCP NAT ports can be reused",
			},
			"usable_nat_ports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure usable NAT ports",
			},
			"usable_nat_ports_end": {
				Type: schema.TypeInt, Optional: true, Description: "End Port of Usable NAT Ports",
			},
			"usable_nat_ports_start": {
				Type: schema.TypeInt, Optional: true, Description: "Start Port of Usable NAT Ports (needs to be even)",
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
func resourceCgnv6NatPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatPoolRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatPoolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatPool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6NatPoolExcludeIp(d []interface{}) []edpt.Cgnv6NatPoolExcludeIp {

	count1 := len(d)
	ret := make([]edpt.Cgnv6NatPoolExcludeIp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6NatPoolExcludeIp
		oi.ExcludeIpStart = in["exclude_ip_start"].(string)
		oi.ExcludeIpEnd = in["exclude_ip_end"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6NatPool(d *schema.ResourceData) edpt.Cgnv6NatPool {
	var ret edpt.Cgnv6NatPool
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.EndAddress = d.Get("end_address").(string)
	ret.Inst.ExcludeIp = getSliceCgnv6NatPoolExcludeIp(d.Get("exclude_ip").([]interface{}))
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.MaxUsersPerIp = d.Get("max_users_per_ip").(int)
	ret.Inst.Netmask = d.Get("netmask").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.PerBatchPortUsageWarningThreshold = d.Get("per_batch_port_usage_warning_threshold").(int)
	ret.Inst.PoolName = d.Get("pool_name").(string)
	ret.Inst.PortBatchV2Size = d.Get("port_batch_v2_size").(string)
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.SimultaneousBatchAllocation = d.Get("simultaneous_batch_allocation").(int)
	ret.Inst.StartAddress = d.Get("start_address").(string)
	ret.Inst.TcpTimeWaitInterval = d.Get("tcp_time_wait_interval").(int)
	ret.Inst.UsableNatPorts = d.Get("usable_nat_ports").(int)
	ret.Inst.UsableNatPortsEnd = d.Get("usable_nat_ports_end").(int)
	ret.Inst.UsableNatPortsStart = d.Get("usable_nat_ports_start").(int)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
