package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwSessionAgingTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_session_aging_tcp`: TCP Aging options\n\n__PLACEHOLDER__",
		CreateContext: resourceFwSessionAgingTcpCreate,
		UpdateContext: resourceFwSessionAgingTcpUpdate,
		ReadContext:   resourceFwSessionAgingTcpRead,
		DeleteContext: resourceFwSessionAgingTcpDelete,

		Schema: map[string]*schema.Schema{
			"force_delete_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number (second))",
			},
			"force_delete_timeout_100ms": {
				Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)",
			},
			"half_close_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP Half Close Idle Timeout (sec), default is off (number)",
			},
			"half_open_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP Half Open Idle Timeout (sec), default is off (number)",
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
			"tcp_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 600, Description: "Idle Timeout (sec), default is 600 (number)",
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
func resourceFwSessionAgingTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwSessionAgingTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwSessionAgingTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwSessionAgingTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwSessionAgingTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwSessionAgingTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwSessionAgingTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwSessionAgingTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwSessionAgingTcpPortCfg(d []interface{}) []edpt.FwSessionAgingTcpPortCfg {

	count1 := len(d)
	ret := make([]edpt.FwSessionAgingTcpPortCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwSessionAgingTcpPortCfg
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

func dataToEndpointFwSessionAgingTcp(d *schema.ResourceData) edpt.FwSessionAgingTcp {
	var ret edpt.FwSessionAgingTcp
	ret.Inst.ForceDeleteTimeout = d.Get("force_delete_timeout").(int)
	ret.Inst.ForceDeleteTimeout100ms = d.Get("force_delete_timeout_100ms").(int)
	ret.Inst.HalfCloseIdleTimeout = d.Get("half_close_idle_timeout").(int)
	ret.Inst.HalfOpenIdleTimeout = d.Get("half_open_idle_timeout").(int)
	ret.Inst.PortCfg = getSliceFwSessionAgingTcpPortCfg(d.Get("port_cfg").([]interface{}))
	ret.Inst.TcpIdleTimeout = d.Get("tcp_idle_timeout").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
