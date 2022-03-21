package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpRouteStaticBfd() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceIpRouteStaticBfdCreate,
        UpdateContext: resourceIpRouteStaticBfdUpdate,
        ReadContext: resourceIpRouteStaticBfdRead,
        DeleteContext: resourceIpRouteStaticBfdDelete,
        Schema: map[string]*schema.Schema{
            "local_ip": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Local IP address",
                ValidateFunc: validation.IsIPv4Address,
            },
            "nexthop_ip": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Nexthop IP address",
                ValidateFunc: validation.IsIPv4Address,
            },
            "template": {
                Type: schema.TypeString, Optional: true, Description: "Configure tracking template (bind tracking template name)",
                ValidateFunc: validation.StringLenBetween(1, 63),
            },
            "threshold": {
                Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
                ValidateFunc: validation.IntBetween(1, 255),
            },
            "action": {
                Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
                ValidateFunc: validation.StringInSlice([]string{"down"}, false),
            },
            "uuid": {Type: schema.TypeString, Computed: true, Description: "uuid of the object"},
        },
    }
}

func resourceIpRouteStaticBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteStaticBfdCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpRouteStaticBfd(d)
        d.SetId(obj.GetId())
        err := PostEndpointIpRouteStaticBfd(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpRouteStaticBfdRead(ctx, d, meta)
    }
    return diags
}

func resourceIpRouteStaticBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteStaticBfdRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointIpRouteStaticBfd(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceIpRouteStaticBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteStaticBfdCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpRouteStaticBfd(d)
        err := PutEndpointIpRouteStaticBfd(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpRouteStaticBfdRead(ctx, d, meta)
    }
    return diags
}

func resourceIpRouteStaticBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpRouteStaticBfdRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointIpRouteStaticBfd(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointIpRouteStaticBfd(d *schema.ResourceData) EndpointIpRouteStaticBfd{
    var ret EndpointIpRouteStaticBfd
    ret.Inst.Action = d.Get("action").(string)
    ret.Inst.LocalIp = d.Get("local_ip").(string)
    ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
    ret.Inst.Template = d.Get("template").(string)
    ret.Inst.Threshold = d.Get("threshold").(int)
    //omit uuid
    return ret
}
