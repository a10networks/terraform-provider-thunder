package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpmi() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipmi`: Perform IPMI related operations\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpmiCreate,
		UpdateContext: resourceSystemIpmiUpdate,
		ReadContext:   resourceSystemIpmiRead,
		DeleteContext: resourceSystemIpmiDelete,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
						},
						"default_gateway": {
							Type: schema.TypeString, Optional: true, Description: "Default gateway address",
						},
					},
				},
			},
			"ipsrc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dhcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addr obtained by BMC running DHCP",
						},
						"static": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually configured static IP address",
						},
					},
				},
			},
			"reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset IPMI Controller",
			},
			"tool": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cmd": {
							Type: schema.TypeString, Optional: true, Description: "Command to execute in double quotes",
						},
					},
				},
			},
			"user": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add": {
							Type: schema.TypeString, Optional: true, Description: "Add a new IPMI user (IPMI User Name)",
						},
						"password": {
							Type: schema.TypeString, Optional: true, Description: "Password",
						},
						"administrator": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Full control",
						},
						"callback": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Lowest privilege level",
						},
						"operator": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most BMC commands are allowed",
						},
						"user": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only 'benign' commands are allowed",
						},
						"disable": {
							Type: schema.TypeString, Optional: true, Description: "Disable an existing IPMI user (IPMI User Name)",
						},
						"privilege": {
							Type: schema.TypeString, Optional: true, Description: "Change an existing IPMI user privilege (IPMI User Name)",
						},
						"setname": {
							Type: schema.TypeString, Optional: true, Description: "Change User Name (Current IPMI User Name)",
						},
						"newname": {
							Type: schema.TypeString, Optional: true, Description: "New IPMI User Name",
						},
						"setpass": {
							Type: schema.TypeString, Optional: true, Description: "Change Password (IPMI User Name)",
						},
						"newpass": {
							Type: schema.TypeString, Optional: true, Description: "New Password",
						},
					},
				},
			},
		},
	}
}
func resourceSystemIpmiCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmi(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpmiUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmi(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpmiDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmi(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpmiRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmi(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemIpmiIp1588(d []interface{}) edpt.SystemIpmiIp1588 {

	count1 := len(d)
	var ret edpt.SystemIpmiIp1588
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.DefaultGateway = in["default_gateway"].(string)
	}
	return ret
}

func getObjectSystemIpmiIpsrc1589(d []interface{}) edpt.SystemIpmiIpsrc1589 {

	count1 := len(d)
	var ret edpt.SystemIpmiIpsrc1589
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.Static = in["static"].(int)
	}
	return ret
}

func getObjectSystemIpmiTool1590(d []interface{}) edpt.SystemIpmiTool1590 {

	count1 := len(d)
	var ret edpt.SystemIpmiTool1590
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cmd = in["cmd"].(string)
	}
	return ret
}

func getObjectSystemIpmiUser1591(d []interface{}) edpt.SystemIpmiUser1591 {

	count1 := len(d)
	var ret edpt.SystemIpmiUser1591
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Add = in["add"].(string)
		ret.Password = in["password"].(string)
		ret.Administrator = in["administrator"].(int)
		ret.Callback = in["callback"].(int)
		ret.Operator = in["operator"].(int)
		ret.User = in["user"].(int)
		ret.Disable = in["disable"].(string)
		ret.Privilege = in["privilege"].(string)
		ret.Setname = in["setname"].(string)
		ret.Newname = in["newname"].(string)
		ret.Setpass = in["setpass"].(string)
		ret.Newpass = in["newpass"].(string)
	}
	return ret
}

func dataToEndpointSystemIpmi(d *schema.ResourceData) edpt.SystemIpmi {
	var ret edpt.SystemIpmi
	ret.Inst.Ip = getObjectSystemIpmiIp1588(d.Get("ip").([]interface{}))
	ret.Inst.Ipsrc = getObjectSystemIpmiIpsrc1589(d.Get("ipsrc").([]interface{}))
	ret.Inst.Reset = d.Get("reset").(int)
	ret.Inst.Tool = getObjectSystemIpmiTool1590(d.Get("tool").([]interface{}))
	ret.Inst.User = getObjectSystemIpmiUser1591(d.Get("user").([]interface{}))
	return ret
}
