data "solana_slot_leaders" "leaders" {
  start_slot = 12345
  limit = 5
}
