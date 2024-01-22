package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLoggingDisableLogByDestinationIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging_disable_log_by_destination_ip`: Configure a filter IP enrty\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingDisableLogByDestinationIpCreate,
		UpdateContext: resourceCgnv6TemplateLoggingDisableLogByDestinationIpUpdate,
		ReadContext:   resourceCgnv6TemplateLoggingDisableLogByDestinationIpRead,
		DeleteContext: resourceCgnv6TemplateLoggingDisableLogByDestinationIpDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
			},
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure an IP subnet",
			},
			"others": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6TemplateLoggingDisableLogByDestinationIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingDisableLogByDestinationIpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingDisableLogByDestinationIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingDisableLogByDestinationIpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingDisableLogByDestinationIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingDisableLogByDestinationIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp(d *schema.ResourceData) edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp {
	var ret edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpUdpList(d.Get("udp_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
