package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityZbar() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_zbar`: Configure zbar\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityZbarCreate,
		UpdateContext: resourceVisibilityZbarUpdate,
		ReadContext:   resourceVisibilityZbarRead,
		DeleteContext: resourceVisibilityZbarDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable zbar infra; 'disable': Disable zbar infra(default);",
			},
			"dest": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"bad_sources": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"truples": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityZbarCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbar(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityZbarRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityZbarUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbar(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityZbarRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityZbarDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbar(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityZbarRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbar(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityZbarDest3138(d []interface{}) edpt.VisibilityZbarDest3138 {

	count1 := len(d)
	var ret edpt.VisibilityZbarDest3138
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.BadSources = getObjectVisibilityZbarDestBadSources3139(in["bad_sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityZbarDestBadSources3139(d []interface{}) edpt.VisibilityZbarDestBadSources3139 {

	var ret edpt.VisibilityZbarDestBadSources3139
	return ret
}

func getObjectVisibilityZbarTruples3140(d []interface{}) edpt.VisibilityZbarTruples3140 {

	var ret edpt.VisibilityZbarTruples3140
	return ret
}

func dataToEndpointVisibilityZbar(d *schema.ResourceData) edpt.VisibilityZbar {
	var ret edpt.VisibilityZbar
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Dest = getObjectVisibilityZbarDest3138(d.Get("dest").([]interface{}))
	ret.Inst.Truples = getObjectVisibilityZbarTruples3140(d.Get("truples").([]interface{}))
	//omit uuid
	return ret
}
