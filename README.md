### AWS-AppStream-Client

Package that uses the AWS AppStream 2.0 API for managing AppStream sessions.

#### Commands

List active AppStream Sessions with this command, executed from the root of the repo.

```bash
AWS_REGION=eu-west-1 AWS_ACCESS_KEY_ID=<aws keyid> AWS_SECRET_ACCESS_KEY=<aws access key> go run ./cmd sessions --stack-name <stack> --fleet-name <fleet> | zap-pretty
```

In this example, [zap-pretty](https://github.com/maoueh/zap-pretty) is used to pretty format the output.

For authentication, environment variables are used as described [here](https://docs.aws.amazon.com/sdk-for-go/api/aws/session). The credentials can also be provided with a combination of AWS_PROFILE and a [credentials file](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) at ~/.awc/credentials.
