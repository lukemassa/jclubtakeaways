This is for deploying the terraform to the AWS account (917497682277). The lambda code is uploaded by the github action, but this manages the infra.

There is also an SSM parameter used as a secret that this created, then I manually uploaded the content of.

## Production deployment protection

Before enabling the updated Lambda workflow, create the `lambda-production`
environment in the GitHub repository settings. Restrict its deployment branches
to `main` only (not all protected branches), add a required reviewer, and enable
preventing self-review. Restrict the existing `github-pages` environment to
`main` as well and add the same review protection. These settings are managed
in GitHub, not by this Terraform configuration or the workflow YAML.

The Lambda deployment role trusts the OIDC subject
`repo:lukemassa/jclubtakeaways:environment:lambda-production`. Environment-based
subjects do not include the branch, so the environment's branch restriction is
required. Configure the environment first, then review and apply the Terraform
trust-policy change before running the updated Lambda deployment. Until the
trust-policy change is applied, the updated workflow cannot assume the role.

Both workflows build with read-only repository permissions and deploy in a
separate job without checking out or executing repository code. Only deployment
jobs can request OIDC tokens. Actions are pinned to commit SHAs; verify new
commits in the official action repositories when updating those pins.
