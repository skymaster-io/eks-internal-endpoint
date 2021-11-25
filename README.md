# eks-internal-endpoint

get eks internal IP for EKS endpoint

Build instruction:

```shell
go get "github.com/aws/aws-sdk-go/aws/session"
go get  "github.com/aws/aws-sdk-go/service/ec2"
go get  "github.com/aws/aws-sdk-go/aws"
go get  "github.com/aws/aws-sdk-go/service/eks"
go build eks-internal-endpoint.go
```

Run:

```shell
eks-internal-endpoint <EKS-cluster-name>
```
