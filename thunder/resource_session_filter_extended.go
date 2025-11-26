package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessionFilterExtended() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_session_filter_extended`: Create a extended convenience Filter to display/clear sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceSessionFilterExtendedCreate,
		UpdateContext: resourceSessionFilterExtendedUpdate,
		ReadContext:   resourceSessionFilterExtendedRead,
		DeleteContext: resourceSessionFilterExtendedDelete,

		Schema: map[string]*schema.Schema{
			"filter_rel": {
				Type: schema.TypeString, Optional: true, Default: "AND", Description: "'AND': both forward and reverse tuple-filter match; 'OR': either forward or reverse tuple-filter match;",
			},
			"fwd_filter": {
				Type: schema.TypeString, Optional: true, Description: "Specify forward tuple-filter name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Session filter name",
			},
			"rev_filter": {
				Type: schema.TypeString, Optional: true, Description: "Specify reverse tuple-filter name",
			},
			"session_type": {
				Type: schema.TypeString, Optional: true, Description: "'sip': SIP sessions;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSessionFilterExtendedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterExtendedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilterExtended(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSessionFilterExtendedRead(ctx, d, meta)
	}
	return diags
}

func resourceSessionFilterExtendedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterExtendedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilterExtended(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSessionFilterExtendedRead(ctx, d, meta)
	}
	return diags
}
func resourceSessionFilterExtendedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterExtendedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilterExtended(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSessionFilterExtendedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterExtendedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilterExtended(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSessionFilterExtended(d *schema.ResourceData) edpt.SessionFilterExtended {
	var ret edpt.SessionFilterExtended
	ret.Inst.FilterRel = d.Get("filter_rel").(string)
	ret.Inst.FwdFilter = d.Get("fwd_filter").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RevFilter = d.Get("rev_filter").(string)
	ret.Inst.SessionType = d.Get("session_type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
