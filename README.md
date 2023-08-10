# clickupctl

This is the command line client for [ClickUp™](https://clickup.com).

## Installation

```console
go install github.com/moutend/clickupctl/cmd/clickupctl@latest
```

## Authentication

To use clickupctl, you'll need to set up a personal API token first. Here are the steps to follow:

1. Log into ClickUp.
2. Click on your avatar in the lower-left corner and select Apps.
3. Under API Token, click Generate.
4. Set the environment variable `CLICKUP_API_TOKEN` using the generated token.

```console
export CLICKUP_API_TOKEN='Your API token'
```

Run the following command to confirm that the authentication has been successfully completed.

```console
clickupctl whoami
```

## Usage

The basic format is:

```
clickupctl <resource> <operation>
```

For instance, use the following command to list all teams:

```console
clickupctl team list
```

For more details, run `clickupctl help`.

## LICENSE

MIT
