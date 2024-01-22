package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecFragmentationOption() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_fragmentation_option`: Configure Fragmentation option for a Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecFragmentationOptionCreate,
		UpdateContext: resourceFlowspecFragmentationOptionUpdate,
		ReadContext:   resourceFlowspecFragmentationOptionRead,
		DeleteContext: resourceFlowspecFragmentationOptionDelete,

		Schema: map[string]*schema.Schema{
			"frag_attribute": {
				Type: schema.TypeString, Required: true, Description: "'is-fragment': Is fragmented packet; 'first-fragment': Is the first fragment packet; 'last-fragment': Is the last fragment; 'dont-fragment': Is DF bit set;",
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
func resourceFlowspecFragmentationOptionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFragmentationOptionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFragmentationOption(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecFragmentationOptionRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecFragmentationOptionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFragmentationOptionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFragmentationOption(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecFragmentationOptionRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecFragmentationOptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFragmentationOptionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFragmentationOption(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecFragmentationOptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecFragmentationOptionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecFragmentationOption(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecFragmentationOption(d *schema.ResourceData) edpt.FlowspecFragmentationOption {
	var ret edpt.FlowspecFragmentationOption
	ret.Inst.FragAttribute = d.Get("frag_attribute").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
