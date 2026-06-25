# Templates

Three starter templates that deploy a Minecraft server with the
`ediri/minecraft` component, one per language. They're identical in behaviour —
pick the language you'd rather read.

| Folder | Runtime | Project |
|--------|---------|---------|
| [`go`](./go) | Go | `minecraft-on-digitalocean-go` |
| [`csharp`](./csharp) | .NET / C# | `minecraft-on-digitalocean-csharp` |
| [`yaml`](./yaml) | YAML | `minecraft-on-digitalocean-yaml` |

## Use one

```bash
pulumi new https://github.com/dirien/pulumi-minecraft/tree/main/examples/go
pulumi install          # fetches the component and generates its local SDK
pulumi up
```

Each template prompts for the server knobs: `region`, `size`, `motd`,
`maxPlayers`, `difficulty`, `gamemode`, and `serverJarUrl`. Everything is
optional — skip a prompt and the component's default applies.

## Authentication

The templates don't ask for a DigitalOcean token. They expect it from a Pulumi
ESC environment, which keeps secrets out of the stack file:

```bash
pulumi config env add <org>/<your-env>   # supplies digitalocean:token
```

That's the seam the talk uses to show ESC composition — the same template,
different environments feeding it credentials and config.
