package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6Global() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_global`: Configure LW-4over6 parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6GlobalCreate,
		UpdateContext: resourceCgnv6Lw4o6GlobalUpdate,
		ReadContext:   resourceCgnv6Lw4o6GlobalRead,
		DeleteContext: resourceCgnv6Lw4o6GlobalDelete,

		Schema: map[string]*schema.Schema{
			"hairpinning": {
				Type: schema.TypeString, Optional: true, Default: "filter-none", Description: "'filter-all': Disable all Hairpinning; 'filter-none': Allow all Hairpinning (default); 'filter-self-ip': Block Hairpinning to same IP; 'filter-self-ip-port': Block hairpinning to same IP and Port combination;",
			},
			"icmp_inbound": {
				Type: schema.TypeString, Optional: true, Default: "handle", Description: "'drop': Drop Inbound ICMP packets; 'handle': Handle Inbound ICMP packets(default);",
			},
			"inside_src_access_list": {
				Type: schema.TypeInt, Optional: true, Description: "Access List for inside IPv4 addresses (ACL ID)",
			},
			"nat_prefix_list": {
				Type: schema.TypeString, Optional: true, Description: "Configure LW-4over6 NAT Prefix List (LW-4over6 NAT Prefix Class-list)",
			},
			"no_forward_match": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_icmpv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send ICMPv6 Type 1 Code 5",
						},
					},
				},
			},
			"no_reverse_match": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send ICMP Type 3 Code 1",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'entry_count': Total Entries Configured; 'self_hairpinning_drop': Self-Hairpinning Drops; 'all_hairpinning_drop': All Hairpinning Drops; 'no_match_icmpv6_sent': No-Forward-Match ICMPv6 Sent; 'no_match_icmp_sent': No-Reverse-Match ICMP Sent; 'icmp_inbound_drop': Inbound ICMP Drops; 'fwd_lookup_failed': Forward Route Lookup Failed; 'rev_lookup_failed': Reverse Route Lookup Failed; 'interface_not_configured': LW-4over6 Interfaces not Configured Drops; 'no_binding_table_matches_fwd': No Forward Binding Table Entry Match Drops; 'no_binding_table_matches_rev': No Reverse Binding Table Entry Match Drops; 'session_count': LW-4over6 Session Count; 'system_address_drop': LW-4over6 System Address Drops;",
						},
					},
				},
			},
			"use_binding_table": {
				Type: schema.TypeString, Optional: true, Description: "Bind LW-4over6 binding table for use (LW-4over6 Binding Table Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Lw4o6GlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6GlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6Global(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6GlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6GlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6GlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6Global(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6GlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6GlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6GlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6Global(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6GlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6GlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6Global(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6Lw4o6GlobalNoForwardMatch(d []interface{}) edpt.Cgnv6Lw4o6GlobalNoForwardMatch {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6GlobalNoForwardMatch
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendIcmpv6 = in["send_icmpv6"].(int)
	}
	return ret
}

func getObjectCgnv6Lw4o6GlobalNoReverseMatch(d []interface{}) edpt.Cgnv6Lw4o6GlobalNoReverseMatch {

	count1 := len(d)
	var ret edpt.Cgnv6Lw4o6GlobalNoReverseMatch
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendIcmp = in["send_icmp"].(int)
	}
	return ret
}

func getSliceCgnv6Lw4o6GlobalSamplingEnable(d []interface{}) []edpt.Cgnv6Lw4o6GlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6GlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6GlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6Global(d *schema.ResourceData) edpt.Cgnv6Lw4o6Global {
	var ret edpt.Cgnv6Lw4o6Global
	ret.Inst.Hairpinning = d.Get("hairpinning").(string)
	ret.Inst.IcmpInbound = d.Get("icmp_inbound").(string)
	ret.Inst.InsideSrcAccessList = d.Get("inside_src_access_list").(int)
	ret.Inst.NatPrefixList = d.Get("nat_prefix_list").(string)
	ret.Inst.NoForwardMatch = getObjectCgnv6Lw4o6GlobalNoForwardMatch(d.Get("no_forward_match").([]interface{}))
	ret.Inst.NoReverseMatch = getObjectCgnv6Lw4o6GlobalNoReverseMatch(d.Get("no_reverse_match").([]interface{}))
	ret.Inst.SamplingEnable = getSliceCgnv6Lw4o6GlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UseBindingTable = d.Get("use_binding_table").(string)
	//omit uuid
	return ret
}
