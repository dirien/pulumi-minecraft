package main

import (
	"strconv"
	"strings"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MinecraftServerArgs are the inputs to a MinecraftServer.
//
// The scalar fields are plain Go types on purpose: the component reads them at
// construction time to render the cloud-init script, so they have to be
// immediate values rather than eventual Inputs.
type MinecraftServerArgs struct {
	Region       string   `pulumi:"region,optional"`
	Size         string   `pulumi:"size,optional"`
	Image        string   `pulumi:"image,optional"`
	ServerJarURL string   `pulumi:"serverJarUrl,optional"`
	JavaMemory   string   `pulumi:"javaMemory,optional"`
	SSHKeys      []string `pulumi:"sshKeys,optional"`

	// server.properties knobs — the things the template exposes and the ESC
	// chapter later composes.
	Motd         string `pulumi:"motd,optional"`
	MaxPlayers   int    `pulumi:"maxPlayers,optional"`
	Difficulty   string `pulumi:"difficulty,optional"`
	Gamemode     string `pulumi:"gamemode,optional"`
	Pvp          *bool  `pulumi:"pvp,optional"`
	OnlineMode   *bool  `pulumi:"onlineMode,optional"`
	ViewDistance int    `pulumi:"viewDistance,optional"`
	ServerPort   int    `pulumi:"serverPort,optional"`
	LevelSeed    string `pulumi:"levelSeed,optional"`
}

// Annotate adds descriptions and defaults to the generated schema, so the
// template prompts and registry docs read well.
func (a *MinecraftServerArgs) Annotate(an infer.Annotator) {
	an.Describe(&a.Region, "DigitalOcean region slug (e.g. fra1, nyc3, sfo3).")
	an.SetDefault(&a.Region, "fra1")
	an.Describe(&a.Size, "Droplet size slug. Minecraft is hungry for RAM — s-2vcpu-4gb is a sane floor.")
	an.SetDefault(&a.Size, "s-2vcpu-4gb")
	an.Describe(&a.Image, "Base image slug. Must be Debian/Ubuntu — the cloud-init uses apt.")
	an.SetDefault(&a.Image, "ubuntu-24-04-x64")
	an.Describe(&a.ServerJarURL, "Direct URL to the Minecraft server.jar. Leave empty to auto-resolve the latest release from Mojang.")
	an.Describe(&a.JavaMemory, "Heap size handed to the JVM for -Xmx and -Xms, e.g. 3G.")
	an.SetDefault(&a.JavaMemory, "3G")
	an.Describe(&a.SSHKeys, "DigitalOcean SSH key IDs or fingerprints to add to the droplet.")

	an.Describe(&a.Motd, "Message of the day shown in the multiplayer server list.")
	an.SetDefault(&a.Motd, "A Minecraft server, deployed with Pulumi")
	an.Describe(&a.MaxPlayers, "Maximum number of concurrent players.")
	an.SetDefault(&a.MaxPlayers, 20)
	an.Describe(&a.Difficulty, "World difficulty: peaceful, easy, normal, or hard.")
	an.SetDefault(&a.Difficulty, "easy")
	an.Describe(&a.Gamemode, "Default game mode: survival, creative, adventure, or spectator.")
	an.SetDefault(&a.Gamemode, "survival")
	an.Describe(&a.Pvp, "Whether players can damage each other. Defaults to true.")
	an.Describe(&a.OnlineMode, "Verify connecting players against Mojang's auth servers. Defaults to true.")
	an.Describe(&a.ViewDistance, "Render distance in chunks (3-32).")
	an.SetDefault(&a.ViewDistance, 10)
	an.Describe(&a.ServerPort, "TCP port the server listens on (also opened in the firewall).")
	an.SetDefault(&a.ServerPort, 25565)
	an.Describe(&a.LevelSeed, "Optional world-generation seed. Empty means a random world.")
}

// MinecraftServer is a Droplet running a Java Minecraft server, fronted by a
// DigitalOcean firewall that exposes SSH and the game port.
type MinecraftServer struct {
	pulumi.ResourceState

	// Ipv4Address is the public IPv4 of the droplet.
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// DropletID is the DigitalOcean droplet ID.
	DropletID pulumi.StringOutput `pulumi:"dropletId"`
	// ServerAddress is the host:port players connect to.
	ServerAddress pulumi.StringOutput `pulumi:"serverAddress"`
}

// Annotate documents the component and its outputs.
func (m *MinecraftServer) Annotate(an infer.Annotator) {
	an.Describe(&m, "A Java Minecraft server running on a DigitalOcean droplet, with a firewall that opens SSH and the game port.")
	an.Describe(&m.Ipv4Address, "Public IPv4 address of the droplet.")
	an.Describe(&m.DropletID, "DigitalOcean droplet ID.")
	an.Describe(&m.ServerAddress, "host:port players use to connect.")
}

// NewMinecraftServer provisions the droplet and firewall.
func NewMinecraftServer(ctx *pulumi.Context, name string, args MinecraftServerArgs, opts ...pulumi.ResourceOption) (*MinecraftServer, error) {
	comp := &MinecraftServer{}
	if err := ctx.RegisterComponentResource(p.GetTypeToken(ctx), name, comp, opts...); err != nil {
		return nil, err
	}

	region := orStr(args.Region, "fra1")
	size := orStr(args.Size, "s-2vcpu-4gb")
	image := orStr(args.Image, "ubuntu-24-04-x64")
	port := orInt(args.ServerPort, 25565)

	userData := buildUserData(cloudInitVars{
		ServerJarURL: args.ServerJarURL,
		JavaMemory:   orStr(args.JavaMemory, "3G"),
		Motd:         orStr(args.Motd, "A Minecraft server, deployed with Pulumi"),
		MaxPlayers:   orInt(args.MaxPlayers, 20),
		Difficulty:   orStr(args.Difficulty, "easy"),
		Gamemode:     orStr(args.Gamemode, "survival"),
		Pvp:          args.Pvp == nil || *args.Pvp,
		OnlineMode:   args.OnlineMode == nil || *args.OnlineMode,
		ViewDistance: orInt(args.ViewDistance, 10),
		ServerPort:   port,
		LevelSeed:    args.LevelSeed,
	})

	dropletArgs := &digitalocean.DropletArgs{
		Image:    pulumi.String(image),
		Size:     pulumi.String(size),
		Region:   pulumi.String(region),
		UserData: pulumi.String(userData),
		Tags: pulumi.StringArray{
			pulumi.String("pulumi"),
			pulumi.String("minecraft"),
		},
	}
	if len(args.SSHKeys) > 0 {
		dropletArgs.SshKeys = pulumi.ToStringArray(args.SSHKeys)
	}

	droplet, err := digitalocean.NewDroplet(ctx, name+"-droplet", dropletArgs, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}

	dropletIDInt := droplet.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)

	allV4 := pulumi.StringArray{pulumi.String("0.0.0.0/0"), pulumi.String("::/0")}

	_, err = digitalocean.NewFirewall(ctx, name+"-firewall", &digitalocean.FirewallArgs{
		DropletIds: pulumi.IntArray{dropletIDInt},
		InboundRules: digitalocean.FirewallInboundRuleArray{
			&digitalocean.FirewallInboundRuleArgs{
				Protocol:        pulumi.String("tcp"),
				PortRange:       pulumi.String("22"),
				SourceAddresses: allV4,
			},
			&digitalocean.FirewallInboundRuleArgs{
				Protocol:        pulumi.String("tcp"),
				PortRange:       pulumi.String(strconv.Itoa(port)),
				SourceAddresses: allV4,
			},
		},
		OutboundRules: digitalocean.FirewallOutboundRuleArray{
			&digitalocean.FirewallOutboundRuleArgs{
				Protocol:             pulumi.String("tcp"),
				PortRange:            pulumi.String("1-65535"),
				DestinationAddresses: allV4,
			},
			&digitalocean.FirewallOutboundRuleArgs{
				Protocol:             pulumi.String("udp"),
				PortRange:            pulumi.String("1-65535"),
				DestinationAddresses: allV4,
			},
			&digitalocean.FirewallOutboundRuleArgs{
				Protocol:             pulumi.String("icmp"),
				DestinationAddresses: allV4,
			},
		},
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}

	comp.Ipv4Address = droplet.Ipv4Address
	comp.DropletID = droplet.ID().ToStringOutput()
	comp.ServerAddress = pulumi.Sprintf("%s:%d", droplet.Ipv4Address, port)

	return comp, nil
}

func orStr(v, def string) string {
	if strings.TrimSpace(v) == "" {
		return def
	}
	return v
}

func orInt(v, def int) int {
	if v == 0 {
		return def
	}
	return v
}
