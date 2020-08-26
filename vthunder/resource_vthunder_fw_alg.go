package vthunder

//vThunder resource FwAlg

import (
    "github.com/go_vthunder/vthunder"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "util"
    "fmt"
    "strconv"
)

func resourceFwAlg() *schema.Resource {
    return &schema.Resource{
        Create: resourceFwAlgCreate,
        Update: resourceFwAlgUpdate,
        Read:   resourceFwAlgRead,
        Delete: resourceFwAlgDelete,
         Schema: map[string]*schema.Schema{ 
"ftp": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"sampling_enable": {
Type: schema.TypeList,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"counters1":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"sip": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"sampling_enable": {
Type: schema.TypeList,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"counters1":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"pptp": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"sampling_enable": {
Type: schema.TypeList,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"counters1":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"rtsp": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"sampling_enable": {
Type: schema.TypeList,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"counters1":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"dns": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"tftp": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"sampling_enable": {
Type: schema.TypeList,
Optional: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"counters1":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
"icmp": {
Type: schema.TypeList,
Optional: true,
MaxItems: 1,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
}, 
 }, 
 }, 
},
    } }
 
func resourceFwAlgCreate(d *schema.ResourceData, meta interface{}) error {
     logger := util.GetLoggerInstance()
     client := meta.(vThunder)
      
    if client.Host != "" {
         logger.Println("[INFO] Creating FwAlg (Inside resourceFwAlgCreate) ")
          
        data := dataToFwAlg(d)
         logger.Println("[INFO] received formatted data from method data to FwAlg --")
         d.SetId(strconv.Itoa('1'))
         go_vthunder.PostFwAlg(client.Token, data, client.Host)
 
        return resourceFwAlgRead(d, meta)

    }
     return nil
 }
    
func resourceFwAlgRead(d *schema.ResourceData, meta interface{}) error {
     logger := util.GetLoggerInstance()
     client := meta.(vThunder)
     logger.Println("[INFO] Reading FwAlg (Inside resourceFwAlgRead)")
 
    if client.Host != "" {
         name := d.Id()
        logger.Println("[INFO] Fetching service Read" + name)
         data, err := go_vthunder.GetFwAlg(client.Token, client.Host)
         if data == nil {
             logger.Println("[INFO] No data found " + name)
             d.SetId("")
             return nil
         }
        return err
    }
     return nil
 }
    
func resourceFwAlgUpdate(d *schema.ResourceData, meta interface{}) error {

    return resourceFwAlgRead(d, meta)
}

func resourceFwAlgDelete(d *schema.ResourceData, meta interface{}) error {

    return resourceFwAlgRead(d, meta)
}
func dataToFwAlg(d *schema.ResourceData) go_vthunder.Fw {
    var vc go_vthunder.Fw
    var c go_vthunder.FwAlgInstance
    
var obj1 go_vthunder.Ftp
prefix := "ftp.0."
obj1.DefaultPortDisable = d.Get(prefix+"default_port_disable").(string)

SamplingEnableCount := d.Get("sampling_enable.#").(int)
obj1.Counters1 = make([]go_vthunder.SamplingEnable, 0, SamplingEnableCount)

for i := 0; i < SamplingEnableCount; i++ {
var obj1 go_vthunder.SamplingEnable
prefix := fmt.Sprintf("sampling_enable.%d.",i)
obj1.Counters1 = d.Get(prefix+"counters1").(string)
obj1.Counters1 = append(obj1.Counters1, obj1)
}


c.DefaultPortDisable = obj1


var obj2 go_vthunder.Sip
prefix := "sip.0."
obj2.DefaultPortDisable = d.Get(prefix+"default_port_disable").(string)

SamplingEnableCount := d.Get("sampling_enable.#").(int)
obj2.Counters1 = make([]go_vthunder.SamplingEnable, 0, SamplingEnableCount)

for i := 0; i < SamplingEnableCount; i++ {
var obj1 go_vthunder.SamplingEnable
prefix := fmt.Sprintf("sampling_enable.%d.",i)
obj1.Counters1 = d.Get(prefix+"counters1").(string)
obj2.Counters1 = append(obj2.Counters1, obj1)
}


c.DefaultPortDisable = obj2


var obj3 go_vthunder.Pptp
prefix := "pptp.0."
obj3.DefaultPortDisable = d.Get(prefix+"default_port_disable").(string)

SamplingEnableCount := d.Get("sampling_enable.#").(int)
obj3.Counters1 = make([]go_vthunder.SamplingEnable, 0, SamplingEnableCount)

for i := 0; i < SamplingEnableCount; i++ {
var obj1 go_vthunder.SamplingEnable
prefix := fmt.Sprintf("sampling_enable.%d.",i)
obj1.Counters1 = d.Get(prefix+"counters1").(string)
obj3.Counters1 = append(obj3.Counters1, obj1)
}


c.DefaultPortDisable = obj3


var obj4 go_vthunder.Rtsp
prefix := "rtsp.0."
obj4.DefaultPortDisable = d.Get(prefix+"default_port_disable").(string)

SamplingEnableCount := d.Get("sampling_enable.#").(int)
obj4.Counters1 = make([]go_vthunder.SamplingEnable, 0, SamplingEnableCount)

for i := 0; i < SamplingEnableCount; i++ {
var obj1 go_vthunder.SamplingEnable
prefix := fmt.Sprintf("sampling_enable.%d.",i)
obj1.Counters1 = d.Get(prefix+"counters1").(string)
obj4.Counters1 = append(obj4.Counters1, obj1)
}


c.DefaultPortDisable = obj4


var obj5 go_vthunder.Dns
prefix := "dns.0."
obj5.DefaultPortDisable = d.Get(prefix+"default_port_disable").(string)

c.DefaultPortDisable = obj5


var obj6 go_vthunder.Tftp
prefix := "tftp.0."
obj6.DefaultPortDisable = d.Get(prefix+"default_port_disable").(string)

SamplingEnableCount := d.Get("sampling_enable.#").(int)
obj6.Counters1 = make([]go_vthunder.SamplingEnable, 0, SamplingEnableCount)

for i := 0; i < SamplingEnableCount; i++ {
var obj1 go_vthunder.SamplingEnable
prefix := fmt.Sprintf("sampling_enable.%d.",i)
obj1.Counters1 = d.Get(prefix+"counters1").(string)
obj6.Counters1 = append(obj6.Counters1, obj1)
}


c.DefaultPortDisable = obj6


var obj7 go_vthunder.Icmp
prefix := "icmp.0."
obj7.Disable = d.Get(prefix+"disable").(string)

c.Disable = obj7


    vc.Ftp = c 
    return vc
}