package thunder

import (
    "context"
    "util"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
)

func resourceIpv6RouteRib() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceIpv6RouteRibCreate,
        UpdateContext: resourceIpv6RouteRibUpdate,
        ReadContext: resourceIpv6RouteRibRead,
        DeleteContext: resourceIpv6RouteRibDelete,
        Schema: map[string]*schema.Schema{
            "ipv6_address": {
                Type: schema.TypeString, Required: true, ForceNew: true, Description: "IPV6 address",
                ValidateFunc: axapi.IsIpv6AddressPlen,
            },
            "ipv6_nexthop_ipv6": {
                Type: schema.TypeList, Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ipv6_nexthop": {
                            Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
                            ValidateFunc: validation.IsIPv6Address,
                        },
                        "ethernet": {
                            Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Ethernet interface number)",
                            ValidateFunc: validation.IntAtLeast(1),
                        },
                        "ve": {
                            Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface (Virtual Ethernet interface number)",
                        },
                        "trunk": {
                            Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
                        },
                        "distance": {
                            Type: schema.TypeInt, Optional: true, Default: 1, Description: "Distance value for this route",
                            ValidateFunc: validation.IntBetween(1, 255),
                        },
                        "description": {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
                        },
                    },
                },
            },
            "ipv6_nexthop_tunnel": {
                Type: schema.TypeList, Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "tunnel": {
                            Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
                            ValidateFunc: validation.IntBetween(1, 128),
                        },
                        "ipv6_nexthop_tunnel_addr": {
                            Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
                            ValidateFunc: validation.IsIPv6Address,
                        },
                        "distance_nexthop_tunnel": {
                            Type: schema.TypeInt, Optional: true, Default: 1, Description: "Distance value for this route",
                            ValidateFunc: validation.IntBetween(1, 255),
                        },
                        "description":  {
                            Type: schema.TypeString, Optional: true, Description: "Description for static route",
                            ValidateFunc: validation.StringLenBetween(1, 63),
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

func resourceIpv6RouteRibCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteRibCreate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpv6RouteRib(d)
        d.SetId(obj.GetId())
        err := PostEndpointIpv6RouteRib(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpv6RouteRibRead(ctx, d, meta)
    }
    return diags
}

func resourceIpv6RouteRibRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteRibRead()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        _,err := GetEndpointIpv6RouteRib(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func resourceIpv6RouteRibUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteRibUpdate()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointIpv6RouteRib(d)
        err := PutEndpointIpv6RouteRib(client.Token, obj, client.Host)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceIpv6RouteRibRead(ctx, d, meta)
    }
    return diags
}

func resourceIpv6RouteRibDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    logger := util.GetLoggerInstance()
    logger.Println("resourceIpv6RouteRibDelete()")
    client := meta.(Thunder)
    var diags diag.Diagnostics
    if client.Host != "" {
        err := DeleteEndpointIpv6RouteRib(client.Token, client.Host, d.Id())
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}

func dataToEndpointIpv6RouteRib(d *schema.ResourceData) EndpointIpv6RouteRib{
    var ret EndpointIpv6RouteRib
    ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)

    count := d.Get("ipv6_nexthop_ipv6.#").(int)
    for i := 0; i < count; i++ {
        prefix := fmt.Sprintf("ipv6_nexthop_ipv6.%d.", i)
        item := Ipv6RouteRibIpv6NexthopIpv6{}
        item.Ipv6Nexthop = d.Get(prefix + "ipv6_nexthop").(string)
        item.Ethernet = d.Get(prefix + "ethernet").(int)
        item.Ve = d.Get(prefix + "ve").(int)
        item.Trunk = d.Get(prefix + "trunk").(int)
        item.Distance = d.Get(prefix + "distance").(int)
        item.Description = d.Get(prefix + "description").(string)
        if (Ipv6RouteRibIpv6NexthopIpv6{} != item) {
            ret.Inst.Ipv6NexthopIpv6 = append(ret.Inst.Ipv6NexthopIpv6, item)
        }
    }

    count = d.Get("ipv6_nexthop_tunnel.#").(int)
    for i := 0; i < count; i++ {
        prefix := fmt.Sprintf("ipv6_nexthop_tunnel.%d.", i)
        item := Ipv6RouteRibIpv6NexthopTunnel{}
        item.Tunnel = d.Get(prefix + "tunnel").(int)
        item.Ipv6NexthopTunnelAddr = d.Get(prefix + "ipv6_nexthop_tunnel_addr").(string)
        item.DistanceNexthopTunnel = d.Get(prefix + "distance_nexthop_tunnel").(int)
        item.Description = d.Get(prefix + "description").(string)
        if (Ipv6RouteRibIpv6NexthopTunnel{} != item) {
            ret.Inst.Ipv6NexthopTunnel = append(ret.Inst.Ipv6NexthopTunnel, item)
        }
    }
    //omit uuid
    return ret
}
