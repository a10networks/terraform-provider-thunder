package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLoggingDisableLogByDestinationIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging_disable_log_by_destination_ip`: Configure a filter of IP entry\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingDisableLogByDestinationIpCreate,
		UpdateContext: resourceFwTemplateLoggingDisableLogByDestinationIpUpdate,
		ReadContext:   resourceFwTemplateLoggingDisableLogByDestinationIpRead,
		DeleteContext: resourceFwTemplateLoggingDisableLogByDestinationIpDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the ICMP traffic",
			},
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure an IP subnet",
			},
			"others": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the other layer-4 protocols",
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
func resourceFwTemplateLoggingDisableLogByDestinationIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestinationIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingDisableLogByDestinationIpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingDisableLogByDestinationIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestinationIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingDisableLogByDestinationIpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingDisableLogByDestinationIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestinationIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingDisableLogByDestinationIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestinationIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTemplateLoggingDisableLogByDestinationIpTcpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIpTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIpTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIpTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationIpUdpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIpUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIpUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIpUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTemplateLoggingDisableLogByDestinationIp(d *schema.ResourceData) edpt.FwTemplateLoggingDisableLogByDestinationIp {
	var ret edpt.FwTemplateLoggingDisableLogByDestinationIp
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceFwTemplateLoggingDisableLogByDestinationIpTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceFwTemplateLoggingDisableLogByDestinationIpUdpList(d.Get("udp_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Logging_name = d.Get("logging_name").(string)
	return ret
}
