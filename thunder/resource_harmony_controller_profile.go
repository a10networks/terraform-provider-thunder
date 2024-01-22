package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_profile`: Harmony controller profile\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerProfileCreate,
		UpdateContext: resourceHarmonyControllerProfileUpdate,
		ReadContext:   resourceHarmonyControllerProfileRead,
		DeleteContext: resourceHarmonyControllerProfileDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'register': Register the device to the controller; 'deregister': Deregister the device from controller;",
			},
			"analytics": {
				Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': Export all the analytics information. This is the default value.; 'system': Export only system level policy for device management.; 'disable': Disable all the exports from the device.;",
			},
			"auto_restart_action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable auto analytics bus restart, default behavior is enable; 'disable': disable auto analytics bus restart;",
			},
			"availability_zone": {
				Type: schema.TypeString, Optional: true, Description: "availablity zone of the thunder-device",
			},
			"cluster_id": {
				Type: schema.TypeString, Optional: true, Description: "id for the cluster in harmony controller, typically an uuid",
			},
			"cluster_name": {
				Type: schema.TypeString, Optional: true, Description: "name of cluster in harmony controller that this device is a member of",
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "Set harmony controller host address",
			},
			"host_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address or FQDN for the host",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "auto analytics bus restart time interval in mins, default is 3 mins",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Set port for remote Harmony Controller",
			},
			"provider1": {
				Type: schema.TypeString, Optional: true, Description: "provider for the harmony-controller",
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
func resourceHarmonyControllerProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHarmonyControllerProfileReSync400(d []interface{}) edpt.HarmonyControllerProfileReSync400 {

	count1 := len(d)
	var ret edpt.HarmonyControllerProfileReSync400
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SchemaRegistry = in["schema_registry"].(int)
		ret.AnalyticsBus = in["analytics_bus"].(int)
	}
	return ret
}

func getObjectHarmonyControllerProfileThunderMgmtIp401(d []interface{}) edpt.HarmonyControllerProfileThunderMgmtIp401 {

	count1 := len(d)
	var ret edpt.HarmonyControllerProfileThunderMgmtIp401
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpAddress = in["ip_address"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		//omit uuid
	}
	return ret
}

func getObjectHarmonyControllerProfileTunnel402(d []interface{}) edpt.HarmonyControllerProfileTunnel402 {

	count1 := len(d)
	var ret edpt.HarmonyControllerProfileTunnel402
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointHarmonyControllerProfile(d *schema.ResourceData) edpt.HarmonyControllerProfile {
	var ret edpt.HarmonyControllerProfile
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Analytics = d.Get("analytics").(string)
	ret.Inst.AutoRestartAction = d.Get("auto_restart_action").(string)
	ret.Inst.AvailabilityZone = d.Get("availability_zone").(string)
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	ret.Inst.ClusterName = d.Get("cluster_name").(string)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.HostIpv6 = d.Get("host_ipv6").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	//omit password_encrypted
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Provider1 = d.Get("provider1").(string)
	ret.Inst.ReSync = getObjectHarmonyControllerProfileReSync400(d.Get("re_sync").([]interface{}))
	ret.Inst.Region = d.Get("region").(string)
	ret.Inst.SecretValue = d.Get("secret_value").(string)
	ret.Inst.ThunderMgmtIp = getObjectHarmonyControllerProfileThunderMgmtIp401(d.Get("thunder_mgmt_ip").([]interface{}))
	ret.Inst.Tunnel = getObjectHarmonyControllerProfileTunnel402(d.Get("tunnel").([]interface{}))
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserName = d.Get("user_name").(string)
	//omit uuid
	return ret
}
