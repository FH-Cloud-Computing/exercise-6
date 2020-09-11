import exoscale

if __name__ == "__main__":
    exo = exoscale.Exoscale(api_key="api-key-here", api_secret="api-secret-here", config_file="")

    zone = exo.compute.get_zone("zone-name-here")

    ip = exo.compute.get_instance_pool("instance-pool-id-here", zone=zone)
    ip.scale(ip.size + 1)
