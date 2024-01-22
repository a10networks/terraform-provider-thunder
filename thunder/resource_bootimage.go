package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBootimage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_bootimage`: Set next Boot Image\n\n__PLACEHOLDER__",
		CreateContext: resourceBootimageCreate,
		UpdateContext: resourceBootimageUpdate,
		ReadContext:   resourceBootimageRead,
		DeleteContext: resourceBootimageDelete,

		Schema: map[string]*schema.Schema{
			"cf_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compact flash",
						},
						"cf_pri": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Primary image",
						},
					},
				},
			},
			"hd_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Hard disk",
						},
						"pri": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Primary image",
						},
						"sec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Secondary image",
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
func resourceBootimageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootimageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootimage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBootimageRead(ctx, d, meta)
	}
	return diags
}

func resourceBootimageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootimageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootimage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBootimageRead(ctx, d, meta)
	}
	return diags
}
func resourceBootimageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootimageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootimage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBootimageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBootimageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBootimage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectBootimageCfCfg(d []interface{}) edpt.BootimageCfCfg {

	count1 := len(d)
	var ret edpt.BootimageCfCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cf = in["cf"].(int)
		ret.CfPri = in["cf_pri"].(int)
	}
	return ret
}

func getObjectBootimageHdCfg(d []interface{}) edpt.BootimageHdCfg {

	count1 := len(d)
	var ret edpt.BootimageHdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hd = in["hd"].(int)
		ret.Pri = in["pri"].(int)
		ret.Sec = in["sec"].(int)
	}
	return ret
}

func dataToEndpointBootimage(d *schema.ResourceData) edpt.Bootimage {
	var ret edpt.Bootimage
	ret.Inst.CfCfg = getObjectBootimageCfCfg(d.Get("cf_cfg").([]interface{}))
	ret.Inst.HdCfg = getObjectBootimageHdCfg(d.Get("hd_cfg").([]interface{}))
	//omit uuid
	return ret
}
