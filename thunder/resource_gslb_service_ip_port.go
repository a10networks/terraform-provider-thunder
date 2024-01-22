package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_service_ip_port`: Server Port\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbServiceIpPortCreate,
		UpdateContext: resourceGslbServiceIpPortUpdate,
		ReadContext:   resourceGslbServiceIpPortRead,
		DeleteContext: resourceGslbServiceIpPortDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this GSLB server port; 'disable': Disable this GSLB server port;",
			},
			"follow_port_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Monitor Name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Health Check Monitor",
			},
			"health_check_follow_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify which port to follow for health status (Port Number)",
			},
			"health_check_protocol_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable GSLB Protocol Health Monitor",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"port_proto": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'active': Active Servers; 'current': Current Connections;",
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
			"node_name": {
				Type: schema.TypeString, Required: true, Description: "NodeName",
			},
		},
	}
}
func resourceGslbServiceIpPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceIpPortRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbServiceIpPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceIpPortRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbServiceIpPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbServiceIpPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbServiceIpPortSamplingEnable(d []interface{}) []edpt.GslbServiceIpPortSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbServiceIpPortSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceIpPortSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbServiceIpPort(d *schema.ResourceData) edpt.GslbServiceIpPort {
	var ret edpt.GslbServiceIpPort
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.FollowPortProtocol = d.Get("follow_port_protocol").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckFollowPort = d.Get("health_check_follow_port").(int)
	ret.Inst.HealthCheckProtocolDisable = d.Get("health_check_protocol_disable").(int)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.PortProto = d.Get("port_proto").(string)
	ret.Inst.SamplingEnable = getSliceGslbServiceIpPortSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.NodeName = d.Get("node_name").(string)
	return ret
}
