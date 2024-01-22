package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileXmlSchemaOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_xml_schema_oper`: Operational Status for the object xml-schema\n\n__PLACEHOLDER__",
		CreateContext: resourceFileXmlSchemaOperCreate,
		UpdateContext: resourceFileXmlSchemaOperUpdate,
		ReadContext:   resourceFileXmlSchemaOperRead,
		DeleteContext: resourceFileXmlSchemaOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"syntax": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"template": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func resourceFileXmlSchemaOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileXmlSchemaOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileXmlSchemaOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileXmlSchemaOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileXmlSchemaOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileXmlSchemaOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileXmlSchemaOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileXmlSchemaOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileXmlSchemaOperOper(d []interface{}) edpt.FileXmlSchemaOperOper {

	count1 := len(d)
	var ret edpt.FileXmlSchemaOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileXmlSchemaOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileXmlSchemaOperOperFileList(d []interface{}) []edpt.FileXmlSchemaOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileXmlSchemaOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileXmlSchemaOperOperFileList
		oi.File = in["file"].(string)
		oi.Syntax = in["syntax"].(string)
		oi.Template = in["template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileXmlSchemaOper(d *schema.ResourceData) edpt.FileXmlSchemaOper {
	var ret edpt.FileXmlSchemaOper
	ret.Inst.Oper = getObjectFileXmlSchemaOperOper(d.Get("oper").([]interface{}))
	return ret
}
