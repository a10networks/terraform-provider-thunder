provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authorization_policy_jwt_claim_map" "thunder_aam_authorization_policy_jwt_claim_map" {

  name         = "test"
  attr_num     = 8
  boolean_type = 0
  claim        = "47"
  type         = 0

}