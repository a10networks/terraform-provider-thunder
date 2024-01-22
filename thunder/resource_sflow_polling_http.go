package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPollingHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling_http`: Poll HTTP counters\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingHttpCreate,
		UpdateContext: resourceSflowPollingHttpUpdate,
		ReadContext:   resourceSflowPollingHttpRead,
		DeleteContext: resourceSflowPollingHttpDelete,

		Schema: map[string]*schema.Schema{
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling HTTP counters; 'disable': Disable polling HTTP counters;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowPollingHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowPollingHttp(d *schema.ResourceData) edpt.SflowPollingHttp {
	var ret edpt.SflowPollingHttp
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
