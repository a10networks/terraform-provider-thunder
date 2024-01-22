package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceIps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_ips`: zone port ips information\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceIpsCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceIpsUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceIpsRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceIpsDelete,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceIpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceIpsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceIpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceIpsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceIpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZonePortZoneServiceIpsSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceIpsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceIpsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceIpsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceIps(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceIps {
	var ret edpt.DdosDstZonePortZoneServiceIps
	ret.Inst.SamplingEnable = getSliceDdosDstZonePortZoneServiceIpsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
