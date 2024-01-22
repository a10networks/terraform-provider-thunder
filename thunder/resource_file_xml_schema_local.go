package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileXmlSchemaLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_xml_schema_local`: XML-Schema file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileXmlSchemaLocalCreate,
		UpdateContext: resourceFileXmlSchemaLocalUpdate,
		ReadContext:   resourceFileXmlSchemaLocalRead,
		DeleteContext: resourceFileXmlSchemaLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
			},
			"dst_file": {
				Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "XML-schema local file name",
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
func resourceFileXmlSchemaLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileXmlSchemaLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileXmlSchemaLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileXmlSchemaLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileXmlSchemaLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileXmlSchemaLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileXmlSchemaLocal(d *schema.ResourceData) edpt.FileXmlSchemaLocal {
	var ret edpt.FileXmlSchemaLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
	return ret
}
