package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteThreatIntel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_threat_intel`: Delete threat-intel files\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteThreatIntelCreate,
		UpdateContext: resourceDeleteThreatIntelUpdate,
		ReadContext:   resourceDeleteThreatIntelRead,
		DeleteContext: resourceDeleteThreatIntelDelete,

		Schema: map[string]*schema.Schema{
			"database": {
				Type: schema.TypeString, Optional: true, Description: "'webroot': Delete webroot module database;",
			},
		},
	}
}
func resourceDeleteThreatIntelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteThreatIntelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteThreatIntel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteThreatIntelRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteThreatIntelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteThreatIntelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteThreatIntel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteThreatIntelRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteThreatIntelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteThreatIntelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteThreatIntel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteThreatIntelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteThreatIntelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteThreatIntel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteThreatIntel(d *schema.ResourceData) edpt.DeleteThreatIntel {
	var ret edpt.DeleteThreatIntel
	ret.Inst.Database = d.Get("database").(string)
	return ret
}
