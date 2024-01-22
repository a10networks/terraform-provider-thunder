package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileTemplateLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_template_local`: upgrade acos wizard template\n\n__PLACEHOLDER__",
		CreateContext: resourceFileTemplateLocalCreate,
		UpdateContext: resourceFileTemplateLocalUpdate,
		ReadContext:   resourceFileTemplateLocalRead,
		DeleteContext: resourceFileTemplateLocalDelete,

		Schema: map[string]*schema.Schema{
			"act_type": {
				Type: schema.TypeString, Optional: true, Description: "Specify ACT package type (i.e. slug/destination name)",
			},
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
			},
			"app": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"act_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"dst_file": {
				Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "aflex local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileTemplateLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTemplateLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileTemplateLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTemplateLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileTemplateLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileTemplateLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileTemplateApp352(d []interface{}) edpt.FileTemplateApp352 {

	count1 := len(d)
	var ret edpt.FileTemplateApp352
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.ActName = in["act_name"].(string)
		ret.Version = in["version"].(string)
	}
	return ret
}

func dataToEndpointFileTemplateLocal(d *schema.ResourceData) edpt.FileTemplateLocal {
	var ret edpt.FileTemplateLocal
	ret.Inst.ActType = d.Get("act_type").(string)
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.App = getObjectFileTemplateApp352(d.Get("app").([]interface{}))
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
	return ret
}
