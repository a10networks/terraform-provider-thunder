package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowCollectorIpCustomizedSetting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_collector_ip_customized_setting`: Customize settings for collector\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowCollectorIpCustomizedSettingCreate,
		UpdateContext: resourceSflowCollectorIpCustomizedSettingUpdate,
		ReadContext:   resourceSflowCollectorIpCustomizedSettingRead,
		DeleteContext: resourceSflowCollectorIpCustomizedSettingDelete,

		Schema: map[string]*schema.Schema{
			"a10_proprietary_polling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable counters for ACOS control blocks",
			},
			"counter_polling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable counter polling",
			},
			"event_notification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable event notification",
			},
			"export_enable": {
				Type: schema.TypeString, Required: true, Description: "'export': Export settings;",
			},
			"packet_sampling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet sampling",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"port": {
				Type: schema.TypeString, Required: true, Description: "Port",
			},
			"addr": {
				Type: schema.TypeString, Required: true, Description: "Addr",
			},
		},
	}
}
func resourceSflowCollectorIpCustomizedSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpCustomizedSettingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpCustomizedSetting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorIpCustomizedSettingRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowCollectorIpCustomizedSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpCustomizedSettingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpCustomizedSetting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorIpCustomizedSettingRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowCollectorIpCustomizedSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpCustomizedSettingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpCustomizedSetting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowCollectorIpCustomizedSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorIpCustomizedSettingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorIpCustomizedSetting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowCollectorIpCustomizedSetting(d *schema.ResourceData) edpt.SflowCollectorIpCustomizedSetting {
	var ret edpt.SflowCollectorIpCustomizedSetting
	ret.Inst.A10ProprietaryPolling = d.Get("a10_proprietary_polling").(int)
	ret.Inst.CounterPolling = d.Get("counter_polling").(int)
	ret.Inst.EventNotification = d.Get("event_notification").(int)
	ret.Inst.ExportEnable = d.Get("export_enable").(string)
	ret.Inst.PacketSampling = d.Get("packet_sampling").(int)
	//omit uuid
	ret.Inst.Port = d.Get("port").(string)
	ret.Inst.Addr = d.Get("addr").(string)
	return ret
}
