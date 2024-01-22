package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTapMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_tap_monitor`: Configure tap monitor port\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTapMonitorCreate,
		UpdateContext: resourceFwTapMonitorUpdate,
		ReadContext:   resourceFwTapMonitorRead,
		DeleteContext: resourceFwTapMonitorDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable tap monitor mode; 'disable': Disable tap monitor mode (Default:Disable);",
			},
			"tap_port_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tap_eth": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
						},
						"tap_vlan": {
							Type: schema.TypeInt, Optional: true, Description: "Vlan number",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwTapMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTapMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTapMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTapMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTapMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTapMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTapMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTapMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTapMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTapMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTapMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTapMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTapMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTapMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTapMonitorTapPortCfg(d []interface{}) []edpt.FwTapMonitorTapPortCfg {

	count1 := len(d)
	ret := make([]edpt.FwTapMonitorTapPortCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTapMonitorTapPortCfg
		oi.TapEth = in["tap_eth"].(int)
		oi.TapVlan = in["tap_vlan"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTapMonitor(d *schema.ResourceData) edpt.FwTapMonitor {
	var ret edpt.FwTapMonitor
	ret.Inst.Status = d.Get("status").(string)
	ret.Inst.TapPortCfg = getSliceFwTapMonitorTapPortCfg(d.Get("tap_port_cfg").([]interface{}))
	//omit uuid
	return ret
}
