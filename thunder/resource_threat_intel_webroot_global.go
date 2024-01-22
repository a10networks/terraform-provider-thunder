package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelWebrootGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_threat_intel_webroot_global`: Global counters for webroot module\n\n__PLACEHOLDER__",
		CreateContext: resourceThreatIntelWebrootGlobalCreate,
		UpdateContext: resourceThreatIntelWebrootGlobalUpdate,
		ReadContext:   resourceThreatIntelWebrootGlobalRead,
		DeleteContext: resourceThreatIntelWebrootGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'spam-sources': Hits for spam sources; 'windows-exploits': Hits for windows exploits; 'web-attacks': Hits for web attacks; 'botnets': Hits for botnets; 'scanners': Hits for scanners; 'dos-attacks': Hits for dos attacks; 'reputation': Hits for reputation; 'phishing': Hits for phishing; 'proxy': Hits for proxy; 'mobile-threats': Hits for mobile threats; 'tor-proxy': Hits for tor-proxy; 'rtu-lookup': Number of lookups in RTU cache; 'database-lookup': Number of lookups in database; 'non-malicious-ips': IP's not found in database or RTU cache;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceThreatIntelWebrootGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceThreatIntelWebrootGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceThreatIntelWebrootGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceThreatIntelWebrootGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceThreatIntelWebrootGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceThreatIntelWebrootGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceThreatIntelWebrootGlobalSamplingEnable(d []interface{}) []edpt.ThreatIntelWebrootGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.ThreatIntelWebrootGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ThreatIntelWebrootGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointThreatIntelWebrootGlobal(d *schema.ResourceData) edpt.ThreatIntelWebrootGlobal {
	var ret edpt.ThreatIntelWebrootGlobal
	ret.Inst.SamplingEnable = getSliceThreatIntelWebrootGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
