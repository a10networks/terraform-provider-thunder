package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPollingA10Proprietary() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling_a10_proprietary`: A10 proprietary counters\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingA10ProprietaryCreate,
		UpdateContext: resourceSflowPollingA10ProprietaryUpdate,
		ReadContext:   resourceSflowPollingA10ProprietaryRead,
		DeleteContext: resourceSflowPollingA10ProprietaryDelete,

		Schema: map[string]*schema.Schema{
			"export_deprecated_counters": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export deprecated counters",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowPollingA10ProprietaryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingA10ProprietaryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingA10Proprietary(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingA10ProprietaryRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingA10ProprietaryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingA10ProprietaryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingA10Proprietary(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingA10ProprietaryRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingA10ProprietaryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingA10ProprietaryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingA10Proprietary(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingA10ProprietaryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingA10ProprietaryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingA10Proprietary(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowPollingA10Proprietary(d *schema.ResourceData) edpt.SflowPollingA10Proprietary {
	var ret edpt.SflowPollingA10Proprietary
	ret.Inst.ExportDeprecatedCounters = d.Get("export_deprecated_counters").(int)
	//omit uuid
	return ret
}
