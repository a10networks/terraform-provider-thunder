





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceDdosZoneTemplateIps() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_ddos_zone_template_ips`: IPS template configuration\n\n__PLACEHOLDER__",
        CreateContext: resourceDdosZoneTemplateIpsCreate,
        UpdateContext: resourceDdosZoneTemplateIpsUpdate,
        ReadContext: resourceDdosZoneTemplateIpsRead,
        DeleteContext: resourceDdosZoneTemplateIpsDelete,

        Schema: map[string]*schema.Schema{
        "high_serverity_action_list_name": {
        Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for high serverity signature",
        },
        "ips_profile_list": {
        Type: schema.TypeList, Optional: true, Description: "",
        Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
        "ips_profile_name": {
        Type: schema.TypeString, Optional: true, Description: "IPS Profile Name",
        },
        },
},
        },
        "ips_tmpl_name": {
        Type: schema.TypeString, Required: true, Description: "DDOS IPS Template Name",
        },
        "low_serverity_action_list_name": {
        Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for low serverity signature",
        },
        "medium_serverity_action_list_name": {
        Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for medium serverity signature",
        },
        "streaming_scan_disable": {
        Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable IPS streaming scan",
        },
        "user_tag": {
        Type: schema.TypeString, Optional: true, Description: "Customized tag",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceDdosZoneTemplateIpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDdosZoneTemplateIpsCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDdosZoneTemplateIps(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceDdosZoneTemplateIpsRead(ctx, d, meta)
    }
    return diags
}

func resourceDdosZoneTemplateIpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDdosZoneTemplateIpsUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDdosZoneTemplateIps(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceDdosZoneTemplateIpsRead(ctx, d, meta)
    }
    return diags
}
func resourceDdosZoneTemplateIpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDdosZoneTemplateIpsDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDdosZoneTemplateIps(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceDdosZoneTemplateIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceDdosZoneTemplateIpsRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointDdosZoneTemplateIps(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func getSliceDdosZoneTemplateIpsIpsProfileList(d []interface{}) []edpt.DdosZoneTemplateIpsIpsProfileList {
    
        count1 := len(d)
            ret := make([]edpt.DdosZoneTemplateIpsIpsProfileList, 0, count1)
            for _, item := range d {
                in := item.(map[string]interface{})
                var oi edpt.DdosZoneTemplateIpsIpsProfileList
            oi.IpsProfileName = in["ips_profile_name"].(string)
            ret = append(ret, oi)
        }
return ret
}



func dataToEndpointDdosZoneTemplateIps(d *schema.ResourceData) edpt.DdosZoneTemplateIps {
    var ret edpt.DdosZoneTemplateIps
	ret.Inst.HighServerityActionListName = d.Get("high_serverity_action_list_name").(string)
	ret.Inst.IpsProfileList = getSliceDdosZoneTemplateIpsIpsProfileList(d.Get("ips_profile_list").([]interface{}))
	ret.Inst.IpsTmplName = d.Get("ips_tmpl_name").(string)
	ret.Inst.LowServerityActionListName = d.Get("low_serverity_action_list_name").(string)
	ret.Inst.MediumServerityActionListName = d.Get("medium_serverity_action_list_name").(string)
	ret.Inst.StreamingScanDisable = d.Get("streaming_scan_disable").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
    return ret
}

