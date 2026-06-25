package main

import (
	"strconv"
	"strings"
)

// cloudInitVars are the values rendered into the leading variable block of the
// cloud-init script.
type cloudInitVars struct {
	ServerJarURL string
	JavaMemory   string
	Motd         string
	MaxPlayers   int
	Difficulty   string
	Gamemode     string
	Pvp          bool
	OnlineMode   bool
	ViewDistance int
	ServerPort   int
	LevelSeed    string
}

// buildUserData renders the full cloud-init script: a variable block computed
// from the component inputs, followed by the static body.
func buildUserData(v cloudInitVars) string {
	var b strings.Builder
	b.WriteString("#!/bin/bash\n")
	b.WriteString("SERVER_JAR_URL=" + shellQuote(v.ServerJarURL) + "\n")
	b.WriteString("JAVA_MEMORY=" + shellQuote(v.JavaMemory) + "\n")
	b.WriteString("MOTD=" + shellQuote(v.Motd) + "\n")
	b.WriteString("MAX_PLAYERS=" + shellQuote(strconv.Itoa(v.MaxPlayers)) + "\n")
	b.WriteString("DIFFICULTY=" + shellQuote(v.Difficulty) + "\n")
	b.WriteString("GAMEMODE=" + shellQuote(v.Gamemode) + "\n")
	b.WriteString("PVP=" + shellQuote(boolStr(v.Pvp)) + "\n")
	b.WriteString("ONLINE_MODE=" + shellQuote(boolStr(v.OnlineMode)) + "\n")
	b.WriteString("VIEW_DISTANCE=" + shellQuote(strconv.Itoa(v.ViewDistance)) + "\n")
	b.WriteString("SERVER_PORT=" + shellQuote(strconv.Itoa(v.ServerPort)) + "\n")
	b.WriteString("LEVEL_SEED=" + shellQuote(v.LevelSeed) + "\n")
	b.WriteString(cloudInitBody)
	return b.String()
}

// shellQuote single-quotes a value so it survives in the bash script verbatim.
func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// cloudInitBody is the static part of the script. It installs a JDK, resolves
// and downloads server.jar (latest release from Mojang when SERVER_JAR_URL is
// empty), accepts the EULA, writes server.properties from the variable block,
// and starts the server in a detached screen session.
const cloudInitBody = `
set -euo pipefail
export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get install -y wget curl jq screen apt-transport-https gpg ca-certificates

useradd -m -d /opt/minecraft minecraft || true
cd /opt/minecraft

# Resolve the server.jar URL and the Java version this build needs. When the
# component didn't pass a URL, ask Mojang's official version manifest for the
# latest release.
JAVA_MAJOR=21
if [ -z "${SERVER_JAR_URL:-}" ]; then
  MANIFEST="$(curl -fsSL https://piston-meta.mojang.com/mc/game/version_manifest_v2.json)"
  LATEST="$(echo "$MANIFEST" | jq -r '.latest.release')"
  VERSION_URL="$(echo "$MANIFEST" | jq -r --arg v "$LATEST" '.versions[] | select(.id==$v) | .url')"
  VERSION_JSON="$(curl -fsSL "$VERSION_URL")"
  SERVER_JAR_URL="$(echo "$VERSION_JSON" | jq -r '.downloads.server.url')"
  JAVA_MAJOR="$(echo "$VERSION_JSON" | jq -r '.javaVersion.majorVersion // 21')"
fi

# Install the JDK this server build actually needs, from Eclipse Adoptium.
# (Modern Minecraft jumps Java versions; e.g. MC 26.2 requires Java 25.)
wget -qO - https://packages.adoptium.net/artifactory/api/gpg/key/public | gpg --dearmor -o /etc/apt/trusted.gpg.d/adoptium.gpg
echo "deb https://packages.adoptium.net/artifactory/deb $(. /etc/os-release && echo "$VERSION_CODENAME") main" > /etc/apt/sources.list.d/adoptium.list
apt-get update
apt-get install -y "temurin-${JAVA_MAJOR}-jdk"
JAVA_BIN="$(command -v java)"

wget -O server.jar "$SERVER_JAR_URL"

# Accept the Mojang EULA (deploying the server implies acceptance).
echo "eula=true" > eula.txt

# server.properties — the knobs the template exposes.
cat > server.properties <<PROPS
motd=${MOTD}
max-players=${MAX_PLAYERS}
difficulty=${DIFFICULTY}
gamemode=${GAMEMODE}
pvp=${PVP}
online-mode=${ONLINE_MODE}
view-distance=${VIEW_DISTANCE}
server-port=${SERVER_PORT}
level-seed=${LEVEL_SEED}
enable-command-block=false
spawn-protection=0
PROPS

chown -R minecraft:minecraft /opt/minecraft

# Start the server in a detached screen session.
su - minecraft -c "cd /opt/minecraft && screen -dmS minecraft_server $JAVA_BIN -Xmx${JAVA_MEMORY} -Xms${JAVA_MEMORY} -jar server.jar nogui"
`
