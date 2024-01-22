package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwSessionAging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_session_aging`: Session aging\n\n__PLACEHOLDER__",
		CreateContext: resourceFwSessionAgingCreate,
		UpdateContext: resourceFwSessionAgingUpdate,
		ReadContext:   resourceFwSessionAgingRead,
		DeleteContext: resourceFwSessionAgingDelete,

		Schema: map[string]*schema.Schema{
			"icmp_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Idle Timeout time (default 2 seconds) (Second, default 2)",
			},
			"ip_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Idle Timeout time(sec), default is 30 (Second)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "session-aging Template (session-aging Template name)",
			},
			"tcp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 600, Description: "Idle Timeout (sec), default is 600 (number)",
						},
						"half_open_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Half Open Idle Timeout (sec), default is off (number)",
						},
						"half_close_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Half Close Idle Timeout (sec), default is off (number)",
						},
						"force_delete_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number (second))",
						},
						"force_delete_timeout_100ms": {
							Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)",
						},
						"port_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_idle_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Idle Timeout (sec), default is 600 (number)",
									},
									"half_open_idle_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Half Open Idle Timeout (sec), default is off (number)",
									},
									"half_close_idle_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Half Close Idle Timeout (sec), default is off (number)",
									},
									"force_delete_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number (second))",
									},
									"force_delete_timeout_100ms": {
										Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"udp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "Idle Timeout (sec), default is 120 (number)",
						},
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
func resourceFwSessionAgingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwSessionAgingRead(ctx, d, meta)
	}
	return diags
}

func resourceFwSessionAgingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwSessionAgingRead(ctx, d, meta)
	}
	return diags
}
func resourceFwSessionAgingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwSessionAgingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFwSessionAgingTcp374(d []interface{}) edpt.FwSessionAgingTcp374 {

	count1 := len(d)
	var ret edpt.FwSessionAgingTcp374
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpIdleTimeout = in["tcp_idle_timeout"].(int)
		ret.HalfOpenIdleTimeout = in["half_open_idle_timeout"].(int)
		ret.HalfCloseIdleTimeout = in["half_close_idle_timeout"].(int)
		ret.ForceDeleteTimeout = in["force_delete_timeout"].(int)
		ret.ForceDeleteTimeout100ms = in["force_delete_timeout_100ms"].(int)
		ret.PortCfg = getSliceFwSessionAgingTcpPortCfg375(in["port_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceFwSessionAgingTcpPortCfg375(d []interface{}) []edpt.FwSessionAgingTcpPortCfg375 {

	count1 := len(d)
	ret := make([]edpt.FwSessionAgingTcpPortCfg375, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwSessionAgingTcpPortCfg375
		oi.TcpPort = in["tcp_port"].(int)
		oi.TcpIdleTimeout = in["tcp_idle_timeout"].(int)
		oi.HalfOpenIdleTimeout = in["half_open_idle_timeout"].(int)
		oi.HalfCloseIdleTimeout = in["half_close_idle_timeout"].(int)
		oi.ForceDeleteTimeout = in["force_delete_timeout"].(int)
		oi.ForceDeleteTimeout100ms = in["force_delete_timeout_100ms"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwSessionAgingUdp376(d []interface{}) edpt.FwSessionAgingUdp376 {

	count1 := len(d)
	var ret edpt.FwSessionAgingUdp376
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpIdleTimeout = in["udp_idle_timeout"].(int)
		ret.PortCfg = getSliceFwSessionAgingUdpPortCfg377(in["port_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceFwSessionAgingUdpPortCfg377(d []interface{}) []edpt.FwSessionAgingUdpPortCfg377 {

	count1 := len(d)
	ret := make([]edpt.FwSessionAgingUdpPortCfg377, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwSessionAgingUdpPortCfg377
		oi.UdpPort = in["udp_port"].(int)
		oi.UdpIdleTimeout = in["udp_idle_timeout"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwSessionAging(d *schema.ResourceData) edpt.FwSessionAging {
	var ret edpt.FwSessionAging
	ret.Inst.IcmpIdleTimeout = d.Get("icmp_idle_timeout").(int)
	ret.Inst.IpIdleTimeout = d.Get("ip_idle_timeout").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Tcp = getObjectFwSessionAgingTcp374(d.Get("tcp").([]interface{}))
	ret.Inst.Udp = getObjectFwSessionAgingUdp376(d.Get("udp").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
