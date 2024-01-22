package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseManager() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_license_manager`: Configure license manager\n\n__PLACEHOLDER__",
		CreateContext: resourceLicenseManagerCreate,
		UpdateContext: resourceLicenseManagerUpdate,
		ReadContext:   resourceLicenseManagerRead,
		DeleteContext: resourceLicenseManagerDelete,

		Schema: map[string]*schema.Schema{
			"bandwidth_base": {
				Type: schema.TypeInt, Optional: true, Description: "Configure feature bandwidth base (Mb)",
			},
			"bandwidth_unrestricted": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the bandwidth to maximum",
			},
			"connect": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connect": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Connect to license manager to activate",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"host_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_ipv4": {
							Type: schema.TypeString, Required: true, Description: "license server ip address (length:1-31)",
						},
						"host_ipv6": {
							Type: schema.TypeString, Required: true, Description: "Configure license manager server ipv6-address",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Default: 443, Description: "Configure the license manager port, default is 443",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"instance_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure instance name [format: (string).(string).(string).(string)]",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Configure interval profile (1 monthly, 2 daily, 3 hourly)",
			},
			"ng_waf_module": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key_id": {
							Type: schema.TypeString, Optional: true, Description: "access-key",
						},
						"secret_access_key": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"overage": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": {
							Type: schema.TypeInt, Optional: true, Description: "Number of days",
						},
						"hours": {
							Type: schema.TypeInt, Optional: true, Description: "Number of hours",
						},
						"minutes": {
							Type: schema.TypeInt, Optional: true, Description: "Number of minutes",
						},
						"seconds": {
							Type: schema.TypeInt, Optional: true, Description: "Number of seconds",
						},
						"gb": {
							Type: schema.TypeInt, Optional: true, Description: "Number of GB",
						},
						"mb": {
							Type: schema.TypeInt, Optional: true, Description: "Number of MB",
						},
						"kb": {
							Type: schema.TypeInt, Optional: true, Description: "Number of KB",
						},
						"bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Number of bytes",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"reminder_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reminder_value": {
							Type: schema.TypeInt, Required: true, Description: "Configure reminder for grace time (Hour)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sn": {
				Type: schema.TypeString, Optional: true, Description: "serial number",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect license server",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLicenseManagerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManager(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerRead(ctx, d, meta)
	}
	return diags
}

func resourceLicenseManagerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManager(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLicenseManagerRead(ctx, d, meta)
	}
	return diags
}
func resourceLicenseManagerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManager(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLicenseManagerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLicenseManagerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLicenseManager(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectLicenseManagerConnect1046(d []interface{}) edpt.LicenseManagerConnect1046 {

	count1 := len(d)
	var ret edpt.LicenseManagerConnect1046
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connect = in["connect"].(int)
		//omit uuid
	}
	return ret
}

func getSliceLicenseManagerHostList(d []interface{}) []edpt.LicenseManagerHostList {

	count1 := len(d)
	ret := make([]edpt.LicenseManagerHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LicenseManagerHostList
		oi.HostIpv4 = in["host_ipv4"].(string)
		oi.HostIpv6 = in["host_ipv6"].(string)
		oi.Port = in["port"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectLicenseManagerNgWafModule1047(d []interface{}) edpt.LicenseManagerNgWafModule1047 {

	count1 := len(d)
	var ret edpt.LicenseManagerNgWafModule1047
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AccessKeyId = in["access_key_id"].(string)
		ret.SecretAccessKey = in["secret_access_key"].(string)
	}
	return ret
}

func getObjectLicenseManagerOverage1048(d []interface{}) edpt.LicenseManagerOverage1048 {

	count1 := len(d)
	var ret edpt.LicenseManagerOverage1048
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Days = in["days"].(int)
		ret.Hours = in["hours"].(int)
		ret.Minutes = in["minutes"].(int)
		ret.Seconds = in["seconds"].(int)
		ret.Gb = in["gb"].(int)
		ret.Mb = in["mb"].(int)
		ret.Kb = in["kb"].(int)
		ret.Bytes = in["bytes"].(int)
		//omit uuid
	}
	return ret
}

func getSliceLicenseManagerReminderList(d []interface{}) []edpt.LicenseManagerReminderList {

	count1 := len(d)
	ret := make([]edpt.LicenseManagerReminderList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LicenseManagerReminderList
		oi.ReminderValue = in["reminder_value"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLicenseManager(d *schema.ResourceData) edpt.LicenseManager {
	var ret edpt.LicenseManager
	ret.Inst.BandwidthBase = d.Get("bandwidth_base").(int)
	ret.Inst.BandwidthUnrestricted = d.Get("bandwidth_unrestricted").(int)
	ret.Inst.Connect = getObjectLicenseManagerConnect1046(d.Get("connect").([]interface{}))
	ret.Inst.HostList = getSliceLicenseManagerHostList(d.Get("host_list").([]interface{}))
	ret.Inst.InstanceName = d.Get("instance_name").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.NgWafModule = getObjectLicenseManagerNgWafModule1047(d.Get("ng_waf_module").([]interface{}))
	ret.Inst.Overage = getObjectLicenseManagerOverage1048(d.Get("overage").([]interface{}))
	ret.Inst.ReminderList = getSliceLicenseManagerReminderList(d.Get("reminder_list").([]interface{}))
	ret.Inst.Sn = d.Get("sn").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
