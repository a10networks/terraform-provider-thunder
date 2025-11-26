package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLoggingEnableLogByDestinationIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging_enable_log_by_destination_ip`: Configure a filter of IP entry\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingEnableLogByDestinationIpCreate,
		UpdateContext: resourceCgnv6TemplateLoggingEnableLogByDestinationIpUpdate,
		ReadContext:   resourceCgnv6TemplateLoggingEnableLogByDestinationIpRead,
		DeleteContext: resourceCgnv6TemplateLoggingEnableLogByDestinationIpDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the ICMP traffic",
			},
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure an IP subnet",
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
func resourceCgnv6TemplateLoggingEnableLogByDestinationIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestinationIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingEnableLogByDestinationIpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingEnableLogByDestinationIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestinationIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingEnableLogByDestinationIpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingEnableLogByDestinationIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestinationIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingEnableLogByDestinationIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestinationIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIpTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIpUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6TemplateLoggingEnableLogByDestinationIp(d *schema.ResourceData) edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp {
	var ret edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIpTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIpUdpList(d.Get("udp_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Logging_name = d.Get("logging_name").(string)
	return ret
}
