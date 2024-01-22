package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Translation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_translation`: Configure CGN translation timeout values\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TranslationCreate,
		UpdateContext: resourceCgnv6TranslationUpdate,
		ReadContext:   resourceCgnv6TranslationRead,
		DeleteContext: resourceCgnv6TranslationDelete,

		Schema: map[string]*schema.Schema{
			"icmp_timeout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_timeout_val": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds (Interval of 60 seconds)",
						},
						"fast": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Fast Aging",
						},
					},
				},
			},
			"service_timeout_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_type": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Protocol; 'udp': UDP Protocol;",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"port_end": {
							Type: schema.TypeInt, Optional: true, Description: "End Port Number",
						},
						"timeout_val": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds (Interval of 60 seconds)",
						},
						"fast": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Fast Aging",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tcp_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "TCP protocol extended translations (Timeout in seconds (Interval of 60 seconds), default is 300 seconds (5 minutes))",
			},
			"udp_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "UDP protocol extended translations (Timeout in seconds (Interval of 60 seconds), default is 300 seconds (5 minutes))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6TranslationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Translation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TranslationRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TranslationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Translation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TranslationRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TranslationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Translation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TranslationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TranslationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Translation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6TranslationIcmpTimeout(d []interface{}) edpt.Cgnv6TranslationIcmpTimeout {

	count1 := len(d)
	var ret edpt.Cgnv6TranslationIcmpTimeout
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpTimeoutVal = in["icmp_timeout_val"].(int)
		ret.Fast = in["fast"].(int)
	}
	return ret
}

func getSliceCgnv6TranslationServiceTimeoutList(d []interface{}) []edpt.Cgnv6TranslationServiceTimeoutList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TranslationServiceTimeoutList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TranslationServiceTimeoutList
		oi.ServiceType = in["service_type"].(string)
		oi.Port = in["port"].(int)
		oi.PortEnd = in["port_end"].(int)
		oi.TimeoutVal = in["timeout_val"].(int)
		oi.Fast = in["fast"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Translation(d *schema.ResourceData) edpt.Cgnv6Translation {
	var ret edpt.Cgnv6Translation
	ret.Inst.IcmpTimeout = getObjectCgnv6TranslationIcmpTimeout(d.Get("icmp_timeout").([]interface{}))
	ret.Inst.ServiceTimeoutList = getSliceCgnv6TranslationServiceTimeoutList(d.Get("service_timeout_list").([]interface{}))
	ret.Inst.TcpTimeout = d.Get("tcp_timeout").(int)
	ret.Inst.UdpTimeout = d.Get("udp_timeout").(int)
	//omit uuid
	return ret
}
