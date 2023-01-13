package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	stdhttp "net/http"

	"github.com/deepfence/ThreatMapper/deepfence_ctl/http"
	"github.com/deepfence/ThreatMapper/deepfence_ctl/output"
	"github.com/deepfence/ThreatMapper/deepfence_server_client"
	"github.com/deepfence/ThreatMapper/deepfence_utils/log"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graph control",
	Long:  `This subcommand controls graph with remote server`,
}

var graphTopologySubCmd = &cobra.Command{
	Use:   "topology",
	Short: "Get Topology graph",
	Long:  `This subcommand retrieve the topology graph`,
	Run: func(cmd *cobra.Command, args []string) {
		host_filter, _ := cmd.Flags().GetString("host-filter")
		host_entries := strings.Split(host_filter, ",")

		region_filter, _ := cmd.Flags().GetString("region-filter")
		region_entries := strings.Split(region_filter, ",")

		provider_filter, _ := cmd.Flags().GetString("provider-filter")
		provider_entries := strings.Split(provider_filter, ",")

		k8s_filter, _ := cmd.Flags().GetString("kubernetes-filter")
		k8s_entries := strings.Split(k8s_filter, ",")

		pod_filter, _ := cmd.Flags().GetString("pod-filter")
		pod_entries := strings.Split(pod_filter, ",")

		filters := deepfence_server_client.ReportersTopologyFilters{
			CloudFilter:      provider_entries,
			HostFilter:       host_entries,
			RegionFilter:     region_entries,
			KubernetesFilter: k8s_entries,
			PodFilter:        pod_entries,
		}

		root, _ := cmd.Flags().GetString("root")

		var res *deepfence_server_client.ApiDocsGraphResult
		var rh *stdhttp.Response
		var err error
		switch root {
		case "":
			req := http.Client().TopologyApi.GetTopologyGraph(context.Background())
			req = req.ReportersTopologyFilters(filters)
			res, rh, err = http.Client().TopologyApi.GetTopologyGraphExecute(req)
		case "hosts":
			req := http.Client().TopologyApi.GetHostsTopologyGraph(context.Background())
			req = req.ReportersTopologyFilters(filters)
			res, rh, err = http.Client().TopologyApi.GetHostsTopologyGraphExecute(req)
		case "containers":
			req := http.Client().TopologyApi.GetContainersTopologyGraph(context.Background())
			req = req.ReportersTopologyFilters(filters)
			res, rh, err = http.Client().TopologyApi.GetContainersTopologyGraphExecute(req)
		case "pods":
			req := http.Client().TopologyApi.GetPodsTopologyGraph(context.Background())
			req = req.ReportersTopologyFilters(filters)
			res, rh, err = http.Client().TopologyApi.GetPodsTopologyGraphExecute(req)
		case "kubernetes":
			req := http.Client().TopologyApi.GetKubernetesTopologyGraph(context.Background())
			req = req.ReportersTopologyFilters(filters)
			res, rh, err = http.Client().TopologyApi.GetKubernetesTopologyGraphExecute(req)
		default:
			log.Fatal().Msgf("Unsupported root:%s", root)
		}

		if err != nil {
			log.Fatal().Msgf("Fail to execute: %v: %v", err, rh)
		}
		output.Out(res)
	},
}

var graphThreatSubCmd = &cobra.Command{
	Use:   "threat",
	Short: "Get Threat graph",
	Long:  `This subcommand retrieve the threat graph`,
	Run: func(cmd *cobra.Command, args []string) {
		req := http.Client().ThreatApi.GetThreatGraph(context.Background())
		res, rh, err := http.Client().ThreatApi.GetThreatGraphExecute(req)

		if err != nil {
			log.Fatal().Msgf("Fail to execute: %v: %v", err, rh)
		}
		output.Out(res)
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
	graphCmd.AddCommand(graphTopologySubCmd)

	graphTopologySubCmd.PersistentFlags().String("host-filter", "", "CSV host filter")
	graphTopologySubCmd.PersistentFlags().String("region-filter", "", "CSV region filter")
	graphTopologySubCmd.PersistentFlags().String("provider-filter", "", "CSV provider filter")
	graphTopologySubCmd.PersistentFlags().String("kubernetes-filter", "", "CSV k8s filter")
	graphTopologySubCmd.PersistentFlags().String("pod-filter", "", "CSV pod filter")

	graphTopologySubCmd.PersistentFlags().String("root", "", "Root can be: ''/hosts/containers/pods/kubernetes")

	graphCmd.AddCommand(graphThreatSubCmd)
}