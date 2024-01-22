package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateAppsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_template_apps_oper`: Operational Status for the object apps\n\n__PLACEHOLDER__",
		ReadContext: resourceTemplateAppsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"installed": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category_slug": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"release_notes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"video_url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"guide_url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"developer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"icon": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"supported_acos": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"version": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"file_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"deleted": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"color": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"available": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category_slug": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"release_notes": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"video_url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"guide_url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"developer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"icon": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"supported_acos": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"version": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"file_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"deleted": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"color": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category_name": {
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

func resourceTemplateAppsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateAppsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateAppsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TemplateAppsOperOper := setObjectTemplateAppsOperOper(res)
		d.Set("oper", TemplateAppsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTemplateAppsOperOper(ret edpt.DataTemplateAppsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"installed": setSliceTemplateAppsOperOperInstalled(ret.DtTemplateAppsOper.Oper.Installed),
			"available": setSliceTemplateAppsOperOperAvailable(ret.DtTemplateAppsOper.Oper.Available),
		},
	}
}

func setSliceTemplateAppsOperOperInstalled(d []edpt.TemplateAppsOperOperInstalled) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["version"] = item.Version
		in["category_slug"] = item.Category_slug
		in["description"] = item.Description
		in["type"] = item.Type
		in["release_notes"] = item.Release_notes
		in["video_url"] = item.Video_url
		in["guide_url"] = item.Guide_url
		in["developer"] = item.Developer
		in["icon"] = item.Icon
		in["supported_acos"] = setSliceTemplateAppsOperOperInstalledSupported_acos(item.Supported_acos)
		in["file_name"] = item.File_name
		in["deleted"] = item.Deleted
		in["color"] = item.Color
		in["category_name"] = item.Category_name
		result = append(result, in)
	}
	return result
}

func setSliceTemplateAppsOperOperInstalledSupported_acos(d []edpt.TemplateAppsOperOperInstalledSupported_acos) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["version"] = item.Version
		result = append(result, in)
	}
	return result
}

func setSliceTemplateAppsOperOperAvailable(d []edpt.TemplateAppsOperOperAvailable) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["version"] = item.Version
		in["category_slug"] = item.Category_slug
		in["description"] = item.Description
		in["type"] = item.Type
		in["release_notes"] = item.Release_notes
		in["video_url"] = item.Video_url
		in["guide_url"] = item.Guide_url
		in["developer"] = item.Developer
		in["icon"] = item.Icon
		in["supported_acos"] = setSliceTemplateAppsOperOperAvailableSupported_acos(item.Supported_acos)
		in["file_name"] = item.File_name
		in["deleted"] = item.Deleted
		in["color"] = item.Color
		in["category_name"] = item.Category_name
		result = append(result, in)
	}
	return result
}

func setSliceTemplateAppsOperOperAvailableSupported_acos(d []edpt.TemplateAppsOperOperAvailableSupported_acos) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["version"] = item.Version
		result = append(result, in)
	}
	return result
}

func getObjectTemplateAppsOperOper(d []interface{}) edpt.TemplateAppsOperOper {

	count1 := len(d)
	var ret edpt.TemplateAppsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Installed = getSliceTemplateAppsOperOperInstalled(in["installed"].([]interface{}))
		ret.Available = getSliceTemplateAppsOperOperAvailable(in["available"].([]interface{}))
	}
	return ret
}

func getSliceTemplateAppsOperOperInstalled(d []interface{}) []edpt.TemplateAppsOperOperInstalled {

	count1 := len(d)
	ret := make([]edpt.TemplateAppsOperOperInstalled, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateAppsOperOperInstalled
		oi.Version = in["version"].(string)
		oi.Category_slug = in["category_slug"].(string)
		oi.Description = in["description"].(string)
		oi.Type = in["type"].(string)
		oi.Release_notes = in["release_notes"].(string)
		oi.Video_url = in["video_url"].(string)
		oi.Guide_url = in["guide_url"].(string)
		oi.Developer = in["developer"].(string)
		oi.Icon = in["icon"].(string)
		oi.Supported_acos = getSliceTemplateAppsOperOperInstalledSupported_acos(in["supported_acos"].([]interface{}))
		oi.File_name = in["file_name"].(string)
		oi.Deleted = in["deleted"].(string)
		oi.Color = in["color"].(string)
		oi.Category_name = in["category_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTemplateAppsOperOperInstalledSupported_acos(d []interface{}) []edpt.TemplateAppsOperOperInstalledSupported_acos {

	count1 := len(d)
	ret := make([]edpt.TemplateAppsOperOperInstalledSupported_acos, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateAppsOperOperInstalledSupported_acos
		oi.Version = in["version"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTemplateAppsOperOperAvailable(d []interface{}) []edpt.TemplateAppsOperOperAvailable {

	count1 := len(d)
	ret := make([]edpt.TemplateAppsOperOperAvailable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateAppsOperOperAvailable
		oi.Version = in["version"].(string)
		oi.Category_slug = in["category_slug"].(string)
		oi.Description = in["description"].(string)
		oi.Type = in["type"].(string)
		oi.Release_notes = in["release_notes"].(string)
		oi.Video_url = in["video_url"].(string)
		oi.Guide_url = in["guide_url"].(string)
		oi.Developer = in["developer"].(string)
		oi.Icon = in["icon"].(string)
		oi.Supported_acos = getSliceTemplateAppsOperOperAvailableSupported_acos(in["supported_acos"].([]interface{}))
		oi.File_name = in["file_name"].(string)
		oi.Deleted = in["deleted"].(string)
		oi.Color = in["color"].(string)
		oi.Category_name = in["category_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTemplateAppsOperOperAvailableSupported_acos(d []interface{}) []edpt.TemplateAppsOperOperAvailableSupported_acos {

	count1 := len(d)
	ret := make([]edpt.TemplateAppsOperOperAvailableSupported_acos, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateAppsOperOperAvailableSupported_acos
		oi.Version = in["version"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateAppsOper(d *schema.ResourceData) edpt.TemplateAppsOper {
	var ret edpt.TemplateAppsOper

	ret.Oper = getObjectTemplateAppsOperOper(d.Get("oper").([]interface{}))
	return ret
}
