package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatInsideIpv4address() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_fixed_nat_inside_ipv4address`: Configure Fixed NAT\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6FixedNatInsideIpv4addressCreate,
		UpdateContext: resourceCgnv6FixedNatInsideIpv4addressUpdate,
		ReadContext:   resourceCgnv6FixedNatInsideIpv4addressRead,
		DeleteContext: resourceCgnv6FixedNatInsideIpv4addressDelete,

		Schema: map[string]*schema.Schema{
			"dest_rule_list": {
				Type: schema.TypeString, Optional: true, Description: "Bind destination based Rule-List (Fixed NAT Rule-List Name)",
			},
			"dynamic_pool_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure size of Dynamic pool (Default: 0)",
			},
			"inside_end_address": {
				Type: schema.TypeString, Required: true, Description: "IPv4 Inside User End Address",
			},
			"inside_netmask": {
				Type: schema.TypeString, Required: true, Description: "IPv4 Netmask",
			},
			"inside_start_address": {
				Type: schema.TypeString, Required: true, Description: "IPv4 Inside User Start Address",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Default: "use-least-nat-ips", Description: "'use-all-nat-ips': Use all the NAT IP addresses configured; 'use-least-nat-ips': Use the least number of NAT IP addresses required (default);",
			},
			"nat_end_address": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 End NAT Address",
			},
			"nat_ip_list": {
				Type: schema.TypeString, Optional: true, Description: "Name of IP List used to specify NAT addresses",
			},
			"nat_netmask": {
				Type: schema.TypeString, Optional: true, Description: "NAT Addresses IP Netmask",
			},
			"nat_start_address": {
				Type: schema.TypeString, Optional: true, Description: "Start NAT Address",
			},
			"offset": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"random": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Randomly choose the first NAT IP address",
						},
						"numeric_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure a numeric offset to the first NAT IP address",
						},
					},
				},
			},
			"partition": {
				Type: schema.TypeString, Required: true, Description: "Inside User Partition (Partition Name)",
			},
			"ports_per_user": {
				Type: schema.TypeInt, Optional: true, Description: "Configure Ports per Inside User (ports-per-user)",
			},
			"respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (Default: off)",
			},
			"session_quota": {
				Type: schema.TypeInt, Optional: true, Description: "Configure per user quota on sessions",
			},
			"skip_ports_on_rollover": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Not using the first block of ports for NAT IPs smaller than the configured offset",
			},
			"usable_nat_ports": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"usable_start_port": {
							Type: schema.TypeInt, Optional: true, Description: "Start Port of Usable NAT Ports",
						},
						"usable_end_port": {
							Type: schema.TypeInt, Optional: true, Description: "End Port of Usable NAT Ports",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6FixedNatInsideIpv4addressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatInsideIpv4addressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatInsideIpv4address(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatInsideIpv4addressRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6FixedNatInsideIpv4addressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatInsideIpv4addressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatInsideIpv4address(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6FixedNatInsideIpv4addressRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6FixedNatInsideIpv4addressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatInsideIpv4addressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatInsideIpv4address(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6FixedNatInsideIpv4addressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatInsideIpv4addressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatInsideIpv4address(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6FixedNatInsideIpv4addressOffset(d []interface{}) edpt.Cgnv6FixedNatInsideIpv4addressOffset {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatInsideIpv4addressOffset
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Random = in["random"].(int)
		ret.NumericOffset = in["numeric_offset"].(int)
	}
	return ret
}

func getObjectCgnv6FixedNatInsideIpv4addressUsableNatPorts(d []interface{}) edpt.Cgnv6FixedNatInsideIpv4addressUsableNatPorts {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatInsideIpv4addressUsableNatPorts
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UsableStartPort = in["usable_start_port"].(int)
		ret.UsableEndPort = in["usable_end_port"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatInsideIpv4address(d *schema.ResourceData) edpt.Cgnv6FixedNatInsideIpv4address {
	var ret edpt.Cgnv6FixedNatInsideIpv4address
	ret.Inst.DestRuleList = d.Get("dest_rule_list").(string)
	ret.Inst.DynamicPoolSize = d.Get("dynamic_pool_size").(int)
	ret.Inst.InsideEndAddress = d.Get("inside_end_address").(string)
	ret.Inst.InsideNetmask = d.Get("inside_netmask").(string)
	ret.Inst.InsideStartAddress = d.Get("inside_start_address").(string)
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.NatEndAddress = d.Get("nat_end_address").(string)
	ret.Inst.NatIpList = d.Get("nat_ip_list").(string)
	ret.Inst.NatNetmask = d.Get("nat_netmask").(string)
	ret.Inst.NatStartAddress = d.Get("nat_start_address").(string)
	ret.Inst.Offset = getObjectCgnv6FixedNatInsideIpv4addressOffset(d.Get("offset").([]interface{}))
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.PortsPerUser = d.Get("ports_per_user").(int)
	ret.Inst.RespondToUserMac = d.Get("respond_to_user_mac").(int)
	ret.Inst.SessionQuota = d.Get("session_quota").(int)
	ret.Inst.SkipPortsOnRollover = d.Get("skip_ports_on_rollover").(int)
	ret.Inst.UsableNatPorts = getObjectCgnv6FixedNatInsideIpv4addressUsableNatPorts(d.Get("usable_nat_ports").([]interface{}))
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
