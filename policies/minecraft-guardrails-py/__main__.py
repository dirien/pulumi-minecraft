"""Mandatory guardrail: Minecraft droplets must run in an approved region."""

from pulumi_policy import (
    EnforcementLevel,
    PolicyPack,
    ReportViolation,
    ResourceValidationArgs,
    ResourceValidationPolicy,
)

DROPLET = "digitalocean:index/droplet:Droplet"
ALLOWED_REGIONS = {"fra1", "nyc1", "nyc3", "ams3", "sfo3"}


def restrict_region(args: ResourceValidationArgs, report_violation: ReportViolation):
    if args.resource_type != DROPLET:
        return
    region = args.props.get("region")
    if region not in ALLOWED_REGIONS:
        report_violation(
            f"Droplet region '{region}' is not approved. "
            f"Use one of: {', '.join(sorted(ALLOWED_REGIONS))}."
        )


PolicyPack(
    name="minecraft-guardrails-py",
    enforcement_level=EnforcementLevel.MANDATORY,
    policies=[
        ResourceValidationPolicy(
            name="restrict-droplet-region",
            description="Droplets must run in an approved DigitalOcean region.",
            enforcement_level=EnforcementLevel.MANDATORY,
            validate=restrict_region,
        ),
    ],
)
