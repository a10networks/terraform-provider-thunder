package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatTranslation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_translation`: Change or Disable NAT translation values\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatTranslationCreate,
		UpdateContext: resourceIpNatTranslationUpdate,
		ReadContext:   resourceIpNatTranslationRead,
		DeleteContext: resourceIpNatTranslationDelete,

		Schema: map[string]*schema.Schema{
			"icmp_timeout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_timeout": {
							Type: schema.TypeString, Optional: true, Default: "fast", Description: "'age': Expiration time; 'fast': Use Fast aging;",
						},
						"icmp_timeout_val": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds (Interval of 60 seconds)",
						},
					},
				},
			},
			"ignore_tcp_msl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "reclaim TCP resource immediately without MSL",
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
						"timeout_type": {
							Type: schema.TypeString, Optional: true, Description: "'age': Expiration time; 'fast': Use Fast aging;",
						},
						"timeout_val": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds (Interval of 60 seconds)",
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
func resourceIpNatTranslationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatTranslationRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatTranslationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatTranslationRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatTranslationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatTranslationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatTranslationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatTranslation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpNatTranslationIcmpTimeout(d []interface{}) edpt.IpNatTranslationIcmpTimeout {

	count1 := len(d)
	var ret edpt.IpNatTranslationIcmpTimeout
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpTimeout = in["icmp_timeout"].(string)
		ret.IcmpTimeoutVal = in["icmp_timeout_val"].(int)
	}
	return ret
}

func getSliceIpNatTranslationServiceTimeoutList(d []interface{}) []edpt.IpNatTranslationServiceTimeoutList {

	count1 := len(d)
	ret := make([]edpt.IpNatTranslationServiceTimeoutList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpNatTranslationServiceTimeoutList
		oi.ServiceType = in["service_type"].(string)
		oi.Port = in["port"].(int)
		oi.TimeoutType = in["timeout_type"].(string)
		oi.TimeoutVal = in["timeout_val"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpNatTranslation(d *schema.ResourceData) edpt.IpNatTranslation {
	var ret edpt.IpNatTranslation
	ret.Inst.IcmpTimeout = getObjectIpNatTranslationIcmpTimeout(d.Get("icmp_timeout").([]interface{}))
	ret.Inst.IgnoreTcpMsl = d.Get("ignore_tcp_msl").(int)
	ret.Inst.ServiceTimeoutList = getSliceIpNatTranslationServiceTimeoutList(d.Get("service_timeout_list").([]interface{}))
	ret.Inst.TcpTimeout = d.Get("tcp_timeout").(int)
	ret.Inst.UdpTimeout = d.Get("udp_timeout").(int)
	//omit uuid
	return ret
}
