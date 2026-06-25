---
theme: "@pulumi/slidev-theme"
title: "Beyond pulumi up: A Tour of Pulumi Cloud and the Neo Agent"
info: |
  Beyond pulumi up: A Tour of Pulumi Cloud and the Neo Agent.
  Engin Diri — Pulumi.
transition: slide-left
mdc: true
canvasWidth: 1920
aspectRatio: 16/9
highlighter: shiki
lineNumbers: false
layout: cover
defaults:
  layout: default
---

<div class="absolute inset-0 flex flex-col justify-center items-start px-20">
  <h1 class="!text-[5rem] !leading-[1.04] !font-semibold !tracking-tight !mb-6 !max-w-[95%]">
    Beyond pulumi up: A Tour of Pulumi Cloud and the Neo Agent
  </h1>
  <img src="/img/event-logo.svg" class="!mt-4 h-[6.5rem] w-auto" alt="Event logo" />
  <p class="!mt-5 !text-[1.8rem] text-[var(--p-fg-muted)] !m-0 !leading-relaxed">
    Engin Diri · Pulumi
  </p>
</div>

<!--
30s hook. Read the title, then the promise: you already use Pulumi Cloud. You just
call it "the backend." It does a lot more than hold your state. I'll show three
features I actually use (IDP, ESC, Policy), then Neo — the agent that drives all of
them. You'll leave with a short list to turn on Monday.
-->

---

<div class="absolute inset-0 flex items-center px-24 gap-20">
  <div class="flex-shrink-0">
    <img src="/img/engin-diri.jpg" class="w-[28rem] rounded-2xl shadow-xl border-4" style="border-color: rgba(126,107,255,0.45)" alt="Engin Diri" />
  </div>
  <div class="flex-1">
    <h1 class="!text-[7rem] !leading-[1.02] !font-semibold !tracking-tight !mb-4 !text-[var(--p-primary)]">Engin Diri</h1>
    <p class="!text-[2.5rem] !leading-relaxed !m-0 opacity-90">
      Senior Solutions Architect at <strong class="!text-[var(--p-primary)]">Pulumi</strong>
    </p>
    <div class="!mt-8 flex items-center gap-8 !text-[1.5rem] opacity-70">
      <span class="flex items-center gap-2"><carbon-logo-x /> @_ediri</span>
      <span class="flex items-center gap-2"><carbon-logo-linkedin /> engin-diri</span>
      <span class="flex items-center gap-2"><carbon-logo-github /> dirien</span>
    </div>
    <p class="!mt-10 !text-[1.75rem] !leading-relaxed opacity-70 !m-0">
      Building platform tooling and infrastructure-as-code.<br/>
      Helping teams ship cloud infrastructure faster.
    </p>
  </div>
</div>

<!--
I'm Engin, Senior Solutions Architect at Pulumi. I build platform tooling and help
teams ship infrastructure faster. Today: a tour of the Pulumi Cloud features I lean
on every week, and the Neo agent that ties them together.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[7.2rem] !font-semibold !tracking-tight !m-0 !text-[var(--p-primary)]">What is Pulumi?</h1>
</div>

---

<div class="absolute inset-0 flex items-center justify-center overflow-hidden">
  <img src="/img/eye-contact-meme.png" class="w-full h-full object-cover" alt="When the speaker asks a question and you're trying to avoid eye contact" />
</div>

<style scoped>
:deep(.pulumi-accent-bar),
:deep(.pulumi-footer) { display: none !important; }
:deep(.pulumi-slide-body) { padding: 0 !important; }
</style>

---

# What is Pulumi?

<div class="zoom-content">

<p v-click class="!mt-8 !text-[1.4rem] !leading-relaxed">
  <span class="hl">TypeScript, Python, Go, .NET, Java, YAML.</span> Pick the language your team
  already speaks. <span class="hl-soft">Loops, conditionals, abstractions, tests.</span>
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  <span class="hl">Not</span> a config <span class="hl">DSL.</span>
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  And it matters more in the <span class="hl">agent era.</span> AI coding agents already speak
  these languages fluently. They can <span class="hl-soft">read, refactor, and test</span> the same
  code your humans do.
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  A config DSL puts a <span class="hl-strike">translation layer</span> between
  intent and execution which is not needed and gets in the way of agents doing their thing.
</p>

</div>

<style scoped>
.zoom-content { zoom: 1.5; }
/* Highlight key phrases. */
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft {
  background: rgba(126, 107, 255, 0.16);
  border-radius: 5px;
  padding: 0.05em 0.3em;
}
.hl-strike {
  color: var(--p-primary);
  font-weight: 700;
}
</style>

<!--
~45s. The phrase that lands: "not a config DSL." Then the AI angle:
agents work directly with real code, no HCL translation step. Real
languages are testable, composable, sit alongside the rest of your code
— so the same agent that writes your app code can ship the infra it runs on.
-->


---
class: dark
---

<div class="platform-image">
  <img src="/img/pulumi-platform.png" alt="Pulumi platform — IaC, Neo, Insights, ESC, IDP, plus Supergraph, Policy, and Workflow" />
</div>

<div v-click class="platform-qr">
  <img src="/img/pulumi-qr.png" alt="pulumi.com" />
  <div class="platform-qr__label">pulumi.com</div>
</div>

<style scoped>
:deep(.pulumi-footer) {
  display: none !important;
}
.platform-image {
  position: absolute;
  inset: 0;
  margin-top: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 0;
}
.platform-image img {
  width: 92%;
  height: 92%;
  object-fit: contain;
}
:deep(h1) {
  position: relative;
  z-index: 10;
}
.platform-qr {
  position: absolute;
  top: 2rem;
  right: 3rem;
  z-index: 20;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
}
.platform-qr img {
  width: 8.5rem;
  height: 8.5rem;
  background: #fff;
  border-radius: 12px;
  padding: 0.5rem;
  box-shadow: 0 10px 28px rgba(0, 0, 0, 0.5);
}
.platform-qr__label {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--p-fg);
}
</style>

<!--
~30s. ESC was just one slice. Walk left-to-right across the diagram:
IaC, Neo (the AI control plane), Insights, ESC, IDP, plus the
Supergraph / Policy / Workflow cross-cutting layers. The message is
"the platform is bigger than what fits in this workshop."
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[7.5rem] !leading-[1.04] !font-semibold !tracking-tight !m-0 !text-[var(--p-primary)]">Pulumi Cloud</h1>
</div>

<!--
Section opener. This is where pulumi up actually phones home — and it's the home
for the three products I'm about to show you.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <p class="!text-[4.68rem] !leading-snug !m-0 opacity-90 !max-w-[85%]">
    Where <code>pulumi up</code> phones home.<br/><br/>Guess what, it does a lot more than hold your state.
  </p>
</div>

<!--
You already use Pulumi Cloud. You just call it "the backend." State, secrets,
RBAC, deployment history. That's the floor, not the ceiling. The same dashboard
also runs your deploys, searches every resource across AWS/Azure/GCP (even the
stuff Pulumi never created), catches drift, and keeps an audit log. And it's where
ESC, Policy, IDP, and Neo live now.
-->

---
class: dark
---

<div class="absolute inset-0 flex items-center justify-center">
  <img src="/img/pulumi-dashboard.png" class="w-full h-full object-contain" alt="Pulumi Cloud dashboard — New task, stacks, environments, resources, Neo Tasks, Policy findings" />
</div>

<style scoped>
:deep(.pulumi-accent-bar),
:deep(.pulumi-footer) { display: none !important; }
:deep(.pulumi-slide-body) { padding: 0 !important; background: #0a0a0a; }
</style>

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <p class="!text-[5.62rem] !leading-snug !font-medium !m-0 !max-w-[88%] opacity-90">
    Three features I'd actually tell my mom about.
  </p>
</div>

<!--
No feature tour. I picked three I actually use every week — IDP, ESC, Policy —
and I'm saving Neo for last, because it's the one that changes how the day feels.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[7.5rem] !leading-[1.04] !font-semibold !tracking-tight !m-0 !text-[var(--p-primary)]">Pulumi IDP</h1>
</div>

<!--
Product divider — the platform-team product. Name only; the promise lands on the next slide.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[6.5rem] !leading-[1.05] !font-semibold !tracking-tight !m-0 !text-[var(--p-fg)] !max-w-[92%]">Golden paths, not tickets</h1>
</div>

<!--
The pitch: the platform team writes the IaC once, everyone else clicks a button.
-->

---

# Golden paths, not tickets

<div class="zoom-content">

<p v-click class="!mt-8 !text-[1.4rem] !leading-relaxed">
  Platform team publishes a component <span class="hl">once</span>:
  <span class="hl-soft">pulumi package publish</span>. Versioned, in the registry, docs generated for you.
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  Developers <span class="hl">never open an editor.</span> New Project, pick a template,
  fill in two fields, Create. That's a <span class="hl-soft">no-code stack.</span>
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  Config lives in <span class="hl">ESC</span>, not scattered across <code>.env</code> files and <code>tfvars</code> that drift out of sync. Day two: tweak and redeploy from the UI.
</p>

</div>

<style scoped>
.zoom-content { zoom: 1.5; }
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft { background: rgba(126, 107, 255, 0.16); border-radius: 5px; padding: 0.05em 0.3em; }
</style>

<!--
Before: dev needs a database, files a Jira ticket, writes a paragraph, waits two
days for someone to hand-write Terraform and get it reviewed. After: the platform
team already published a "service-with-database" template. Dev clicks New Project,
picks it, names it, clicks Create. Live in under a minute, on the exact pattern the
platform team signed off on. No IaC written by the dev, no ticket.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[10rem] !leading-none !font-bold !tracking-[0.08em] !m-0 !text-[var(--p-fg)]">DEMO</h1>
</div>

<!--
Live demo — the IDP golden path: publish a component, then New Project → template → fill two fields → Create.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[7.5rem] !leading-[1.04] !font-semibold !tracking-tight !m-0 !text-[var(--p-primary)]">Pulumi ESC</h1>
</div>

<!--
Product divider — secrets & short-lived creds. Name only; the promise lands on the next slide.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[6.5rem] !leading-[1.05] !font-semibold !tracking-tight !m-0 !text-[var(--p-fg)] !max-w-[92%]">Kill config and secret sprawl</h1>
</div>

<!--
The one-liner: stop pasting AWS keys into .env files. Run any command with creds
that die in an hour.
-->

---

# Kill config and secret sprawl

```bash
esc run acme/platform/aws-prod -- aws s3 ls
# creds minted on the spot · short-lived · no `aws configure`
```

<div class="esc-points">

<p v-click class="!mt-7 !text-[1.5rem] !leading-relaxed">
  Works for <span class="hl">any</span> command: <code>aws</code>, <code>kubectl</code>, your app, CI. Not just <code>pulumi up</code>.
</p>

<p v-click class="!mt-5 !text-[1.5rem] !leading-relaxed">
  Environments <span class="hl">import each other.</span> A <span class="hl-soft">base</span> env holds the defaults; <code>demo</code> and <code>prod</code> import it and override just what's different. Secrets resolve live from 1Password, Vault, or AWS, so when you rotate one upstream, every consumer gets the new value.
</p>

<p v-click class="!mt-5 !text-[1.5rem] !leading-relaxed">
  The OIDC setup is <span class="hl-soft">a few lines of YAML.</span> Delete the IAM user. Nothing static is left to leak.
</p>

</div>

<style scoped>
/* A focused, larger code snippet — not a full-width band. */
:deep(.slidev-code) {
  font-size: 1.7rem !important;
  line-height: 1.55 !important;
  width: fit-content;
  max-width: 95%;
  margin: 1rem 0 0;
  padding: 1.1rem 1.6rem;
  border-radius: 12px;
}
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft { background: rgba(126, 107, 255, 0.16); border-radius: 5px; padding: 0.05em 0.3em; }
</style>

<!--
The classic: aws s3 ls → "Unable to locate credentials." Old fix: dig out an access
key, aws configure, hope it's the right account. ESC fix: esc run, and it lists the
buckets using a session token minted on the spot. It's short-lived — commonly an hour,
set by `duration` / the role's max session, configurable up to 12h. Same env feeds your
stack, your laptop, and GitHub Actions. One source of truth instead of the same secret
copied into five places that all drift apart.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[10rem] !leading-none !font-bold !tracking-[0.08em] !m-0 !text-[var(--p-fg)]">DEMO</h1>
</div>

<!--
Live demo — ESC composition. A `minecraft/base` env imports `pulumi-idp/auth` for the
DigitalOcean creds and sets the default server.properties. `demo` and `prod` import base
and override just the MOTD, gamemode, and difficulty. `pulumi env open` shows the merged
result; point a stack at demo vs prod and the same template ships a different server.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[7.5rem] !leading-[1.04] !font-semibold !tracking-tight !m-0 !text-[var(--p-primary)]">Pulumi Policy</h1>
</div>

<!--
Product divider — guardrails. Name only; the promise lands on the next slide.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[5rem] !leading-[1.12] !font-semibold !tracking-tight !m-0 !text-[var(--p-fg)] !max-w-[90%]">Guardrails, because nobody should dance on razor blades</h1>
</div>

<!--
The one-liner: the policy runs on pulumi preview, so the public bucket gets blocked
before it ever exists.
-->

---

# Guardrails, because nobody should dance on razor blades

```typescript
// require-minecraft-tag · mandatory · runs on every `pulumi preview`
validateResourceOfType(digitalocean.Droplet, (droplet, _, report) => {
  if (!droplet.tags?.includes("minecraft"))
    report("Droplets must be tagged 'minecraft'.");
});
```

<div class="policy-points">

<p v-click class="!mt-6 !text-[1.5rem] !leading-relaxed">
  Real <span class="hl">code</span>, not YAML: <span class="hl-soft">TypeScript, Python, or OPA/Rego.</span> It runs on <code>preview</code>, so the bad config is caught <span class="hl">before it exists</span>, not by a scanner three hours later.
</p>

<p v-click class="!mt-5 !text-[1.5rem] !leading-relaxed">
  Four levels per policy: <span class="hl-soft">advisory · mandatory · remediate · disabled.</span> Same pack, different teeth.
</p>

<p v-click class="!mt-5 !text-[1.5rem] !leading-relaxed">
  <span class="hl">Org-wide.</span> Publish the pack, enable it for a <span class="hl-soft">Policy Group</span>, and every stack in the org gets checked. No opting out, nothing to wire per project.
</p>

</div>

<style scoped>
:deep(.slidev-code) {
  font-size: 1.3rem !important;
  line-height: 1.5 !important;
  width: fit-content;
  max-width: 96%;
  margin: 1.25rem 0 0;
  padding: 1rem 1.4rem;
  border-radius: 12px;
}
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft { background: rgba(126, 107, 255, 0.16); border-radius: 5px; padding: 0.05em 0.3em; }
</style>

<!--
A cloud security scanner finds the public bucket after it leaked. CrossGuard stops
the deploy, so the dangerous state never lands in the account. Roll out advisory in
dev to avoid the pitchfork mob, flip to mandatory in prod. Org-wide Policy Groups
live in Pulumi Cloud, plus compliance-ready packs for PCI/ISO/CIS.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[10rem] !leading-none !font-bold !tracking-[0.08em] !m-0 !text-[var(--p-fg)]">DEMO</h1>
</div>

<!--
Live demo — Policy. Publish the three packs (TypeScript, Python, Rego), enable them org-wide,
then `pulumi preview` the Minecraft template: the two mandatory checks pass, the Rego advisory
warns about backups. Flip a value — `region: lon1` — and the mandatory policy blocks it before
the droplet exists. Disable the packs afterward so they don't gate your other stacks.
-->

---

<div class="absolute inset-0 flex items-center justify-center px-20 text-center">
  <h1 class="!text-[6rem] !font-semibold !tracking-tight !leading-tight !m-0 !text-[var(--p-primary)] flex items-center gap-6">
    Why AI and Pulumi
    <span>=</span>
    <span class="!text-[7rem]">❤️</span>
  </h1>
</div>

<!--
The payoff: agents speak real languages, Pulumi *is* real languages, and ESC
hands them short-lived creds — so an agent can plan, provision, and verify
infra end-to-end. AI + Pulumi is a natural match.
-->

---

# Why AI and Pulumi <span class="title-eq">=</span> <span class="title-heart">❤️</span>

<div class="why8">

  <!-- Agent (base) -->
  <img src="/img/agent-bot.png" class="why8-img" style="left:18%; top:62%; height:25rem;" alt="Agent" />
  <div class="why8-lbl green" style="left:18%; top:93%;">Agent</div>

  <!-- Python logo (click 1) -->
  <img v-click="1" src="/img/python-logo.png" class="why8-img" style="left:54%; top:42%; height:15rem;" alt="Python" />
  <svg v-click="1" class="why8-arrows" viewBox="0 0 1920 1080" preserveAspectRatio="none">
    <defs><marker id="ah1" markerWidth="34" markerHeight="34" refX="28" refY="17" orient="auto" markerUnits="userSpaceOnUse"><path d="M0,0 L34,17 L0,34 Z" fill="#55cd48"/></marker></defs>
    <line x1="600" y1="710" x2="845" y2="565" stroke="#55cd48" stroke-width="9" stroke-linecap="round" marker-end="url(#ah1)"/>
  </svg>
  <div v-click="[1, 4]" class="why8-lbl green" style="left:37%; top:77%;">Written in<br/>Python</div>
  <div v-click="4" class="why8-lbl green" style="left:37%; top:77%;">Written in<br/><span class="struck">Python</span> TypeScript</div>

  <!-- IaC (click 2) -->
  <img v-click="2" src="/logos/pulumi-logo-mark-color-light.svg" class="why8-img" style="left:73%; top:73%; height:9rem;" alt="Pulumi IaC" />
  <div v-click="2" class="why8-lbl green" style="left:72.5%; top:90%;">IaC</div>
  <svg v-click="2" class="why8-arrows" viewBox="0 0 1920 1080" preserveAspectRatio="none">
    <defs><marker id="ah2" markerWidth="34" markerHeight="34" refX="28" refY="17" orient="auto" markerUnits="userSpaceOnUse"><path d="M0,0 L34,17 L0,34 Z" fill="#55cd48"/></marker></defs>
    <line x1="1255" y1="700" x2="1180" y2="615" stroke="#55cd48" stroke-width="9" stroke-linecap="round" marker-end="url(#ah2)"/>
  </svg>
  <div v-click="[2, 4]" class="why8-lbl green" style="left:59%; top:80%;">Written in<br/>Python</div>
  <div v-click="4" class="why8-lbl green" style="left:59%; top:80%;">Written in<br/><span class="struck">Python</span> TypeScript</div>

  <!-- Provision resources -> AWS wheel (click 3) -->
  <img v-click="3" src="/img/aws-services-wheel.png" class="why8-img" style="left:88%; top:37%; height:16rem;" alt="Provision resources" />
  <svg v-click="3" class="why8-arrows" viewBox="0 0 1920 1080" preserveAspectRatio="none">
    <defs><marker id="ah3" markerWidth="34" markerHeight="34" refX="28" refY="17" orient="auto" markerUnits="userSpaceOnUse"><path d="M0,0 L34,17 L0,34 Z" fill="#55cd48"/></marker></defs>
    <line x1="1545" y1="675" x2="1660" y2="600" stroke="#55cd48" stroke-width="9" stroke-linecap="round" marker-end="url(#ah3)"/>
  </svg>
  <div v-click="3" class="why8-lbl green" style="left:89%; top:67%;">Provision<br/>Resources</div>

  <!-- click 4: TypeScript logo on top of the Python logo -->
  <img v-click="4" src="/img/typescript-logo.png" class="why8-img" style="left:54%; top:42%; height:15rem;" alt="TypeScript" />

</div>

<style scoped>
:deep(.pulumi-slide-body) { position: relative !important; padding: 0 !important; }
:deep(.pulumi-slide-body > h1:first-child) { z-index: 5; }

.why8 { position: absolute; inset: 0; }

.why8-img { position: absolute; width: auto; transform: translate(-50%, -50%); }

.why8-lbl {
  position: absolute; transform: translate(-50%, -50%);
  font-weight: 800; font-size: 2.2rem; line-height: 1.05; text-align: center; white-space: nowrap;
}
.green { color: #55cd48; }
.struck { text-decoration: line-through; opacity: 0.6; }

.why8-arrows { position: absolute; inset: 0; width: 100%; height: 100%; pointer-events: none; overflow: visible; }

/* Title "= ❤️" matches slide 7's divider. */
.title-heart { font-size: 1.15em; vertical-align: -0.08em; }
</style>

---

# Why AI and Pulumi <span class="title-eq">=</span> <span class="title-heart">❤️</span>

<div class="zoom-content">

<p v-click class="!mt-7 !text-[1.4rem] !leading-relaxed">
  Agents already run Pulumi. They write it just as well, because it's the same TypeScript, Python, and Go they already know.
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  Pulumi ships <span class="hl">Agent Skills</span>, so the agent uses real components, ESC, and policy instead of guessing at APIs.
</p>

<div v-click class="flex flex-wrap gap-2 !mt-4">
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-overview</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-best-practices</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-component</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-esc</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-automation-api</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-terraform-to-pulumi</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-cdk-to-pulumi</span>
  <span class="font-mono !text-[1rem] rounded-md px-3 py-1" style="background: rgba(126,107,255,0.16); color: var(--p-primary)">pulumi-neo-handoff</span>
</div>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  I built this whole demo <span class="hl-soft">with a coding agent</span>: the component, the ESC envs, the policy packs.
</p>

<p v-click class="!mt-6 !text-[1.4rem] !leading-relaxed">
  <span class="hl">Neo</span> skips all of it. Those skills are for your other agents. It's Pulumi's own, so it already knows Pulumi and your stacks.
</p>

</div>

<style scoped>
.zoom-content { zoom: 1.3; margin-top: 1.5rem; }
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft { background: rgba(126, 107, 255, 0.16); border-radius: 5px; padding: 0.05em 0.3em; }
.hl-strike { text-decoration: line-through; opacity: 0.55; }
.title-heart { font-size: 1.15em; vertical-align: -0.08em; }
</style>

<!--
The authoring half of the ❤️. why8 showed agents PROVISION; this is that they AUTHOR — well.
Coding agents speak the languages; Pulumi's Agent Skills (pulumi-typescript, -python, -go,
-cli, -automation-api, -neo) inject the Pulumi patterns so they reach for components/ESC/policy
instead of inventing APIs. Be specific: this entire demo was built with a coding agent + these
skills. Then pivot to the Neo section — the one agent that needs no skill is Neo, because it's native.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[7.5rem] !leading-[1.04] !font-semibold !tracking-tight !m-0 !text-[var(--p-primary)]">Pulumi Neo</h1>
</div>

<!--
Product divider — the payoff product. Name only; the promise lands on the next slide.
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[6.4rem] !leading-[1.05] !font-semibold !tracking-tight !m-0 !text-[var(--p-fg)] !max-w-[92%]">Now an agent drives all of this</h1>
</div>

<!--
The payoff. Everything I just showed you — IDP, ESC, Policy — Neo can drive.
It shows up in three places: your terminal, Pulumi Cloud, and on your pull
requests. Let me show you each.
-->

---

# In your terminal

<div class="neo-float">
  <img src="/img/neo-cli.png" class="max-h-[100%] max-w-[222%] w-auto object-contain" alt="Pulumi Neo — coding agent CLI: pulumi neo --org ediri" />
</div>

<div v-click class="neo-qr">
  <img src="/img/neo-blog-qr.png" alt="Read the Pulumi Neo CLI blog post" />
  <div class="neo-qr__label">Read the<br/>announcement</div>
</div>

<style scoped>
/* Float the screenshot over the whole slide (incl. the footer band). */
.neo-float {
  position: absolute;
  top: 9rem;          /* clear the title */
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 20;        /* above the footer */
  pointer-events: none;
}

/* Click-revealed QR card, bottom-right, above the footer. */
.neo-qr {
  position: absolute;
  bottom: 2.5rem;
  right: 3rem;
  z-index: 21;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.6rem;
}
.neo-qr img {
  width: 9rem;
  height: 9rem;
  background: #fff;
  border-radius: 12px;
  padding: 0.5rem;
  box-shadow: 0 10px 28px rgba(0, 0, 0, 0.45);
}
.neo-qr__label {
  font-size: 1.05rem;
  font-weight: 600;
  text-align: center;
  line-height: 1.2;
  color: var(--p-fg);
}
</style>

<!--
The "turn it on" moment. Same Neo, now in your shell: `pulumi neo` from any project
with a Pulumi.yaml, no creds to provision, no context to paste. It picks up the CLIs
and kubeconfigs you already have. You stay in the loop: --approval-mode balanced runs
the safe stuff but stops before `pulumi up`; --permission-mode read-only lets it plan
but never mutate. Shift+Tab before your first message drops it into plan mode. And `-p`
lets Claude Code or Cursor hand a deploy off to Neo.
-->

---

# In Pulumi Cloud

<div class="neo-cols">
  <div class="neo-text">
    <p v-click>Same Neo, in the console. <span class="hl">Neo → Agent Tasks</span>: type the job in plain English, scoped to <span class="hl">your</span> permissions, not a god-mode token.</p>
    <p v-click>Tasks run <span class="hl-soft">async</span>. Close the laptop; Neo comes back with a <span class="hl">pull request</span>, so review and CI still gate the merge.</p>
    <p v-click><span class="hl-soft">Review</span> mode, the default, stops before preview, before <code>up</code>, before the PR.</p>
    <p v-click>Schedule it with <span class="hl">Automations</span>: a weekly encryption or backup audit that files the fix as a PR.</p>
  </div>
  <div class="neo-img">
    <img src="/img/neo-tasks.png" alt="Pulumi Cloud Neo Tasks — 'What do you want done today?', task list with status and token usage" />
  </div>
</div>

<style scoped>
.neo-cols {
  position: absolute;
  top: 8.5rem;
  left: 0;
  right: 0;
  bottom: 2.4rem;
  display: flex;
  align-items: center;
  gap: 2.5rem;
  padding: 0 3.5rem;
}
.neo-text { flex: 0 0 33%; }
.neo-text p { font-size: 1.15rem; line-height: 1.55; margin: 0 0 1.1rem; }
.neo-img { flex: 1 1 auto; min-width: 0; height: 100%; display: flex; align-items: center; justify-content: center; }
.neo-img img {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  border-radius: 10px;
  border: 1px solid rgba(126, 107, 255, 0.35);
  box-shadow: 0 12px 36px rgba(0, 0, 0, 0.5);
}
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft { background: rgba(126, 107, 255, 0.16); border-radius: 5px; padding: 0.05em 0.3em; }
</style>

<!--
Same Neo, in the console: Neo → Agent Tasks, type the job in plain English. It works across your
whole org, scoped to YOUR permissions, not a god-mode token. Tasks run async — close the laptop,
Neo comes back with a pull request, so your normal review and CI still gate the merge. Review mode,
the default, stops before preview, before `up`, and before it opens the PR. Put it on a schedule with
Automations: a weekly encryption or backup audit that files the fix as a PR (brand new, May '26).
Demo story: Friday a policy scan flags forty unencrypted buckets across a dozen stacks —
"find every unencrypted bucket and fix them" — it plans, shows the preview, you approve, a PR per repo.
-->

---

# On your pull requests

<div class="pr-card">
  <div class="pr-card__head">
    <carbon-logo-github class="!text-[1.6rem]" />
    <span class="pr-card__who">@pulumi-neo</span>
    <span class="pr-card__meta">reviewed <code>rds.ts</code></span>
  </div>
  <div class="pr-card__body">
    Heads up: this property forces a <strong class="rep">replace</strong> on
    <code>aws:rds/instance prod-db</code>. Pulumi will destroy and recreate it — downtime, and
    data loss without a final snapshot. Two stacks not in this PR, <code>api-prod</code> and
    <code>worker-prod</code>, read this endpoint and will need to re-resolve it.
  </div>
</div>

<p v-click class="cr-point !mt-8 !text-[1.6rem] !leading-relaxed">
  A bot that read your diff is cute. Neo read the <span class="hl">pulumi preview</span> and your
  <span class="hl">live state</span>, so it knows a one-word edit deletes prod.
</p>

<p v-click class="cr-point !mt-5 !text-[1.6rem] !leading-relaxed">
  It follows the <span class="hl-soft">blast radius across stacks</span>, comments inline, and you can argue
  back: <code>@pulumi-neo</code> in the thread. <span class="opacity-50">(public preview, GitHub.com)</span>
</p>

<style scoped>
.hl { color: var(--p-primary); font-weight: 600; }
.hl-soft { background: rgba(126, 107, 255, 0.16); border-radius: 5px; padding: 0.05em 0.3em; }
.cr-point { max-width: 90%; }
.pr-card {
  max-width: 84%;
  margin: 1.6rem 0 0;
  border: 1px solid rgba(126, 107, 255, 0.4);
  border-radius: 14px;
  background: rgba(126, 107, 255, 0.06);
  overflow: hidden;
}
.pr-card__head {
  display: flex; align-items: center; gap: 0.6rem;
  padding: 0.8rem 1.4rem;
  background: rgba(126, 107, 255, 0.14);
  font-size: 1.25rem;
}
.pr-card__who { font-weight: 700; color: var(--p-primary); }
.pr-card__meta { opacity: 0.6; font-size: 1.05rem; }
.pr-card__body { padding: 1.1rem 1.4rem; font-size: 1.4rem; line-height: 1.5; }
.pr-card__body .rep { color: #ff6b6b; }
</style>

<!--
The killer difference. A generic LLM PR bot reads the text of the diff. Neo reads the
pulumi preview output and comments inline — that's documented. So is the cross-stack
part: it can walk through impact on downstream stacks the PR never touched. The "forces
a replace → destroy/recreate prod-db" line in the card is MY illustration of what a
preview surfaces (replace ops are core Pulumi), not a feature the docs name explicitly —
so present it as "the kind of thing it catches," not a quoted capability. @pulumi-neo in
a comment and it answers in-thread. Honest status: public preview, GitHub.com only, GA
targeted for July 1 2026 (a forward target — it can slip; we're days away).
-->

---

<div class="absolute inset-0 flex flex-col items-center justify-center px-20 text-center">
  <h1 class="!text-[10rem] !leading-none !font-bold !tracking-[0.08em] !m-0 !text-[var(--p-fg)]">DEMO</h1>
</div>

<!--
Live demo — Neo, the three surfaces. `pulumi neo` in the terminal, Neo → Agent Tasks in
Pulumi Cloud, and @pulumi-neo on a pull request. Hand it the unencrypted-buckets job: it plans,
you approve, it opens a PR. The payoff that ties IDP, ESC, and Policy together.
-->

---

<div class="absolute inset-0 flex flex-col justify-center items-center px-20">
  <div class="opacity-80 tracking-[0.6em] uppercase !text-[1.6rem] !mb-4 text-[var(--p-fg-muted)]">Thank you</div>
  <h1 class="!text-[4.5rem] !leading-[1.02] !font-semibold !tracking-tight !mb-16 text-center">
    Stay in <span class="!text-[var(--p-primary)]">touch,</span>
  </h1>

  <div class="flex gap-16 justify-center items-start">
    <div class="text-center">
      <img src="/img/engin-diri.jpg" class="w-32 h-32 rounded-full mx-auto mb-4 border-4 object-cover" style="border-color: rgba(126,107,255,0.35)" alt="Engin Diri" />
      <div class="!text-[1.7rem] !font-bold">Engin Diri</div>
      <div class="opacity-60 !text-[1.2rem]">Pulumi</div>
      <div class="flex items-center justify-center gap-4 mt-2 !text-[1.1rem] opacity-60">
        <span class="flex items-center gap-1"><carbon-logo-github /> dirien</span>
        <span class="flex items-center gap-1"><carbon-logo-linkedin /> engin-diri</span>
      </div>
      <div class="mt-5 bg-white rounded-lg p-2 inline-block shadow-lg">
        <img src="/img/linkedin-qr.png" class="w-32 h-32" alt="LinkedIn QR" />
      </div>
    </div>
    <div class="text-center">
      <div class="w-32 h-32 rounded-full mx-auto mb-1 border-4 overflow-hidden flex items-center justify-center" style="border-color: rgba(126,107,255,0.35)">
        <carbon-logo-github class="!text-[11.25rem] leading-none" />
      </div>
      <div class="!text-[1.7rem] !font-bold">Demo + Component</div>
      <div class="opacity-60 !text-[1.2rem]">dirien/pulumi-minecraft</div>
      <div class="mt-2 !text-[1.1rem] opacity-0">&nbsp;</div>
      <div class="mt-5 bg-white rounded-lg p-2 inline-block shadow-lg">
        <img src="/img/repo-qr.png" class="w-32 h-32" alt="Demo repo QR" />
      </div>
    </div>
    <div class="text-center">
      <div class="w-32 h-32 rounded-full mx-auto mb-4 border-4 bg-white flex items-center justify-center" style="border-color: rgba(126,107,255,0.35)">
        <img src="/logos/pulumi-logo-mark-color-light.svg" class="w-20 h-20" alt="Pulumi" />
      </div>
      <div class="!text-[1.7rem] !font-bold">Pulumi</div>
      <div class="opacity-60 !text-[1.2rem]">pulumi.com</div>
      <div class="mt-2 !text-[1.1rem] opacity-0">&nbsp;</div>
      <div class="mt-5 bg-white rounded-lg p-2 inline-block shadow-lg">
        <img src="/img/pulumi-qr.png" class="w-32 h-32" alt="Pulumi website QR" />
      </div>
    </div>
  </div>
</div>

<!--
Thank you! Scan to connect on LinkedIn, or grab the slides and workshop code
from the repo. Then jump into Module 0.
-->
