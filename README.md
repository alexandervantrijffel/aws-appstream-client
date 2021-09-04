# AWS-AppStream-Client

Package that uses the AWS AppStream 2.0 API for managing AppStream sessions.

## AWS Authentication

The credentials for AWS authentication can be provided in multiple ways. Either pass environment variables as described [here](https://docs.aws.amazon.com/sdk-for-go/api/aws/session) or with a combination of AWS_PROFILE and a [credentials file](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) at ~/.awc/credentials.

Example of passing environment variables to the command:

AWS_REGION=eu-west-1 AWS_ACCESS_KEY_ID=<aws keyid> AWS_SECRET_ACCESS_KEY=<aws access key> go run ./cmd sessions --stack-name <stack> --fleet-name <fleet>

Another method is to copy the `.env.sample` file to `.env`, add the secrets to `.env` and run the following command:

````bash
export $(grep -v '^#' .env | xargs -d '\n')
```

Unset the exported environment variables with the following command

```bash
unset $(grep -v '^#' .env | sed -E 's/(.)=./\1/' | xargs)
```

## AWS Commands

### sessions

List active AppStream Sessions with this command, executed from the root of the repo.

```bash
go run ./cmd sessions --stack-name <stack> --fleet-name <fleet> | zap-pretty
````

In this example, [zap-pretty](https://github.com/maoueh/zap-pretty) is used to pretty format the output.
