package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLocalDnsResolution() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_local_dns_resolution`: local DNS resolver configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLocalDnsResolutionCreate,
		UpdateContext: resourceSlbTemplateDnsLocalDnsResolutionUpdate,
		ReadContext:   resourceSlbTemplateDnsLocalDnsResolutionRead,
		DeleteContext: resourceSlbTemplateDnsLocalDnsResolutionDelete,

		Schema: map[string]*schema.Schema{
			"host_list_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostnames": {
							Type: schema.TypeString, Optional: true, Description: "Hostnames class-list name (dns type)",
						},
					},
				},
			},
			"local_resolver_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_resolver": {
							Type: schema.TypeString, Optional: true, Description: "Local dns servers (address)",
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
func resourceSlbTemplateDnsLocalDnsResolutionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLocalDnsResolutionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLocalDnsResolution(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLocalDnsResolutionRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLocalDnsResolutionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLocalDnsResolutionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLocalDnsResolution(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLocalDnsResolutionRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLocalDnsResolutionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLocalDnsResolutionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLocalDnsResolution(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLocalDnsResolutionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLocalDnsResolutionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLocalDnsResolution(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg
		oi.LocalResolver = in["local_resolver"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsLocalDnsResolution(d *schema.ResourceData) edpt.SlbTemplateDnsLocalDnsResolution {
	var ret edpt.SlbTemplateDnsLocalDnsResolution
	ret.Inst.HostListCfg = getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg(d.Get("host_list_cfg").([]interface{}))
	ret.Inst.LocalResolverCfg = getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg(d.Get("local_resolver_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
