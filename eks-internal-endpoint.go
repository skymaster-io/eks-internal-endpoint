package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"

	"os"
	"fmt"
	"strings"
)
var clusterName string

func main() {
	if (len(os.Args) != 2){
		fmt.Println("Usage: "+ os.Args[0] + " <aks cluster name>" )
		os.Exit(3)
	}
	clusterName = os.Args[1]
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	getEndpointURL := getEndpointURL(sess)
	ip := getEndpointIp(sess)
	
	fmt.Println(ip + "\t" + getEndpointURL )	
}

func getEndpointIp(sess *session.Session) string {
	ec2Svc := ec2.New(sess)
	filters := []*ec2.Filter{
		&ec2.Filter{
			Name: aws.String("description"),
			Values: []*string{
				aws.String("Amazon EKS " + clusterName ),
			},
		},
	}
	input := ec2.DescribeNetworkInterfacesInput{Filters: filters}
	result, err := ec2Svc.DescribeNetworkInterfaces(&input)
	if err != nil {
		fmt.Println("Error", err)		
	}
	if (len(result.NetworkInterfaces) == 0) {
		fmt.Println("Unable to find Network interface with description: \"Amazon EKS " + clusterName + "\"" )
		os.Exit(2)
	}
	return *result.NetworkInterfaces[0].PrivateIpAddress
}

func getEndpointURL(sess *session.Session) string {
	ec2Svc := eks.New(sess)
	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}
	result, err := ec2Svc.DescribeCluster(input)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(3)
	}
	return strings.Replace(*result.Cluster.Endpoint,"https://", "", 1)
}