package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLoggingEnableLogByDestinationIp6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging_enable_log_by_destination_ip6`: Configure a filter of IPv6 entry\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingEnableLogByDestinationIp6Create,
		UpdateContext: resourceFwTemplateLoggingEnableLogByDestinationIp6Update,
		ReadContext:   resourceFwTemplateLoggingEnableLogByDestinationIp6Read,
		DeleteContext: resourceFwTemplateLoggingEnableLogByDestinationIp6Delete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the ICMP traffic",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure an IPv6 subnet",
			},
			"others": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the other layer-4 protocols",
			},
			"tcp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
						},
						"tcp_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Port Range End",
						},
					},
				},
			},
			"udp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
						},
						"udp_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Port Range End",
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
			"logging_name": {
				Type: schema.TypeString, Required: true, Description: "Logging_name",
			},
		},
	}
}
func resourceFwTemplateLoggingEnableLogByDestinationIp6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationIp6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestinationIp6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingEnableLogByDestinationIp6Read(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingEnableLogByDestinationIp6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationIp6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestinationIp6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingEnableLogByDestinationIp6Read(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingEnableLogByDestinationIp6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationIp6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestinationIp6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingEnableLogByDestinationIp6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationIp6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestinationIp6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTemplateLoggingEnableLogByDestinationIp6TcpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIp6TcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIp6TcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIp6TcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationIp6UdpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIp6UdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIp6UdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIp6UdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTemplateLoggingEnableLogByDestinationIp6(d *schema.ResourceData) edpt.FwTemplateLoggingEnableLogByDestinationIp6 {
	var ret edpt.FwTemplateLoggingEnableLogByDestinationIp6
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceFwTemplateLoggingEnableLogByDestinationIp6TcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceFwTemplateLoggingEnableLogByDestinationIp6UdpList(d.Get("udp_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Logging_name = d.Get("logging_name").(string)
	return ret
}
