package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelThreatList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_threat_intel_threat_list`: Threat Categories for malicious IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceThreatIntelThreatListCreate,
		UpdateContext: resourceThreatIntelThreatListUpdate,
		ReadContext:   resourceThreatIntelThreatListRead,
		DeleteContext: resourceThreatIntelThreatListDelete,

		Schema: map[string]*schema.Schema{
			"all_categories": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all categories",
			},
			"botnets": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Botnet C&C channels, and infected zombie machines controlled by Bot master",
			},
			"dos_attacks": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's participating in DOS, DDOS, anomalous sync flood, and anomalous traffic detection",
			},
			"mobile_threats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's associated with mobile threats",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Threat category List name",
			},
			"phishing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addresses hosting phishing sites, ad click fraud or gaming fraud",
			},
			"proxy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addresses providing proxy services",
			},
			"reputation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP addresses currently known to be infected with malware",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'spam-sources': Hits for spam sources; 'windows-exploits': Hits for windows exploits; 'web-attacks': Hits for web attacks; 'botnets': Hits for botnets; 'scanners': Hits for scanners; 'dos-attacks': Hits for dos attacks; 'reputation': Hits for reputation; 'phishing': Hits for phishing; 'proxy': Hits for proxy; 'mobile-threats': Hits for mobile threats; 'tor-proxy': Hits for tor-proxy; 'total-hits': Total hits for threat-list;",
						},
					},
				},
			},
			"scanners": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's associated with probes, host scan, domain scan, and password brute force attack",
			},
			"spam_sources": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's tunneling spam messages through a proxy, anomalous SMTP activities, and forum spam activities",
			},
			"tor_proxy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's providing tor proxy services",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'webroot': Configure Webroot threat categories;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"web_attacks": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's associated with cross site scripting, iFrame injection, SQL injection, cross domain injection, or domain password brute fo",
			},
			"windows_exploits": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP's associated with malware, shell code, rootkits, worms or viruses",
			},
		},
	}
}
func resourceThreatIntelThreatListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceThreatIntelThreatListRead(ctx, d, meta)
	}
	return diags
}

func resourceThreatIntelThreatListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceThreatIntelThreatListRead(ctx, d, meta)
	}
	return diags
}
func resourceThreatIntelThreatListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceThreatIntelThreatListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceThreatIntelThreatListSamplingEnable(d []interface{}) []edpt.ThreatIntelThreatListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.ThreatIntelThreatListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ThreatIntelThreatListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointThreatIntelThreatList(d *schema.ResourceData) edpt.ThreatIntelThreatList {
	var ret edpt.ThreatIntelThreatList
	ret.Inst.AllCategories = d.Get("all_categories").(int)
	ret.Inst.Botnets = d.Get("botnets").(int)
	ret.Inst.DosAttacks = d.Get("dos_attacks").(int)
	ret.Inst.MobileThreats = d.Get("mobile_threats").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Phishing = d.Get("phishing").(int)
	ret.Inst.Proxy = d.Get("proxy").(int)
	ret.Inst.Reputation = d.Get("reputation").(int)
	ret.Inst.SamplingEnable = getSliceThreatIntelThreatListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Scanners = d.Get("scanners").(int)
	ret.Inst.SpamSources = d.Get("spam_sources").(int)
	ret.Inst.TorProxy = d.Get("tor_proxy").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.WebAttacks = d.Get("web_attacks").(int)
	ret.Inst.WindowsExploits = d.Get("windows_exploits").(int)
	return ret
}
