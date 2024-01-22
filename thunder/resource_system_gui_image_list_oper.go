package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGuiImageListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_gui_image_list_oper`: Operational Status for the object gui-image-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGuiImageListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pre_pri_gui": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pre_sec_gui": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"gui_list_pri": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gui_image": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"path": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"gui_list_sec": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gui_image": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"path": {
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

func resourceSystemGuiImageListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGuiImageListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGuiImageListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGuiImageListOperOper := setObjectSystemGuiImageListOperOper(res)
		d.Set("oper", SystemGuiImageListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGuiImageListOperOper(ret edpt.DataSystemGuiImageListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pre_pri_gui":  ret.DtSystemGuiImageListOper.Oper.PrePriGui,
			"pre_sec_gui":  ret.DtSystemGuiImageListOper.Oper.PreSecGui,
			"gui_list_pri": setSliceSystemGuiImageListOperOperGuiListPri(ret.DtSystemGuiImageListOper.Oper.GuiListPri),
			"gui_list_sec": setSliceSystemGuiImageListOperOperGuiListSec(ret.DtSystemGuiImageListOper.Oper.GuiListSec),
		},
	}
}

func setSliceSystemGuiImageListOperOperGuiListPri(d []edpt.SystemGuiImageListOperOperGuiListPri) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["gui_image"] = item.GuiImage
		in["path"] = item.Path
		result = append(result, in)
	}
	return result
}

func setSliceSystemGuiImageListOperOperGuiListSec(d []edpt.SystemGuiImageListOperOperGuiListSec) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["gui_image"] = item.GuiImage
		in["path"] = item.Path
		result = append(result, in)
	}
	return result
}

func getObjectSystemGuiImageListOperOper(d []interface{}) edpt.SystemGuiImageListOperOper {

	count1 := len(d)
	var ret edpt.SystemGuiImageListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrePriGui = in["pre_pri_gui"].(string)
		ret.PreSecGui = in["pre_sec_gui"].(string)
		ret.GuiListPri = getSliceSystemGuiImageListOperOperGuiListPri(in["gui_list_pri"].([]interface{}))
		ret.GuiListSec = getSliceSystemGuiImageListOperOperGuiListSec(in["gui_list_sec"].([]interface{}))
	}
	return ret
}

func getSliceSystemGuiImageListOperOperGuiListPri(d []interface{}) []edpt.SystemGuiImageListOperOperGuiListPri {

	count1 := len(d)
	ret := make([]edpt.SystemGuiImageListOperOperGuiListPri, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGuiImageListOperOperGuiListPri
		oi.GuiImage = in["gui_image"].(string)
		oi.Path = in["path"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGuiImageListOperOperGuiListSec(d []interface{}) []edpt.SystemGuiImageListOperOperGuiListSec {

	count1 := len(d)
	ret := make([]edpt.SystemGuiImageListOperOperGuiListSec, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGuiImageListOperOperGuiListSec
		oi.GuiImage = in["gui_image"].(string)
		oi.Path = in["path"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGuiImageListOper(d *schema.ResourceData) edpt.SystemGuiImageListOper {
	var ret edpt.SystemGuiImageListOper

	ret.Oper = getObjectSystemGuiImageListOperOper(d.Get("oper").([]interface{}))
	return ret
}
