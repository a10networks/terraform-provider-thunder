package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_destination`: Configure destination where netflow records will be sent\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorDestinationCreate,
		UpdateContext: resourceNetflowMonitorDestinationUpdate,
		ReadContext:   resourceNetflowMonitorDestinationRead,
		DeleteContext: resourceNetflowMonitorDestinationDelete,

		Schema: map[string]*schema.Schema{
			"ip_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "IP address of netflow collector",
						},
						"port4": {
							Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Port number, default is 9996",
						},
					},
				},
			},
			"ipv6_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 address of netflow collector",
						},
						"port6": {
							Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Port number, default is 9996",
						},
					},
				},
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Service-group for load balancing between multiple collector servers",
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
func resourceNetflowMonitorDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetflowMonitorDestinationIpCfg(d []interface{}) edpt.NetflowMonitorDestinationIpCfg {

	count1 := len(d)
	var ret edpt.NetflowMonitorDestinationIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(string)
		ret.Port4 = in["port4"].(int)
	}
	return ret
}

func getObjectNetflowMonitorDestinationIpv6Cfg(d []interface{}) edpt.NetflowMonitorDestinationIpv6Cfg {

	count1 := len(d)
	var ret edpt.NetflowMonitorDestinationIpv6Cfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6 = in["ipv6"].(string)
		ret.Port6 = in["port6"].(int)
	}
	return ret
}

func dataToEndpointNetflowMonitorDestination(d *schema.ResourceData) edpt.NetflowMonitorDestination {
	var ret edpt.NetflowMonitorDestination
	ret.Inst.IpCfg = getObjectNetflowMonitorDestinationIpCfg(d.Get("ip_cfg").([]interface{}))
	ret.Inst.Ipv6Cfg = getObjectNetflowMonitorDestinationIpv6Cfg(d.Get("ipv6_cfg").([]interface{}))
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
