package minecraft_guardrails_opa

# METADATA
# title: Recommend droplet backups
# description: Minecraft droplets should have backups enabled.
warn_missing_backups[msg] {
    input.type == "digitalocean:index/droplet:Droplet"
    not input.backups
    msg := sprintf("Droplet '%s' has no backups enabled; consider turning them on.", [input.__name])
}
