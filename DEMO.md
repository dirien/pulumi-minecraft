# Live Demo Runbook — Beyond `pulumi up`

The demo arc follows the deck: **Component → Template (IDP) → ESC → Policy**. One
running example throughout: a Java Minecraft server on a DigitalOcean droplet.

Everything below is already published/created (see [Pre-flight](#pre-flight)).
During the talk you mostly *show* and *run preview*, not wait on long deploys.

---

## Artifacts (already live in the `ediri` org)

| Thing | Reference |
|-------|-----------|
| Component (registry) | `ediri/minecraft@0.1.0` |
| Component (source, what templates consume) | `github.com/dirien/pulumi-minecraft@v0.1.0` |
| Templates | `examples/go`, `examples/csharp`, `examples/yaml` |
| ESC base env (creds + defaults) | `ediri/minecraft/base` → imports `pulumi-idp/auth` |
| ESC overlays | `ediri/minecraft/demo`, `ediri/minecraft/prod` |
| Policy: require tag (mandatory, TS) | `ediri/minecraft-guardrails-ts` |
| Policy: restrict region (mandatory, Python) | `ediri/minecraft-guardrails-py` |
| Policy: recommend backups (advisory, Rego) | `ediri/minecraft_guardrails_opa` |

---

## Pre-flight (do this the morning of the talk)

```bash
pulumi whoami -v                       # logged in, org list includes 'ediri'
esc env open ediri/pulumi-idp/auth --format json | jq '.pulumiConfig["digitalocean:token"] != null'   # true
pulumi package get-schema ediri/minecraft@0.1.0 | jq .name      # "minecraft"
pulumi env open ediri/minecraft/demo --format json | jq .pulumiConfig.motd   # demo MOTD resolves
```

Pick **one region/size you can afford** and pre-warm a server if you want a live
world to connect to (cloud-init takes 3–5 min — don't do it cold on stage):

```bash
cd examples/yaml
pulumi stack init ediri/minecraft-on-digitalocean-yaml/talk
pulumi config env add minecraft/demo
pulumi up --yes        # ~1 min to create; server reachable a few min later
pulumi stack output serverAddress      # paste into a Minecraft client beforehand
```

> Java client connect: **Multiplayer → Direct Connect → `<serverAddress>`**.

---

## Part 1 — The Component (slides: "Pulumi IDP / Golden paths")

**Say:** the platform team writes the infra once, publishes it, and never gets a ticket again.

**Show** the component source — it's real Go, not YAML:

```bash
bat minecraftserver.go        # inputs (motd, region, size, …), Droplet + Firewall
```

**Say:** one command put it in the registry for the whole org:

```bash
# (already done — show the registry page, don't re-run live)
# pulumi package publish github.com/dirien/pulumi-minecraft@0.1.0 --publisher ediri
```

Open the registry page: `https://app.pulumi.com/ediri/registry` → **minecraft**.

---

## Part 2 — The Template (slides: "no-code stack", DEMO after IDP)

**Say:** the developer never opens an editor. New Project, pick the template, fill two
fields, Create. Same component, any language.

**Option A — local `pulumi new` (reliable on stage):**

```bash
cd /tmp
pulumi new https://github.com/dirien/pulumi-minecraft/tree/main/examples/yaml \
  --name mc-live --stack ediri/mc-live/dev
# answer the prompts: motd, region, size, difficulty, gamemode…
pulumi config env add minecraft/demo      # creds from ESC (Part 3 sets this up)
pulumi up
```

**Option B — Pulumi Cloud "New Project":** show the template in the Cloud UI and click
through. (Polished, but needs network + the template registered in the org.)

**Say:** three languages, one component. Show that `examples/go`, `examples/csharp`,
and `examples/yaml` all deploy the same server — the dev picks what they read.

---

## Part 3 — ESC (slides: "Kill config and secret sprawl", DEMO after ESC)

**Say:** the template never asked for a cloud token. Auth comes from ESC, and the
config composes — a base env holds defaults, `demo` and `prod` import it and override.

**Show the composition (instant, no deploy):**

```bash
esc env get ediri/minecraft/base        # imports: [pulumi-idp/auth]; default server.properties
pulumi env open ediri/minecraft/demo  --format json | jq .pulumiConfig   # base + demo overrides
pulumi env open ediri/minecraft/prod  --format json | jq .pulumiConfig   # base + prod overrides
```

**Point at:** `motd`, `gamemode`, `difficulty`, `maxPlayers` differ between demo/prod,
but `digitalocean:token` is present in both — pulled through `base` → `pulumi-idp/auth`.

**Say:** same template, swap the env, different server:

```bash
pulumi config env add minecraft/prod --stack <stack>   # instead of demo
pulumi preview                                          # plan now reflects prod values
```

> Live deploy of a second server is optional — the `pulumi env open` diff tells the
> story instantly. Mention rotation: change a secret in `base`, every consumer gets it.

---

## Part 4 — Policy (slides: "Guardrails, because nobody should dance on razor blades")

**Say:** guardrails run on `preview` — before any resource exists. Published once,
enabled org-wide, every stack is checked.

**Enable the three packs (do this just before the section):**

```bash
pulumi policy enable ediri/minecraft-guardrails-ts latest
pulumi policy enable ediri/minecraft-guardrails-py latest
pulumi policy enable ediri/minecraft_guardrails_opa latest
```

**Show them fire on a normal preview (no `--policy-pack` flag — they're org-wide):**

```bash
cd examples/yaml
pulumi preview --stack ediri/minecraft-on-digitalocean-yaml/talk
```

Expected `Policies:` block:
```
✅ minecraft-guardrails-ts@…    require-minecraft-tag (mandatory)  — pass
✅ minecraft-guardrails-py@…    restrict-droplet-region (mandatory) — pass
⚠️ minecraft_guardrails_opa@…   warn_missing_backups (advisory)    — warns
```

**The money moment — show a mandatory block.** Set a region that isn't allowed and
preview again; the Python policy stops it cold *before* the droplet exists:

```bash
pulumi config set region lon1 --stack ediri/minecraft-on-digitalocean-yaml/talk
pulumi preview --stack ediri/minecraft-on-digitalocean-yaml/talk    # ❌ mandatory violation, blocked
pulumi config rm region --stack ediri/minecraft-on-digitalocean-yaml/talk   # undo
```

**Say:** TypeScript, Python, or Rego — same framework, four levels (advisory → mandatory
→ remediate → disabled), enforced across the whole org.

---

## Reset (between rehearsals) and Cleanup (after the talk)

```bash
# turn policies OFF so they don't block your other ediri stacks
pulumi policy disable ediri/minecraft-guardrails-ts
pulumi policy disable ediri/minecraft-guardrails-py
pulumi policy disable ediri/minecraft_guardrails_opa

# tear down any servers you deployed
cd examples/yaml && pulumi destroy --yes --stack ediri/minecraft-on-digitalocean-yaml/talk
# (repeat for any /tmp/mc-live stack)
```

> ⚠️ The two mandatory policies apply to **every DigitalOcean droplet in `ediri`** when
> enabled. Leave them **disabled** except during the policy section so you don't block
> unrelated deploys.

---

## Panic recovery

- **Cloud-init not up yet** → connect to the server you pre-warmed in Pre-flight, or
  show `pulumi env open` and the slides instead of a live world.
- **OPA policy "language plugin opa not found"** → first run only; the analyzer plugin
  installs on demand. Re-run the preview once.
- **`pulumi new` prompts hang on network** → fall back to `cd examples/yaml && pulumi up`.
- **DO token missing** → `pulumi config env add minecraft/demo` (it carries the token).
</content>
