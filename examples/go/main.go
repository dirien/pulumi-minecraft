package main

import (
	"github.com/minecraft/minecraft/sdk/go/minecraft"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		// Only forward knobs that were actually set; the component fills in
		// sensible defaults for the rest.
		server, err := minecraft.NewMinecraftServer(ctx, "mc", &minecraft.MinecraftServerArgs{
			Region:       strPtr(cfg, "region"),
			Size:         strPtr(cfg, "size"),
			Motd:         strPtr(cfg, "motd"),
			Difficulty:   strPtr(cfg, "difficulty"),
			Gamemode:     strPtr(cfg, "gamemode"),
			ServerJarUrl: strPtr(cfg, "serverJarUrl"),
			MaxPlayers:   intPtr(cfg, "maxPlayers"),
		})
		if err != nil {
			return err
		}

		ctx.Export("serverAddress", server.ServerAddress)
		ctx.Export("ipv4Address", server.Ipv4Address)
		ctx.Export("dropletId", server.DropletId)
		return nil
	})
}

func strPtr(c *config.Config, key string) *string {
	if v := c.Get(key); v != "" {
		return &v
	}
	return nil
}

func intPtr(c *config.Config, key string) *int {
	if v, err := c.TryInt(key); err == nil {
		return &v
	}
	return nil
}
