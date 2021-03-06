package integration

type ClusterPlan struct {
	Cluster struct {
		Name                     string
		AdminPassword            string `yaml:"admin_password"`
		AllowPackageInstallation string `yaml:"allow_package_installation"`
		Networking               struct {
			Type             string
			PodCIDRBlock     string `yaml:"pod_cidr_block"`
			ServiceCIDRBlock string `yaml:"service_cidr_block"`
		}
		Certificates struct {
			Expiry          string
			LocationCity    string `yaml:"location_city"`
			LocationState   string `yaml:"location_state"`
			LocationCountry string `yaml:"location_country"`
		}
		SSH struct {
			User string
			Key  string `yaml:"ssh_key"`
			Port int    `yaml:"ssh_port"`
		}
	}
	Etcd struct {
		ExpectedCount int `yaml:"expected_count"`
		Nodes         []NodePlan
	}
	Master struct {
		ExpectedCount         int `yaml:"expected_count"`
		Nodes                 []NodePlan
		LoadBalancedFQDN      string `yaml:"load_balanced_fqdn"`
		LoadBalancedShortName string `yaml:"load_balanced_short_name"`
	}
	Worker struct {
		ExpectedCount int `yaml:"expected_count"`
		Nodes         []NodePlan
	}
}

type NodePlan struct {
	host       string
	ip         string
	internalip string
}
