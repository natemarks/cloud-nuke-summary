# cloud-nuke-summary

The first time I ran this on a sandbox account, I got 500K likes of output - too much to check manually

```bash
cloud-nuke inspect-aws \
--region us-east-1 \
--region us-west-1 \
--region us-east-2 \
--region us-west-2 \
--exclude-resource-type iam \
--exclude-resource-type iam-group \
--exclude-resource-type iam-policy \
--exclude-resource-type iam-role \
--exclude-resource-type iam-service-linked-role \
--exclude-resource-type secretsmanager \
--exclude-resource-type dynamodb \
--exclude-resource-type kmscustomerkeys \
--exclude-resource-type ec2-keypairs \
--exclude-resource-type s3 2>cloud-nuke-output.txt
```

So I created this little parser to summarize the output


