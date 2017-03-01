#Treadmill

A tool to export orders from Woocommerce to fulfillment providers

See this [realtime board](https://realtimeboard.com/app/board/o9J_k0pqWJc=/) for a system overview

## Configuration

 The following env vars are required to run the binary:

 `WOOF_REG` to specify the AWS region for the bucket e.g. `eu-west-1`

 `WOOF_BUC` to specify the name of the AWS S3 bucket e.g. `treadmill`

 `WOOF_ENV` to specify the environment e.g. `dev`, `test` or `uat`

 A TOML file in the `conf` folder of the bucket provides further detailed configuration

## Development


```
docker-compose run treadmill gvt fetch github.com/aws/aws-sdk-go/service/s3
```
