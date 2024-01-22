provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_rule_oper" "thunder_rule_set_rule_oper" {
  oper {
    hitcount           = hitcount
    permitbytes        = permitbytes
    denybytes          = denybytes
    resetbytes         = resetbytes
    totalbytes         = totalbytes
    permitpackets      = permitpackets
    denypackets        = denypackets
    resetpackets       = resetpackets
    totalpackets       = totalpackets
    activesessiontcp   = activesessiontcp
    activesessionudp   = activesessionudp
    activesessionicmp  = activesessionicmp
    activesessionsctp  = activesessionsctp
    activesessionother = activesessionother
    activesessiontotal = activesessiontotal
    sessiontcp         = sessiontcp
    sessionudp         = sessionudp
    sessionicmp        = sessionicmp
    sessionsctp        = sessionsctp
    sessionother       = sessionother
    sessiontotal       = sessiontotal
    ratelimitdrops     = ratelimitdrops
  }

}
output "get_rule_set_rule_oper" {
  value = ["${data.thunder_rule_set_rule_oper.thunder_rule_set_rule_oper}"]
}