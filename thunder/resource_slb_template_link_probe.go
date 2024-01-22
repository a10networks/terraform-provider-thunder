package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLinkProbe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_link_probe`: Link probe template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateLinkProbeCreate,
		UpdateContext: resourceSlbTemplateLinkProbeUpdate,
		ReadContext:   resourceSlbTemplateLinkProbeRead,
		DeleteContext: resourceSlbTemplateLinkProbeDelete,

		Schema: map[string]*schema.Schema{
			"destination": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type: schema.TypeString, Optional: true, Description: "Target Hostname",
						},
						"resolve_as": {
							Type: schema.TypeString, Optional: true, Description: "'resolve-to-ipv4': Use A Query only to resolve the configured hostname; 'resolve-to-ipv6': Use AAAA Query only to resolve the configured hostname;",
						},
						"static_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "Target IPv4 Address",
						},
						"static_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Target IPv6 Address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Probe template",
			},
			"expected_status_code": {
				Type: schema.TypeString, Optional: true, Default: "200", Description: "Specify response code range (e.g. 200,400-430), default is 200 (Format is xx,xx-xx (xx between [100, 899]), default is 200)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "probe template Name",
			},
			"probe_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between each probe that needs to be sent out (in seconds, default is 5)",
			},
			"probes_per_test": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Total number of probes to be sent out for each test",
			},
			"rtt_method": {
				Type: schema.TypeString, Optional: true, Default: "http-rtt", Description: "'http-rtt': Calculate Round Trip Time between HTTP request and response; 'tcp-srtt': Use TCP Smoothed round trip time in the HTTP connection;",
			},
			"selection_rule": {
				Type: schema.TypeString, Optional: true, Default: "fastest-link-always", Description: "'threshold': Use all links below a threshold before selecting the fastest link; 'fastest-link-always': Always use the link with the lowest average latency;",
			},
			"test_interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "time interval between subsequent tests (in seconds, default is 60)",
			},
			"threshold_value": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use all links below a threshold before selecting the fastest link (RTT in milliseconds)",
			},
			"url": {
				Type: schema.TypeString, Optional: true, Default: "/", Description: "Specify URL to which probes should be sent out. Default is /",
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
func resourceSlbTemplateLinkProbeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkProbeRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateLinkProbeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLinkProbeRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateLinkProbeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateLinkProbeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLinkProbeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLinkProbe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateLinkProbeDestination1449(d []interface{}) edpt.SlbTemplateLinkProbeDestination1449 {

	count1 := len(d)
	var ret edpt.SlbTemplateLinkProbeDestination1449
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostname = in["hostname"].(string)
		ret.ResolveAs = in["resolve_as"].(string)
		ret.StaticIpv4Addr = in["static_ipv4_addr"].(string)
		ret.StaticIpv6Addr = in["static_ipv6_addr"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointSlbTemplateLinkProbe(d *schema.ResourceData) edpt.SlbTemplateLinkProbe {
	var ret edpt.SlbTemplateLinkProbe
	ret.Inst.Destination = getObjectSlbTemplateLinkProbeDestination1449(d.Get("destination").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.ExpectedStatusCode = d.Get("expected_status_code").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ProbeInterval = d.Get("probe_interval").(int)
	ret.Inst.ProbesPerTest = d.Get("probes_per_test").(int)
	ret.Inst.RttMethod = d.Get("rtt_method").(string)
	ret.Inst.SelectionRule = d.Get("selection_rule").(string)
	ret.Inst.TestInterval = d.Get("test_interval").(int)
	ret.Inst.ThresholdValue = d.Get("threshold_value").(int)
	ret.Inst.Url = d.Get("url").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
