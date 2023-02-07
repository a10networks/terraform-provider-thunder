package thunder

//Thunder resource RouterBgpNeighborVeNeighbor

import (
	"context"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborVeNeighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpNeighborVeNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborVeNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborVeNeighborRead,
		DeleteContext: resourceRouterBgpNeighborVeNeighborDelete,
		Schema: map[string]*schema.Schema{
			"ve": {
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
		},
	}
}

func resourceRouterBgpNeighborVeNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborVeNeighbor (Inside resourceRouterBgpNeighborVeNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("ve").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborVeNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborVeNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		err := go_thunder.PostRouterBgpNeighborVeNeighbor(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborVeNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborVeNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpNeighborVeNeighbor (Inside resourceRouterBgpNeighborVeNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborVeNeighbor(client.Token, name1, name2, client.Host)
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

func resourceRouterBgpNeighborVeNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborVeNeighbor   (Inside resourceRouterBgpNeighborVeNeighborUpdate) ")
		data := dataToRouterBgpNeighborVeNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborVeNeighbor ")
		err := go_thunder.PutRouterBgpNeighborVeNeighbor(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborVeNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborVeNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborVeNeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborVeNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouterBgpNeighborVeNeighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborVeNeighbor {
	var vc go_thunder.RouterBgpNeighborVeNeighbor
	var c go_thunder.RouterBgpNeighborVeNeighborInstance
	c.Ve = d.Get("ve").(int)
	c.Unnumbered = d.Get("unnumbered").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ve = c
	return vc
}
