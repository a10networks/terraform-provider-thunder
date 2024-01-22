package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_template_logging`: NAT Logging Template\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatTemplateLoggingCreate,
		UpdateContext: resourceIpNatTemplateLoggingUpdate,
		ReadContext:   resourceIpNatTemplateLoggingRead,
		DeleteContext: resourceIpNatTemplateLoggingDelete,

		Schema: map[string]*schema.Schema{
			"facility": {
				Type: schema.TypeString, Optional: true, Default: "local0", Description: "'kernel': 0: Kernel; 'user': 1: User-level; 'mail': 2: Mail; 'daemon': 3: System daemons; 'security-authorization': 4: Security/authorization; 'syslog': 5: Syslog internal; 'line-printer': 6: Line printer; 'news': 7: Network news; 'uucp': 8: UUCP subsystem; 'cron': 9: Time-related; 'security-authorization-private': 10: Private security/authorization; 'ftp': 11: FTP; 'ntp': 12: NTP; 'audit': 13: Audit; 'alert': 14: Alert; 'clock': 15: Clock-related; 'local0': 16: Local use 0; 'local1': 17: Local use 1; 'local2': 18: Local use 2; 'local3': 19: Local use 3; 'local4': 20: Local use 4; 'local5': 21: Local use 5; 'local6': 22: Local use 6; 'local7': 23: Local use 7;",
			},
			"include_destination": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the destination IP and port in logs",
			},
			"include_rip_rport": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the IP and port of real server in logs",
			},
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_mappings": {
							Type: schema.TypeString, Optional: true, Description: "'creation': Log creation of NAT mappgins; 'disable': Disable Log creation and deletion of NAT mappings;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "NAT logging template name",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Set NAT logging service-group",
			},
			"severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity_string": {
							Type: schema.TypeString, Optional: true, Default: "debug", Description: "'emergency': 0: Emergency; 'alert': 1: Alert; 'critical': 2: Critical; 'error': 3: Error; 'warning': 4: Warning; 'notice': 5: Notice; 'informational': 6: Informational; 'debug': 7: Debug;",
						},
						"severity_val": {
							Type: schema.TypeInt, Optional: true, Default: 7, Description: "Logging severity level",
						},
					},
				},
			},
			"source_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_port_num": {
							Type: schema.TypeInt, Optional: true, Default: 514, Description: "Set source port for sending NAT syslogs (default: 514)",
						},
						"any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use any source port",
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
func resourceIpNatTemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTemplateLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTemplateLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatTemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTemplateLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTemplateLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatTemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTemplateLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTemplateLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatTemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTemplateLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTemplateLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpNatTemplateLoggingLog(d []interface{}) edpt.IpNatTemplateLoggingLog {

	count1 := len(d)
	var ret edpt.IpNatTemplateLoggingLog
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortMappings = in["port_mappings"].(string)
	}
	return ret
}

func getObjectIpNatTemplateLoggingSeverity(d []interface{}) edpt.IpNatTemplateLoggingSeverity {

	count1 := len(d)
	var ret edpt.IpNatTemplateLoggingSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SeverityString = in["severity_string"].(string)
		ret.SeverityVal = in["severity_val"].(int)
	}
	return ret
}

func getObjectIpNatTemplateLoggingSourcePort(d []interface{}) edpt.IpNatTemplateLoggingSourcePort {

	count1 := len(d)
	var ret edpt.IpNatTemplateLoggingSourcePort
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SourcePortNum = in["source_port_num"].(int)
		ret.Any = in["any"].(int)
	}
	return ret
}

func dataToEndpointIpNatTemplateLogging(d *schema.ResourceData) edpt.IpNatTemplateLogging {
	var ret edpt.IpNatTemplateLogging
	ret.Inst.Facility = d.Get("facility").(string)
	ret.Inst.IncludeDestination = d.Get("include_destination").(int)
	ret.Inst.IncludeRipRport = d.Get("include_rip_rport").(int)
	ret.Inst.Log = getObjectIpNatTemplateLoggingLog(d.Get("log").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.Severity = getObjectIpNatTemplateLoggingSeverity(d.Get("severity").([]interface{}))
	ret.Inst.SourcePort = getObjectIpNatTemplateLoggingSourcePort(d.Get("source_port").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
