package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAsPath() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_as_path`: Configure an AS-path list for BGP\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAsPathCreate,
		UpdateContext: resourceIpAsPathUpdate,
		ReadContext:   resourceIpAsPathRead,
		DeleteContext: resourceIpAsPathDelete,

		Schema: map[string]*schema.Schema{
			"access_list": {
				Type: schema.TypeString, Required: true, Description: "Specify an access list name (Regular expression access-list name)",
			},
			"action": {
				Type: schema.TypeString, Required: true, Description: "'deny': Specify packets to reject; 'permit': Specify packets to forward;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Required: true, Description: "A regular-expression to match the BGP AS paths",
			},
		},
	}
}
func resourceIpAsPathCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAsPathCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAsPath(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAsPathRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAsPathUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAsPathUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAsPath(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAsPathRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAsPathDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAsPathDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAsPath(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAsPathRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAsPathRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAsPath(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpAsPath(d *schema.ResourceData) edpt.IpAsPath {
	var ret edpt.IpAsPath
	ret.Inst.AccessList = d.Get("access_list").(string)
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	return ret
}
