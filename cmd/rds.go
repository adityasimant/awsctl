/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg"
)

// rdsCmd represent the rds command
var rdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "Retrieve information about AWS RDS instances",
	Long: `This command retrieves information about AWS RDS instances, including their instance IDs, engines, statuses, and endpoints.
Example :
	awsctl get rds`,
	RunE: getRdsCommand,
}

func getRdsCommand(cmd *cobra.Command, args []string) error {

	rdsInstances, err := pkg.GetRDSInstances()

	if err != nil {
		return errors.New("failed to retrive RDS instances : " + err.Error())
	}

	if len(rdsInstances) == 0 {
		fmt.Println("No RDS instances found")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 5, 2, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "DB_INSTANCE_ID\tENGINE\tSTATUS\tENDPOINT")

	for _, i := range rdsInstances {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", i.DBInstanceID, i.Engine, i.Status, i.Endpoint)
	}
	return nil
}

func init() {
	getCmd.AddCommand(rdsCmd)
}
