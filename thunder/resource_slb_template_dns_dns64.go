package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsDns64() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_dns64`: Enable DNS64\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsDns64Create,
		UpdateContext: resourceSlbTemplateDnsDns64Update,
		ReadContext:   resourceSlbTemplateDnsDns64Read,
		DeleteContext: resourceSlbTemplateDnsDns64Delete,

		Schema: map[string]*schema.Schema{
			"cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a cached A-query response to provide AAAA query responses for the same hostname",
			},
			"change_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always change incoming AAAA DNS Query to A",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS64",
			},
			"parallel_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward AAAA Query & generate A Query in parallel",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Retry count, default is 3 (Retry Number)",
			},
			"single_response_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Single Response which is used to avoid ambiguity",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Timeout to send additional Queries, unit: second, default is 1",
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
func resourceSlbTemplateDnsDns64Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsDns64Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsDns64(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsDns64Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsDns64Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsDns64Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsDns64(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsDns64Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsDns64Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsDns64Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsDns64(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsDns64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsDns64Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsDns64(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsDns64(d *schema.ResourceData) edpt.SlbTemplateDnsDns64 {
	var ret edpt.SlbTemplateDnsDns64
	ret.Inst.Cache = d.Get("cache").(int)
	ret.Inst.ChangeQuery = d.Get("change_query").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.ParallelQuery = d.Get("parallel_query").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.SingleResponseDisable = d.Get("single_response_disable").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
