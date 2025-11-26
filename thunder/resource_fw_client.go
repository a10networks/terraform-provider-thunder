package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwClient() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_client`: Configure the inside list\n\n__PLACEHOLDER__",
		CreateContext: resourceFwClientCreate,
		UpdateContext: resourceFwClientUpdate,
		ReadContext:   resourceFwClientRead,
		DeleteContext: resourceFwClientDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'ipv4': Make class-list type IPv4; 'ipv6': Make class-list type IPv6;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwClientCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClientCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClient(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwClientRead(ctx, d, meta)
	}
	return diags
}

func resourceFwClientUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClientUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClient(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwClientRead(ctx, d, meta)
	}
	return diags
}
func resourceFwClientDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClientDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClient(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwClientRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwClientRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwClient(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwClient(d *schema.ResourceData) edpt.FwClient {
	var ret edpt.FwClient
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	return ret
}
