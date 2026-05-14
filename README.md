# Twilio Resource Provider

The Twilio Resource Provider lets you manage [Twilio](https://www.twilio.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @nellisauction/pulumi-twilio
```

or `yarn`:

```bash
yarn add @nellisauction/pulumi-twilio
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_twilio
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/nellisauction/pulumi-twilio/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Twilio
```

## Configuration

The following configuration points are available for the `twilio` provider:

- `twilio:username` (environment: `TWILIO_API_KEY` or `TWILIO_ACCOUNT_SID`) - your API Key or Account SID
- `twilio:password` (environment: `TWILIO_API_SECRET` or `TWILIO_AUTH_TOKEN`) - your API Secret or Auth Token
- `twilio:accountSid` (environment: `TWILIO_SUBACCOUNT_SID` or `TWILIO_ACCOUNT_SID`) - sub-account SID (optional)
- `twilio:edge` (environment: `TWILIO_EDGE`) - Twilio edge location (optional)
- `twilio:region` (environment: `TWILIO_REGION`) - Twilio region (optional)

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/twilio/api-docs/).
