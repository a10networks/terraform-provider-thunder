package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileTechsupportLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_techsupport_local`: Generate the current showtech file\n\n__PLACEHOLDER__",
		CreateContext: resourceFileTechsupportLocalCreate,
		UpdateContext: resourceFileTechsupportLocalUpdate,
		ReadContext:   resourceFileTechsupportLocalRead,
		DeleteContext: resourceFileTechsupportLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'check': check; 'delete': delete; 'export': export;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "techsupport local file name",
			},
			"slot": {
				Type: schema.TypeInt, Optional: true, Description: "specify slot id in chassis",
			},
			"status": {
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
func resourceFileTechsupportLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTechsupportLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileTechsupportLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTechsupportLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileTechsupportLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileTechsupportLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileTechsupportStatus351(d []interface{}) edpt.FileTechsupportStatus351 {

	var ret edpt.FileTechsupportStatus351
	return ret
}

func dataToEndpointFileTechsupportLocal(d *schema.ResourceData) edpt.FileTechsupportLocal {
	var ret edpt.FileTechsupportLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.Slot = d.Get("slot").(int)
	ret.Inst.Status = getObjectFileTechsupportStatus351(d.Get("status").([]interface{}))
	//omit uuid
	return ret
}
