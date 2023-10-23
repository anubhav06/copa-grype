# 🔌Copa-Grype

Plugin for [Copacetic](https://github.com/project-copacetic/copacetic) to support patching grype produced results.  

Learn more about Copacetic's scanner plugins [here](https://project-copacetic.github.io/copacetic/website/next/scanner-plugins)  

## Installation

You can download the latest and previous versions of `copa-grype` from the [GitHub releases page](https://github.com/anubhav06/copa-grype/releases).
Make sure to add it to your `PATH` environment variable.


Otherwise, install using the CLI:

```shell
# Install the binary
curl -sL https://github.com/anubhav06/copa-grype/releases/latest/download/copa-grype -o copa-grype
# Add to PATH
export PATH=$PATH:/path/to/copagrype/directory
```

## Example Usage
```shell
# test plugin with example config
copa-grype grype_report.json
# this will print the report in JSON format. Example:
# {"apiVersion":"v1alpha1","metadata":{"os":{"type":"FakeOS","version":"42"},"config":{"arch":"amd64"}},"updates":[{"name":"foo","installedVersion":"1.0.0","fixedVersion":"1.0.1","vulnerabilityID":"VULN001"},{"name":"bar","installedVersion":"2.0.0","fixedVersion":"2.0.1","vulnerabilityID":"VULN002"}]}

```
