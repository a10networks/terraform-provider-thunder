package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_tcp`: TCP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodTcpCreate,
		UpdateContext: resourceHealthMonitorMethodTcpUpdate,
		ReadContext:   resourceHealthMonitorMethodTcpRead,
		DeleteContext: resourceHealthMonitorMethodTcpDelete,

		Schema: map[string]*schema.Schema{
			"maintenance": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify response text for maintenance",
			},
			"maintenance_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify text for maintenance",
			},
			"method_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP type",
			},
			"port_halfopen": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set TCP SYN check",
			},
			"port_resp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_contains": {
							Type: schema.TypeString, Optional: true, Description: "Mark server up if response string contains string (Specify the string)",
						},
					},
				},
			},
			"port_send": {
				Type: schema.TypeString, Optional: true, Description: "Send a string to server (Specify the string)",
			},
			"tcp_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify TCP port (Specify port number)",
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
func resourceHealthMonitorMethodTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethodTcpPortResp(d []interface{}) edpt.HealthMonitorMethodTcpPortResp {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodTcpPortResp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortContains = in["port_contains"].(string)
	}
	return ret
}

func dataToEndpointHealthMonitorMethodTcp(d *schema.ResourceData) edpt.HealthMonitorMethodTcp {
	var ret edpt.HealthMonitorMethodTcp
	ret.Inst.Maintenance = d.Get("maintenance").(int)
	ret.Inst.MaintenanceText = d.Get("maintenance_text").(string)
	ret.Inst.MethodTcp = d.Get("method_tcp").(int)
	ret.Inst.PortHalfopen = d.Get("port_halfopen").(int)
	ret.Inst.PortResp = getObjectHealthMonitorMethodTcpPortResp(d.Get("port_resp").([]interface{}))
	ret.Inst.PortSend = d.Get("port_send").(string)
	ret.Inst.TcpPort = d.Get("tcp_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
