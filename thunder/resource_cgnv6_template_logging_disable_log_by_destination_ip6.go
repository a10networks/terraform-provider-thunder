package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLoggingDisableLogByDestinationIp6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging_disable_log_by_destination_ip6`: Configure a filter IPv6 enrty\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Create,
		UpdateContext: resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Update,
		ReadContext:   resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Read,
		DeleteContext: resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Delete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure an IPv6 subnet",
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
func resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationIp6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6TcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6TcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6TcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6TcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6UdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6UdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6UdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6UdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6TemplateLoggingDisableLogByDestinationIp6(d *schema.ResourceData) edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6 {
	var ret edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6TcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6UdpList(d.Get("udp_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
