using System.Collections.Generic;
using Pulumi;
using Minecraft.Minecraft;

return await Deployment.RunAsync(() =>
{
    var config = new Config();
    var args = new MinecraftServerArgs();

    // Only forward knobs that were actually set; the component fills in
    // sensible defaults for the rest.
    var region = config.Get("region");
    if (!string.IsNullOrEmpty(region)) args.Region = region;

    var size = config.Get("size");
    if (!string.IsNullOrEmpty(size)) args.Size = size;

    var motd = config.Get("motd");
    if (!string.IsNullOrEmpty(motd)) args.Motd = motd;

    var difficulty = config.Get("difficulty");
    if (!string.IsNullOrEmpty(difficulty)) args.Difficulty = difficulty;

    var gamemode = config.Get("gamemode");
    if (!string.IsNullOrEmpty(gamemode)) args.Gamemode = gamemode;

    var serverJarUrl = config.Get("serverJarUrl");
    if (!string.IsNullOrEmpty(serverJarUrl)) args.ServerJarUrl = serverJarUrl;

    var maxPlayers = config.GetInt32("maxPlayers");
    if (maxPlayers.HasValue) args.MaxPlayers = maxPlayers.Value;

    var server = new MinecraftServer("mc", args);

    return new Dictionary<string, object?>
    {
        ["serverAddress"] = server.ServerAddress,
        ["ipv4Address"] = server.Ipv4Address,
        ["dropletId"] = server.DropletId,
    };
});
