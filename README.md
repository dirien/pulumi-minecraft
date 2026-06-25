# Minecraft Server on DigitalOcean

A Pulumi component that runs a Java Edition Minecraft server on a single
DigitalOcean droplet. You give it a region and a handful of `server.properties`
values; it hands back an IP and a `host:port` you can paste straight into the
multiplayer menu.

Under the hood it creates two resources:

- a **Droplet** whose cloud-init installs a JDK, fetches `server.jar`, writes
  `server.properties`, accepts the EULA, and starts the server in a `screen`
  session;
- a **Firewall** that opens SSH and the game port, and nothing else.

It is written in Go with [pulumi-go-provider](https://github.com/pulumi/pulumi-go-provider),
so you can consume it from TypeScript, Python, Go, .NET, Java, or YAML.

## Install

```bash
pulumi package add ediri/minecraft@0.1.0
```

That generates a typed SDK in your project's language. In Go:

```go
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"<your-module>/sdk/go/minecraft"
)

server, err := minecraft.NewMinecraftServer(ctx, "mc", &minecraft.MinecraftServerArgs{
	Region:     pulumi.StringRef("fra1"),
	Motd:       pulumi.StringRef("Welcome to the build server"),
	Difficulty: pulumi.StringRef("hard"),
	MaxPlayers: pulumi.IntRef(30),
})
if err != nil {
	return err
}
ctx.Export("address", server.ServerAddress)
```

Every input is optional. Set only what you care about; the rest fall back to
the defaults below.

## Authentication

The component talks to the DigitalOcean API, so it needs a token. It reads the
standard `digitalocean:token` config value (or the `DIGITALOCEAN_TOKEN`
environment variable). The tidiest way to supply it is a Pulumi ESC environment
that mints the token and exposes it as config:

```bash
pulumi config env add <org>/<your-env>
```

Then nothing static lives in your stack file.

## Inputs

| Property | Type | Default | What it does |
|----------|------|---------|--------------|
| `region` | string | `fra1` | DigitalOcean region slug (fra1, nyc3, sfo3, sgp1, …). |
| `size` | string | `s-2vcpu-4gb` | Droplet size slug. Minecraft is hungry for RAM. |
| `image` | string | `ubuntu-24-04-x64` | Base image. Must be Debian/Ubuntu — cloud-init uses apt. |
| `serverJarUrl` | string | _latest_ | Direct `server.jar` URL. Empty means resolve the latest release from Mojang. |
| `javaMemory` | string | `3G` | Heap size for `-Xmx` and `-Xms`. |
| `sshKeys` | string[] | _none_ | DigitalOcean SSH key IDs or fingerprints to add to the droplet. |
| `motd` | string | `A Minecraft server, deployed with Pulumi` | Message of the day in the server list. |
| `maxPlayers` | int | `20` | Maximum concurrent players. |
| `difficulty` | string | `easy` | `peaceful`, `easy`, `normal`, or `hard`. |
| `gamemode` | string | `survival` | `survival`, `creative`, `adventure`, or `spectator`. |
| `pvp` | bool | `true` | Whether players can damage each other. |
| `onlineMode` | bool | `true` | Verify players against Mojang's auth servers. |
| `viewDistance` | int | `10` | Render distance in chunks (3–32). |
| `serverPort` | int | `25565` | TCP port the server listens on (also opened in the firewall). |
| `levelSeed` | string | _random_ | World-generation seed. |

## Outputs

| Property | Type | What it is |
|----------|------|------------|
| `ipv4Address` | string | Public IPv4 of the droplet. |
| `dropletId` | string | DigitalOcean droplet ID. |
| `serverAddress` | string | The `host:port` players connect to. |

## A few things worth knowing

- **The EULA.** Deploying writes `eula=true`. By running this you accept the
  [Minecraft EULA](https://www.minecraft.net/eula) on the server's behalf.
- **Jar resolution.** Leave `serverJarUrl` empty and the droplet asks Mojang's
  version manifest for the current release. Pin a URL when you want a specific
  version.
- **It costs money.** The droplet bills by the hour while it runs. Run
  `pulumi destroy` when you're done with the world.
- **Java Edition only.** Bedrock clients can't connect.
- **First boot takes a few minutes.** Installing the JDK and generating the
  world happen after `pulumi up` returns, so give it a moment before the server
  shows up in the list.
