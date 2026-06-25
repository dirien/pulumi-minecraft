import { PolicyPack } from "@pulumi/policy";

const DROPLET = "digitalocean:index/droplet:Droplet";

new PolicyPack("minecraft-guardrails-ts", {
    policies: [
        {
            name: "require-minecraft-tag",
            description: "Every droplet must carry the 'minecraft' tag.",
            enforcementLevel: "mandatory",
            validateResource: (args, reportViolation) => {
                if (args.type !== DROPLET) {
                    return;
                }
                const tags: string[] = args.props.tags ?? [];
                if (!tags.includes("minecraft")) {
                    reportViolation(
                        "Droplet must be tagged 'minecraft' so the platform team can find and govern it.",
                    );
                }
            },
        },
    ],
});
