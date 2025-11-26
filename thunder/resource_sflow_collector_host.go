package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowCollectorHost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_collector_host`: Configure sFlow collector using FQDN\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowCollectorHostCreate,
		UpdateContext: resourceSflowCollectorHostUpdate,
		ReadContext:   resourceSflowCollectorHostRead,
		DeleteContext: resourceSflowCollectorHostDelete,

		Schema: map[string]*schema.Schema{
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Configure FQDN name of sFlow Receiver",
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
func resourceSflowCollectorHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorHostRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowCollectorHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowCollectorHostRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowCollectorHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowCollectorHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowCollectorHostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowCollectorHost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSflowCollectorHostCustomizedSetting1490(d []interface{}) edpt.SflowCollectorHostCustomizedSetting1490 {

	count1 := len(d)
	var ret edpt.SflowCollectorHostCustomizedSetting1490
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

func dataToEndpointSflowCollectorHost(d *schema.ResourceData) edpt.SflowCollectorHost {
	var ret edpt.SflowCollectorHost
	ret.Inst.CustomizedSetting = getObjectSflowCollectorHostCustomizedSetting1490(d.Get("customized_setting").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
