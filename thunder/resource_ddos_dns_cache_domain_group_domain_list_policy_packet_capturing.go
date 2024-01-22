package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_domain_group_domain_list_policy_packet_capturing`: Capture packet trace of specified root zones when doing zone transfer\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingCreate,
		UpdateContext: resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingUpdate,
		ReadContext:   resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRead,
		DeleteContext: resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingDelete,

		Schema: map[string]*schema.Schema{
			"root_zone_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"root_zone": {
							Type: schema.TypeString, Optional: true, Description: "Specify root zone to be captured",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Capture-config name",
						},
						"capture_mode": {
							Type: schema.TypeString, Optional: true, Description: "'regular': Capture packet anyway; 'capture-on-failure': Capture packet if last XFR was failed;",
						},
					},
				},
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
func resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList(d []interface{}) []edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList
		oi.RootZone = in["root_zone"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		oi.CaptureMode = in["capture_mode"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheDomainGroupDomainListPolicyPacketCapturing(d *schema.ResourceData) edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing {
	var ret edpt.DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing
	ret.Inst.RootZoneList = getSliceDdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList(d.Get("root_zone_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
