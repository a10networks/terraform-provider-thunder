package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeIps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_range_ips`: zone port ips information\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortRangeIpsCreate,
		UpdateContext: resourceDdosDstZonePortRangeIpsUpdate,
		ReadContext:   resourceDdosDstZonePortRangeIpsRead,
		DeleteContext: resourceDdosDstZonePortRangeIpsDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'ips_matched_total': IPS Matched Total; 'ips_matched_action_pass': IPS Matched Action Pass; 'ips_matched_action_drop': IPS Matched Action Drop; 'ips_matched_action_blacklist': IPS Matched Action Blacklist; 'ips_matched_severity_high': IPS Matched Severity High; 'ips_matched_severity_medium': IPS Matched Severity Medium; 'ips_matched_severity_low': IPS Matched Severity Low; 'src_ips_matched_action_pass': Src IPS Matched Action Pass; 'src_ips_matched_action_drop': Src IPS Matched Action Drop; 'src_ips_matched_action_blacklist': Src IPS Matched Action Blacklist; 'src_ips_matched_severity_high': Src IPS Matched Severity High; 'src_ips_matched_severity_medium': Src IPS Matched Severity Medium; 'src_ips_matched_severity_low': Src IPS Matched Severity Low;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}
func resourceDdosDstZonePortRangeIpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeIpsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortRangeIpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeIpsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortRangeIpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortRangeIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZonePortRangeIpsSamplingEnable(d []interface{}) []edpt.DdosDstZonePortRangeIpsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeIpsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeIpsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeIps(d *schema.ResourceData) edpt.DdosDstZonePortRangeIps {
	var ret edpt.DdosDstZonePortRangeIps
	ret.Inst.SamplingEnable = getSliceDdosDstZonePortRangeIpsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
