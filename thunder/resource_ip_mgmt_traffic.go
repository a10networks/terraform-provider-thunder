package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpMgmtTraffic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_mgmt_traffic`: Management traffic IP parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceIpMgmtTrafficCreate,
		UpdateContext: resourceIpMgmtTrafficUpdate,
		ReadContext:   resourceIpMgmtTrafficRead,
		DeleteContext: resourceIpMgmtTrafficDelete,

		Schema: map[string]*schema.Schema{
			"source_interface": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"lif": {
							Type: schema.TypeInt, Optional: true, Description: "Logical interface (Lif interface number)",
						},
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
					},
				},
			},
			"traffic_type": {
				Type: schema.TypeString, Required: true, Description: "'all': All; 'ftp': FTP; 'ntp': NTP; 'snmp-trap': SNMP Trap; 'ssh': SSH and SCP; 'syslog': SYSLOG; 'telnet': Telnet; 'tftp': TFTP; 'web': Web - HTTP and HTTPS;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpMgmtTrafficCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMgmtTrafficCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMgmtTraffic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpMgmtTrafficRead(ctx, d, meta)
	}
	return diags
}

func resourceIpMgmtTrafficUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMgmtTrafficUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMgmtTraffic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpMgmtTrafficRead(ctx, d, meta)
	}
	return diags
}
func resourceIpMgmtTrafficDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMgmtTrafficDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMgmtTraffic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpMgmtTrafficRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpMgmtTrafficRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpMgmtTraffic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpMgmtTrafficSourceInterface(d []interface{}) edpt.IpMgmtTrafficSourceInterface {

	count1 := len(d)
	var ret edpt.IpMgmtTrafficSourceInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Loopback = in["loopback"].(int)
		ret.Ethernet = in["ethernet"].(int)
		ret.Ve = in["ve"].(int)
		ret.Trunk = in["trunk"].(int)
		ret.Lif = in["lif"].(int)
		ret.Tunnel = in["tunnel"].(int)
	}
	return ret
}

func dataToEndpointIpMgmtTraffic(d *schema.ResourceData) edpt.IpMgmtTraffic {
	var ret edpt.IpMgmtTraffic
	ret.Inst.SourceInterface = getObjectIpMgmtTrafficSourceInterface(d.Get("source_interface").([]interface{}))
	ret.Inst.TrafficType = d.Get("traffic_type").(string)
	//omit uuid
	return ret
}
