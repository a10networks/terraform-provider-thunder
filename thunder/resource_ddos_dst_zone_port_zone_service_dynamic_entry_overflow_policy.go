package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_dynamic_entry_overflow_policy`: Configure overflow policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
			},
			"dummy_name": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quic": {
							Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
						},
						"dns": {
							Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
						},
						"http": {
							Type: schema.TypeString, Optional: true, Description: "DDOS http template",
						},
						"ssl_l4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
						},
						"sip": {
							Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
						},
						"tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
						},
						"udp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
						},
						"encap": {
							Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
						},
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
						},
					},
				},
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
func resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceDynamicEntryOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceDynamicEntryOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceDynamicEntryOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceDynamicEntryOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceDynamicEntryOverflowPolicy(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicy {
	var ret edpt.DdosDstZonePortZoneServiceDynamicEntryOverflowPolicy
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZonePortZoneServiceDynamicEntryOverflowPolicyZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
