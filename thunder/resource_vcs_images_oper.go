package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsImagesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vcs_images_oper`: Operational Status for the object images\n\n__PLACEHOLDER__",
		ReadContext: resourceVcsImagesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"images_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"images_name": {
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

func resourceVcsImagesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsImagesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsImagesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VcsImagesOperOper := setObjectVcsImagesOperOper(res)
		d.Set("oper", VcsImagesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVcsImagesOperOper(ret edpt.DataVcsImagesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"images_list": setSliceVcsImagesOperOperImagesList(ret.DtVcsImagesOper.Oper.ImagesList),
		},
	}
}

func setSliceVcsImagesOperOperImagesList(d []edpt.VcsImagesOperOperImagesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["images_name"] = item.ImagesName
		result = append(result, in)
	}
	return result
}

func getObjectVcsImagesOperOper(d []interface{}) edpt.VcsImagesOperOper {

	count1 := len(d)
	var ret edpt.VcsImagesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ImagesList = getSliceVcsImagesOperOperImagesList(in["images_list"].([]interface{}))
	}
	return ret
}

func getSliceVcsImagesOperOperImagesList(d []interface{}) []edpt.VcsImagesOperOperImagesList {

	count1 := len(d)
	ret := make([]edpt.VcsImagesOperOperImagesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsImagesOperOperImagesList
		oi.Type = in["type"].(string)
		oi.ImagesName = in["images_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsImagesOper(d *schema.ResourceData) edpt.VcsImagesOper {
	var ret edpt.VcsImagesOper

	ret.Oper = getObjectVcsImagesOperOper(d.Get("oper").([]interface{}))
	return ret
}
