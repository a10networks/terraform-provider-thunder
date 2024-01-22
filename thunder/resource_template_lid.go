package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_lid`: Create an Lid\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateLidCreate,
		UpdateContext: resourceTemplateLidUpdate,
		ReadContext:   resourceTemplateLidRead,
		DeleteContext: resourceTemplateLidDelete,

		Schema: map[string]*schema.Schema{
			"ddos_protection_factor": {
				Type: schema.TypeInt, Optional: true, Description: "Enable DDoS Protection (Multiplier of the downlink PPS)",
			},
			"downlink_pps": {
				Type: schema.TypeInt, Optional: true, Description: "Downlink PPS limit (Number of Packets per second)",
			},
			"downlink_throughput": {
				Type: schema.TypeInt, Optional: true, Description: "Downlink Throughput limit (Mega Bits per second)",
			},
			"lid_number": {
				Type: schema.TypeInt, Required: true, Description: "Lid Number",
			},
			"limit_cps": {
				Type: schema.TypeInt, Optional: true, Description: "Enable Connections Per Second Rate Limit (Number of Connections per second)",
			},
			"limit_rate": {
				Type: schema.TypeString, Optional: true, Description: "'limit-pps': Enable Packets Per Second Rate Limit; 'limit-throughput': Enable Throughput Rate Limit;",
			},
			"respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default:off)",
			},
			"src_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"concurrent_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Concurrent Session Limit per Source IP Address (Number of Concurrent Sessions)",
						},
						"log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when Session Limit is exceeded",
						},
						"prefix_length": {
							Type: schema.TypeInt, Optional: true, Description: "Source prefix length",
						},
						"burstsize_downlink_pps": {
							Type: schema.TypeInt, Optional: true, Description: "Downlink PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
						},
						"burstsize_uplink_pps": {
							Type: schema.TypeInt, Optional: true, Description: "Uplink PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
						},
						"burstsize_total_pps": {
							Type: schema.TypeInt, Optional: true, Description: "Total PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
						},
						"burstsize_downlink_throughput": {
							Type: schema.TypeInt, Optional: true, Description: "Downlink Throughput Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
						},
						"burstsize_uplink_throughput": {
							Type: schema.TypeInt, Optional: true, Description: "Uplink Throughput Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
						},
						"burstsize_total_throughput": {
							Type: schema.TypeInt, Optional: true, Description: "Total Throughput Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
						},
						"burstsize_cps": {
							Type: schema.TypeInt, Optional: true, Description: "CPS Token Bucket Size (Must Exceed Configured Rate) (In Connections per second)",
						},
					},
				},
			},
			"total_pps": {
				Type: schema.TypeInt, Optional: true, Description: "Total PPS limit (Number of Packets per second)",
			},
			"total_throughput": {
				Type: schema.TypeInt, Optional: true, Description: "Total Throughput limit (Mega Bits per second)",
			},
			"uplink_pps": {
				Type: schema.TypeInt, Optional: true, Description: "Uplink PPS limit (Number of Packets per second)",
			},
			"uplink_throughput": {
				Type: schema.TypeInt, Optional: true, Description: "Uplink Throughput limit (Mega Bits per second)",
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
func resourceTemplateLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLidRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLidRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTemplateLidSrcIp(d []interface{}) edpt.TemplateLidSrcIp {

	count1 := len(d)
	var ret edpt.TemplateLidSrcIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentSessions = in["concurrent_sessions"].(int)
		ret.Log = in["log"].(int)
		ret.PrefixLength = in["prefix_length"].(int)
		ret.BurstsizeDownlinkPps = in["burstsize_downlink_pps"].(int)
		ret.BurstsizeUplinkPps = in["burstsize_uplink_pps"].(int)
		ret.BurstsizeTotalPps = in["burstsize_total_pps"].(int)
		ret.BurstsizeDownlinkThroughput = in["burstsize_downlink_throughput"].(int)
		ret.BurstsizeUplinkThroughput = in["burstsize_uplink_throughput"].(int)
		ret.BurstsizeTotalThroughput = in["burstsize_total_throughput"].(int)
		ret.BurstsizeCps = in["burstsize_cps"].(int)
	}
	return ret
}

func dataToEndpointTemplateLid(d *schema.ResourceData) edpt.TemplateLid {
	var ret edpt.TemplateLid
	ret.Inst.DdosProtectionFactor = d.Get("ddos_protection_factor").(int)
	ret.Inst.DownlinkPps = d.Get("downlink_pps").(int)
	ret.Inst.DownlinkThroughput = d.Get("downlink_throughput").(int)
	ret.Inst.LidNumber = d.Get("lid_number").(int)
	ret.Inst.LimitCps = d.Get("limit_cps").(int)
	ret.Inst.LimitRate = d.Get("limit_rate").(string)
	ret.Inst.RespondToUserMac = d.Get("respond_to_user_mac").(int)
	ret.Inst.SrcIp = getObjectTemplateLidSrcIp(d.Get("src_ip").([]interface{}))
	ret.Inst.TotalPps = d.Get("total_pps").(int)
	ret.Inst.TotalThroughput = d.Get("total_throughput").(int)
	ret.Inst.UplinkPps = d.Get("uplink_pps").(int)
	ret.Inst.UplinkThroughput = d.Get("uplink_throughput").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
