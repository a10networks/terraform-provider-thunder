package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowCollectorIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_collector_ip`: Configure sFlow IPv4 collector\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowCollectorIpCreate,
		UpdateContext: resourceSflowCollectorIpUpdate,
		ReadContext:   resourceSflowCollectorIpRead,
		DeleteContext: resourceSflowCollectorIpDelete,

		Schema: map[string]*schema.Schema{
			"addr": {
				Type: schema.TypeString, Required: true, Description: "Configure sFlow collector IP address",
			},
			"customized_setting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_enable": {
							Type: schema.TypeString, Optional: true, Description: "'export': Export settings;",
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
func resourceSflowCollectorIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorIpRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowCollectorIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorIpRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowCollectorIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowCollectorIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSflowCollectorIpCustomizedSetting1400(d []interface{}) edpt.SflowCollectorIpCustomizedSetting1400 {

	count1 := len(d)
	var ret edpt.SflowCollectorIpCustomizedSetting1400
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

func dataToEndpointSflowCollectorIp(d *schema.ResourceData) edpt.SflowCollectorIp {
	var ret edpt.SflowCollectorIp
	ret.Inst.Addr = d.Get("addr").(string)
	ret.Inst.CustomizedSetting = getObjectSflowCollectorIpCustomizedSetting1400(d.Get("customized_setting").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
