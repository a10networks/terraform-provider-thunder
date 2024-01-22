package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLinkStartupConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_link_startup_config`: Startup Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceLinkStartupConfigCreate,
		UpdateContext: resourceLinkStartupConfigUpdate,
		ReadContext:   resourceLinkStartupConfigRead,
		DeleteContext: resourceLinkStartupConfigDelete,

		Schema: map[string]*schema.Schema{
			"all_partitions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All-partitions",
			},
			"default": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default startup-config",
			},
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Local Configuration Profile Name",
			},
			"primary": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create link in primary partition",
			},
			"secondary": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create link in secondary partition",
			},
		},
	}
}
func resourceLinkStartupConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLinkStartupConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLinkStartupConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLinkStartupConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceLinkStartupConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLinkStartupConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLinkStartupConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLinkStartupConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceLinkStartupConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLinkStartupConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLinkStartupConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLinkStartupConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLinkStartupConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLinkStartupConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLinkStartupConfig(d *schema.ResourceData) edpt.LinkStartupConfig {
	var ret edpt.LinkStartupConfig
	ret.Inst.AllPartitions = d.Get("all_partitions").(int)
	ret.Inst.Default = d.Get("default").(int)
	ret.Inst.FileName = d.Get("file_name").(string)
	ret.Inst.Primary = d.Get("primary").(int)
	ret.Inst.Secondary = d.Get("secondary").(int)
	return ret
}
