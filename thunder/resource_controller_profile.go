package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_profile`: A10 controller profile\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerProfileCreate,
		UpdateContext: resourceControllerProfileUpdate,
		ReadContext:   resourceControllerProfileRead,
		DeleteContext: resourceControllerProfileDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'register': Register the device to the controller; 'deregister': Deregister the device from controller;",
			},
			"analytics": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'all': Export all the analytics information.; 'system': Export only system level policy for device management.; 'disable': Disable all the exports from the device. This is the default value.;",
			},
			"analytics_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use data port for analytics information exportation, if not set analytics will follow 'host use-mgmt-port' option",
			},
			"api_key": {
				Type: schema.TypeString, Optional: true, Description: "API key for authentication",
			},
			"auto_restart_action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable auto analytics bus restart, default behavior is enable; 'disable': disable auto analytics bus restart;",
			},
			"availability_zone": {
				Type: schema.TypeString, Optional: true, Description: "availablity zone of the thunder-device",
			},
			"cluster_id": {
				Type: schema.TypeString, Optional: true, Description: "id for the cluster in controller, typically an uuid",
			},
			"cluster_name": {
				Type: schema.TypeString, Optional: true, Description: "name of cluster in controller that this device is a member of",
			},
			"force": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deregister": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "forcefully deregister thunder from harmony controller",
						},
					},
				},
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "Set controller host address",
			},
			"host_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address or FQDN for the host",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "auto analytics bus restart time interval in mins, default is 3 mins",
			},
			"organization": {
				Type: schema.TypeString, Optional: true, Description: "organization for the controller",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Set port for remote Controller",
			},
			"re_sync": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schema_registry": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "re-sync the schema registry",
						},
						"analytics_bus": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "re-sync analtyics bus connections",
						},
					},
				},
			},
			"region": {
				Type: schema.TypeString, Optional: true, Description: "region of the thunder-device",
			},
			"secret_value": {
				Type: schema.TypeString, Optional: true, Description: "Specify the password for the user",
			},
			"thunder_mgmt_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address (IPv4 address)",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 address for the host",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tunnel": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Tunnel Enable; 'disable': Tunnel Disable;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections",
			},
			"user_name": {
				Type: schema.TypeString, Optional: true, Description: "user-name for the tenant",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceControllerProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectControllerProfileForce140(d []interface{}) edpt.ControllerProfileForce140 {

	count1 := len(d)
	var ret edpt.ControllerProfileForce140
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Deregister = in["deregister"].(int)
	}
	return ret
}

func getObjectControllerProfileReSync141(d []interface{}) edpt.ControllerProfileReSync141 {

	count1 := len(d)
	var ret edpt.ControllerProfileReSync141
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SchemaRegistry = in["schema_registry"].(int)
		ret.AnalyticsBus = in["analytics_bus"].(int)
	}
	return ret
}

func getObjectControllerProfileThunderMgmtIp142(d []interface{}) edpt.ControllerProfileThunderMgmtIp142 {

	count1 := len(d)
	var ret edpt.ControllerProfileThunderMgmtIp142
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpAddress = in["ip_address"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		//omit uuid
	}
	return ret
}

func getObjectControllerProfileTunnel143(d []interface{}) edpt.ControllerProfileTunnel143 {

	count1 := len(d)
	var ret edpt.ControllerProfileTunnel143
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointControllerProfile(d *schema.ResourceData) edpt.ControllerProfile {
	var ret edpt.ControllerProfile
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Analytics = d.Get("analytics").(string)
	ret.Inst.AnalyticsData = d.Get("analytics_data").(int)
	ret.Inst.ApiKey = d.Get("api_key").(string)
	ret.Inst.AutoRestartAction = d.Get("auto_restart_action").(string)
	ret.Inst.AvailabilityZone = d.Get("availability_zone").(string)
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	ret.Inst.ClusterName = d.Get("cluster_name").(string)
	ret.Inst.Force = getObjectControllerProfileForce140(d.Get("force").([]interface{}))
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.HostIpv6 = d.Get("host_ipv6").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Organization = d.Get("organization").(string)
	//omit password_encrypted
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ReSync = getObjectControllerProfileReSync141(d.Get("re_sync").([]interface{}))
	ret.Inst.Region = d.Get("region").(string)
	ret.Inst.SecretValue = d.Get("secret_value").(string)
	ret.Inst.ThunderMgmtIp = getObjectControllerProfileThunderMgmtIp142(d.Get("thunder_mgmt_ip").([]interface{}))
	ret.Inst.Tunnel = getObjectControllerProfileTunnel143(d.Get("tunnel").([]interface{}))
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserName = d.Get("user_name").(string)
	//omit uuid
	return ret
}
