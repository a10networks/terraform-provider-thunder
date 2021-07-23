package thunder

//Thunder resource RouterBgpNeighborTrunkNeighbor

import (
	"context"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborTrunkNeighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpNeighborTrunkNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborTrunkNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborTrunkNeighborRead,
		DeleteContext: resourceRouterBgpNeighborTrunkNeighborDelete,
		Schema: map[string]*schema.Schema{
			"trunk": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"unnumbered": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"peer_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpNeighborTrunkNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborTrunkNeighbor (Inside resourceRouterBgpNeighborTrunkNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("trunk").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborTrunkNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborTrunkNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		err := go_thunder.PostRouterBgpNeighborTrunkNeighbor(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborTrunkNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborTrunkNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpNeighborTrunkNeighbor (Inside resourceRouterBgpNeighborTrunkNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborTrunkNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceRouterBgpNeighborTrunkNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborTrunkNeighbor   (Inside resourceRouterBgpNeighborTrunkNeighborUpdate) ")
		data := dataToRouterBgpNeighborTrunkNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborTrunkNeighbor ")
		err := go_thunder.PutRouterBgpNeighborTrunkNeighbor(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborTrunkNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborTrunkNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborTrunkNeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborTrunkNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouterBgpNeighborTrunkNeighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborTrunkNeighbor {
	var vc go_thunder.RouterBgpNeighborTrunkNeighbor
	var c go_thunder.RouterBgpNeighborTrunkNeighborInstance
	c.Trunk = d.Get("trunk").(int)
	c.Unnumbered = d.Get("unnumbered").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Trunk = c
	return vc
}
