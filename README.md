# eks-internal-endpoint
get eks internal IP of EKS endpoint

Build instruction:
```
go get "github.com/aws/aws-sdk-go/aws/session"
go get  "github.com/aws/aws-sdk-go/service/ec2"
go get  "github.com/aws/aws-sdk-go/aws"
go get  "github.com/aws/aws-sdk-go/service/eks"
go build eks-internal-endpoint.go
```

Run:
```
eks-internal-endpoint <EKS-cluster-name>
```
