## Go AWS Lambda

A simple AWS Lambda function written in Go. This document describes how to deploy the function code to an AWS Lambda using Docker and AWS Elastic Container Registry (ECR). 

### Prerequisites

- AWS CLI installed and setup,
- Go installed,
- Docker installed.

### Deploy

#### 1. Build
```shell
docker build --platform linux/amd64 -t <AWS_ECR_REPOSITORY_URI>:latest .
```
- Replace `<AWS_ECR_REPOSITORY_URI>` with your registry URI, like `111122223333.dkr.ecr.us-east-1.amazonaws.com/repository-name`. If you don't have a repository yet, see how to [create a new repository](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/ecr/create-repository.html).

#### 2. Authenticate Docker to your ECR registry
```shell
aws ecr get-login-password --region <AWS_REGION> | docker login --username AWS --password-stdin <AWS_ECR_REGISTRY_URI>
```
- Replace `<AWS_REGION>` with a target AWS region, like `us-east-1`,
- Replace `<AWS_ECR_REGISTRY_URI>` with your registry URI, like `111122223333.dkr.ecr.us-east-1.amazonaws.com`.

#### 3. Push the image to a repository
```shell
docker push <AWS_ECR_REPOSITORY_URI>:latest
```

#### 4. Create a new Lambda function
```shell
aws lambda create-function \
  --function-name <NAME> \
  --package-type Image \
  --code ImageUri=<AWS_ECR_REPOSITORY_URI>:latest \
  --role <AWS_LAMBDA_EXECUTION_ROLE_ARN>
```
- Replace `<NAME>` with a function name,
- Replace `<AWS_ECR_REPOSITORY_URI>` with your repository URI, like `111122223333.dkr.ecr.us-east-1.amazonaws.com/repository-name`,
- Replace `<AWS_LAMBDA_EXECUTION_ROLE_ARN>` with ARN of the Lambda execution role (see how to [create an execution role](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-awscli.html#with-userapp-walkthrough-custom-events-create-iam-role)).

### Invoke Lambda using AWS CLI
```shell
aws lambda invoke \
  --function-name <NAME> \
  --payload '{ "Name": "New event" }' \
  --cli-binary-format raw-in-base64-out \
  response.json
```
- Replace `<NAME>` with a function name.

### Update function code

```shell
aws lambda update-function-code \
  --function-name <NAME> \
  --image-uri <AWS_ECR_REPOSITORY_URI>:latest
```
- Replace `<NAME>` with a function name,
- Replace `<AWS_ECR_REPOSITORY_URI>` with your repository URI, like `111122223333.dkr.ecr.us-east-1.amazonaws.com/repository-name`.