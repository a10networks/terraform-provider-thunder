package thunder

import (
    "context"
    "util"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpv6RouteStaticBfdVe() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceIpv6RouteStaticBfdVeCreate,
        UpdateContext: resourceIpv6RouteStaticBfdVeUpdate,
        ReadContext: resourceIpv6RouteStaticBfdVeRead,
        DeleteContext: resourceIpv6RouteStaticBfdVeDelete,
        Schema: map[string]*schema.Schema{
            "ve_num": {
                Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Virtual ethernet interface",
                ValidateFunc: validation.IntBetween(2, 4094),
            },
            "nexthop_ipv6_ll": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "Nexthop IPv6 address (Link_local)",
                ValidateFunc: validation.IsIPv6Address,
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
            "uuid": {
                Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
            },
        },
    }
}

func resourceIpv6RouteStaticBfdVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdVeCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpv6RouteStaticBfdVe(d)
        d.SetId(obj.GetId())
        err := PostEndpointIpv6RouteStaticBfdVe(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpv6RouteStaticBfdVeRead(ctx, d, meta)
    }
    return diags
}

func resourceIpv6RouteStaticBfdVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdVeRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointIpv6RouteStaticBfdVe(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceIpv6RouteStaticBfdVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdVeUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpv6RouteStaticBfdVe(d)
        err := PutEndpointIpv6RouteStaticBfdVe(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpv6RouteStaticBfdVeRead(ctx, d, meta)
    }
    return diags
}

func resourceIpv6RouteStaticBfdVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteStaticBfdVeDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointIpv6RouteStaticBfdVe(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointIpv6RouteStaticBfdVe(d *schema.ResourceData) EndpointIpv6RouteStaticBfdVe{
    var ret EndpointIpv6RouteStaticBfdVe
    ret.Inst.VeNum = d.Get("ve_num").(int)
    ret.Inst.NexthopIpv6Ll = d.Get("nexthop_ipv6_ll").(string)
    ret.Inst.Template = d.Get("template").(string)
    ret.Inst.Threshold = d.Get("threshold").(int)
    ret.Inst.Action = d.Get("action").(string)
    //omit uuid
    return ret
}
