package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sessions`: \n\n__PLACEHOLDER__",
		CreateContext: resourceSessionsCreate,
		UpdateContext: resourceSessionsUpdate,
		ReadContext:   resourceSessionsRead,
		DeleteContext: resourceSessionsDelete,

		Schema: map[string]*schema.Schema{
			"ext": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"smp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"smp_table": {
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
func resourceSessionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSessionsRead(ctx, d, meta)
	}
	return diags
}

func resourceSessionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSessionsRead(ctx, d, meta)
	}
	return diags
}
func resourceSessionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSessionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSessionsExt1397(d []interface{}) edpt.SessionsExt1397 {

	var ret edpt.SessionsExt1397
	return ret
}

func getObjectSessionsSmp1398(d []interface{}) edpt.SessionsSmp1398 {

	var ret edpt.SessionsSmp1398
	return ret
}

func getObjectSessionsSmpTable1399(d []interface{}) edpt.SessionsSmpTable1399 {

	var ret edpt.SessionsSmpTable1399
	return ret
}

func dataToEndpointSessions(d *schema.ResourceData) edpt.Sessions {
	var ret edpt.Sessions
	ret.Inst.Ext = getObjectSessionsExt1397(d.Get("ext").([]interface{}))
	ret.Inst.Smp = getObjectSessionsSmp1398(d.Get("smp").([]interface{}))
	ret.Inst.SmpTable = getObjectSessionsSmpTable1399(d.Get("smp_table").([]interface{}))
	//omit uuid
	return ret
}
