package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6GlobalDomainList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_global_domain_list`: Configure the global parameters for the domain-list\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6GlobalDomainListCreate,
		UpdateContext: resourceCgnv6GlobalDomainListUpdate,
		ReadContext:   resourceCgnv6GlobalDomainListRead,
		DeleteContext: resourceCgnv6GlobalDomainListDelete,

		Schema: map[string]*schema.Schema{
			"aaaa_query": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable the DNS AAAA query for each domain in the domain list.; 'disable': Disable the DNS AAAA query for each domain in the domain list.;",
			},
			"fail_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Set up the global failure interval in second for the DNS resolution.",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Set up the global query interval in minute for the DNS resolution.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6GlobalDomainListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalDomainListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6GlobalDomainList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6GlobalDomainListRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6GlobalDomainListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalDomainListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6GlobalDomainList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6GlobalDomainListRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6GlobalDomainListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalDomainListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6GlobalDomainList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6GlobalDomainListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalDomainListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6GlobalDomainList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6GlobalDomainList(d *schema.ResourceData) edpt.Cgnv6GlobalDomainList {
	var ret edpt.Cgnv6GlobalDomainList
	ret.Inst.AaaaQuery = d.Get("aaaa_query").(string)
	ret.Inst.FailInterval = d.Get("fail_interval").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	//omit uuid
	return ret
}
