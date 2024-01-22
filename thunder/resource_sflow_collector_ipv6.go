package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowCollectorIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_collector_ipv6`: Configure sFlow IPv6 collector\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowCollectorIpv6Create,
		UpdateContext: resourceSflowCollectorIpv6Update,
		ReadContext:   resourceSflowCollectorIpv6Read,
		DeleteContext: resourceSflowCollectorIpv6Delete,

		Schema: map[string]*schema.Schema{
			"addr": {
				Type: schema.TypeString, Required: true, Description: "Configure sFlow collector IPv6 address",
			},
			"customized_setting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_enable": {
							Type: schema.TypeString, Optional: true, Description: "'export': Customizes export settings for collector;",
						},
						"packet_sampling": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet sampling",
						},
						"counter_polling": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable counter polling",
						},
						"a10_proprietary_polling": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable counters for ACOS control blocks",
						},
						"event_notification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable event notification",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port number (default is 6343)",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "sFlow collector is through out-of-band management",
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
func resourceSflowCollectorIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceSflowCollectorIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceSflowCollectorIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowCollectorIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSflowCollectorIpv6CustomizedSetting1401(d []interface{}) edpt.SflowCollectorIpv6CustomizedSetting1401 {

	count1 := len(d)
	var ret edpt.SflowCollectorIpv6CustomizedSetting1401
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExportEnable = in["export_enable"].(string)
		ret.PacketSampling = in["packet_sampling"].(int)
		ret.CounterPolling = in["counter_polling"].(int)
		ret.A10ProprietaryPolling = in["a10_proprietary_polling"].(int)
		ret.EventNotification = in["event_notification"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSflowCollectorIpv6(d *schema.ResourceData) edpt.SflowCollectorIpv6 {
	var ret edpt.SflowCollectorIpv6
	ret.Inst.Addr = d.Get("addr").(string)
	ret.Inst.CustomizedSetting = getObjectSflowCollectorIpv6CustomizedSetting1401(d.Get("customized_setting").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
