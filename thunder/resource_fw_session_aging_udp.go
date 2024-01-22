package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwSessionAgingUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_session_aging_udp`: UDP Aging options\n\n__PLACEHOLDER__",
		CreateContext: resourceFwSessionAgingUdpCreate,
		UpdateContext: resourceFwSessionAgingUdpUpdate,
		ReadContext:   resourceFwSessionAgingUdpRead,
		DeleteContext: resourceFwSessionAgingUdpDelete,

		Schema: map[string]*schema.Schema{
			"port_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"udp_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "Idle Timeout (sec), default is 120 (number)",
						},
					},
				},
			},
			"udp_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Idle Timeout (sec), default is 120 (number)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceFwSessionAgingUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwSessionAgingUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwSessionAgingUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwSessionAgingUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwSessionAgingUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwSessionAgingUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwSessionAgingUdpPortCfg(d []interface{}) []edpt.FwSessionAgingUdpPortCfg {

	count1 := len(d)
	ret := make([]edpt.FwSessionAgingUdpPortCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwSessionAgingUdpPortCfg
		oi.UdpPort = in["udp_port"].(int)
		oi.UdpIdleTimeout = in["udp_idle_timeout"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwSessionAgingUdp(d *schema.ResourceData) edpt.FwSessionAgingUdp {
	var ret edpt.FwSessionAgingUdp
	ret.Inst.PortCfg = getSliceFwSessionAgingUdpPortCfg(d.Get("port_cfg").([]interface{}))
	ret.Inst.UdpIdleTimeout = d.Get("udp_idle_timeout").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
