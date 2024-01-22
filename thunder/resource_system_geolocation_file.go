package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocationFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_geolocation_file`: Geolocation File\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGeolocationFileCreate,
		UpdateContext: resourceSystemGeolocationFileUpdate,
		ReadContext:   resourceSystemGeolocationFileRead,
		DeleteContext: resourceSystemGeolocationFileDelete,

		Schema: map[string]*schema.Schema{
			"error_info": {
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
func resourceSystemGeolocationFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocationFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocationFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocationFileRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGeolocationFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocationFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocationFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeolocationFileRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGeolocationFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocationFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocationFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGeolocationFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocationFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocationFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemGeolocationFileErrorInfo1573(d []interface{}) edpt.SystemGeolocationFileErrorInfo1573 {

	var ret edpt.SystemGeolocationFileErrorInfo1573
	return ret
}

func dataToEndpointSystemGeolocationFile(d *schema.ResourceData) edpt.SystemGeolocationFile {
	var ret edpt.SystemGeolocationFile
	ret.Inst.ErrorInfo = getObjectSystemGeolocationFileErrorInfo1573(d.Get("error_info").([]interface{}))
	//omit uuid
	return ret
}
