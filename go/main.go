package main

import (
	"context"
	"flag"
	"github.com/exoscale/egoscale"
	"log"
)

func main() {
	instancePoolId := ""
	exoscaleEndpoint := "https://api.exoscale.ch/v1/"
	exoscaleZoneId := ""
	exoscaleApiKey := ""
	exoscaleApiSecret := ""
	flag.StringVar(
		&instancePoolId,
		"instance-pool-id",
		instancePoolId,
		"ID of the instance pool to manage",
	)
	flag.StringVar(
		&exoscaleZoneId,
		"exoscale-zone-id",
		exoscaleZoneId,
		"Exoscale zone ID",
	)
	flag.StringVar(
		&exoscaleEndpoint,
		"exoscale-endpoint",
		exoscaleEndpoint,
		"Endpoint URL of the Exoscale API",
	)
	flag.StringVar(
		&exoscaleApiKey,
		"exoscale-api-key",
		exoscaleApiKey,
		"API key for Exoscale",
	)
	flag.StringVar(
		&exoscaleApiSecret,
		"exoscale-api-secret",
		exoscaleApiSecret,
		"API secret for Exoscale",
	)
	flag.Parse()

	zoneId, err := egoscale.ParseUUID(exoscaleZoneId)
	if err != nil {
		log.Fatalf("invalid zone ID (%v)", err)
	}

	poolId, err := egoscale.ParseUUID(instancePoolId)
	if err != nil {
		log.Fatalf("invalid pool ID (%v)", err)
	}

	client := egoscale.NewClient(exoscaleEndpoint, exoscaleApiKey, exoscaleApiSecret)
	ctx := context.Background()
	resp, err := client.RequestWithContext(ctx, egoscale.GetInstancePool{
		ZoneID: zoneId,
		ID:     poolId,
	})
	response := resp.(egoscale.GetInstancePoolResponse)
	if len(response.InstancePools) == 0 {
		log.Fatalf("instance pool not found")
	} else if len(response.InstancePools) > 1 {
		//This should never happen
		log.Fatalf("more than one instance pool returned")
	}
	instancePool := response.InstancePools[0]

	_, err = client.RequestWithContext(ctx, egoscale.ScaleInstancePool{
		ZoneID: zoneId,
		ID:     poolId,
		Size:   instancePool.Size + 1,
	})
	if err != nil {
		log.Fatalf("Failed to increase instance pool size (%v)", err)
	}
}
