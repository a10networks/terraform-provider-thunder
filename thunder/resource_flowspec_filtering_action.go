package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecFilteringAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_filtering_action`: Configure Filtering Actions for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecFilteringActionCreate,
		UpdateContext: resourceFlowspecFilteringActionUpdate,
		ReadContext:   resourceFlowspecFilteringActionRead,
		DeleteContext: resourceFlowspecFilteringActionDelete,

		Schema: map[string]*schema.Schema{
			"copy_ip_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
			},
			"copy_ip_host_nlri": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
			},
			"copy_ipv6_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
			},
			"copy_ipv6_host_nlri": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
			},
			"dscp_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set DSCP value",
			},
			"ecomm_custom_hex": {
				Type: schema.TypeString, Optional: true, Description: "Custom Extended Community in Hex",
			},
			"ip_host": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 host address",
			},
			"ip_host_nlri": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 host address",
			},
			"ip_host_rt": {
				Type: schema.TypeString, Optional: true, Description: "Type 0x8108 - Route Target IPv4",
			},
			"ipv6_host": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 host address",
			},
			"ipv6_host_nlri": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 host address",
			},
			"next_hop_nlri_type": {
				Type: schema.TypeString, Optional: true, Description: "'ip': Type 0x0800 - IPv4 Address; 'ipv6': Type 0x0800 - IPv6 Address;",
			},
			"next_hop_type": {
				Type: schema.TypeString, Optional: true, Description: "'ip': Type 0x0800 - IPv4 Address; 'ipv6': Type 0x0800 - IPv6 Address;",
			},
			"redirect": {
				Type: schema.TypeString, Optional: true, Description: "'next-hop-nlri': Type 0x0800 - IP encoded in MP_REACH_NLRI Next-hop network; 'next-hop': Type 0x0800 - Extended community Next-hop (Per v2 dated Feb 2015); 'vrf-route-target': Type 0x8008 - Redirect to VRF Route Target;",
			},
			"sample_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable traffic sampling and logging",
			},
			"terminal_action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Evaluation stops after this rule if not set",
			},
			"traffic_class": {
				Type: schema.TypeInt, Optional: true, Description: "Set IPv6 Traffic Class value",
			},
			"traffic_marking": {
				Type: schema.TypeString, Optional: true, Description: "'dscp': IPv4 DSCP; 'ipv6-traffic-class': IPv6 Traffic Class;",
			},
			"traffic_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Type 0x8006 - Apply rate (in Bytes per second) for this class of traffic",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value_ip_host": {
				Type: schema.TypeInt, Optional: true, Description: "2-byte decimal value(local-administrator)",
			},
			"vrf_target_ip": {
				Type: schema.TypeString, Optional: true, Description: "'ip': Type 0x8108 - Redirect to route-target IP;",
			},
			"vrf_target_string": {
				Type: schema.TypeString, Optional: true, Description: "Type 0x8008(ASN-2:Index), 0x8208(ASN-4:Index) - Route Target AS",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceFlowspecFilteringActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFilteringActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFilteringAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecFilteringActionRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecFilteringActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFilteringActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFilteringAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecFilteringActionRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecFilteringActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFilteringActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFilteringAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecFilteringActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFilteringActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFilteringAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecFilteringAction(d *schema.ResourceData) edpt.FlowspecFilteringAction {
	var ret edpt.FlowspecFilteringAction
	ret.Inst.CopyIpHost = d.Get("copy_ip_host").(int)
	ret.Inst.CopyIpHostNlri = d.Get("copy_ip_host_nlri").(int)
	ret.Inst.CopyIpv6Host = d.Get("copy_ipv6_host").(int)
	ret.Inst.CopyIpv6HostNlri = d.Get("copy_ipv6_host_nlri").(int)
	ret.Inst.DscpVal = d.Get("dscp_val").(int)
	ret.Inst.EcommCustomHex = d.Get("ecomm_custom_hex").(string)
	ret.Inst.IpHost = d.Get("ip_host").(string)
	ret.Inst.IpHostNlri = d.Get("ip_host_nlri").(string)
	ret.Inst.IpHostRt = d.Get("ip_host_rt").(string)
	ret.Inst.Ipv6Host = d.Get("ipv6_host").(string)
	ret.Inst.Ipv6HostNlri = d.Get("ipv6_host_nlri").(string)
	ret.Inst.NextHopNlriType = d.Get("next_hop_nlri_type").(string)
	ret.Inst.NextHopType = d.Get("next_hop_type").(string)
	ret.Inst.Redirect = d.Get("redirect").(string)
	ret.Inst.SampleLog = d.Get("sample_log").(int)
	ret.Inst.TerminalAction = d.Get("terminal_action").(int)
	ret.Inst.TrafficClass = d.Get("traffic_class").(int)
	ret.Inst.TrafficMarking = d.Get("traffic_marking").(string)
	ret.Inst.TrafficRate = d.Get("traffic_rate").(int)
	//omit uuid
	ret.Inst.ValueIpHost = d.Get("value_ip_host").(int)
	ret.Inst.VrfTargetIp = d.Get("vrf_target_ip").(string)
	ret.Inst.VrfTargetString = d.Get("vrf_target_string").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
