package shoot

import (
	"fmt"
	"os"

	"github.com/junderhill/density/internal/location"
	"github.com/junderhill/density/internal/meteoblue"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var long string
var lat string

func init() {
	ShootCmd.Flags().StringVar(&lat, "lat", "", "Location Lat to check")
	ShootCmd.Flags().StringVar(&long, "long", "", "Location Lon to check")
	viper.BindPFlag("latitude", ShootCmd.Flags().Lookup("latitude"))
	viper.BindPFlag("longitude", ShootCmd.Flags().Lookup("longitude"))
	ShootCmd.MarkFlagRequired("latitude")
	ShootCmd.MarkFlagRequired("longitude")
}

var ShootCmd = &cobra.Command{
	Use:   "shoot",
	Short: "Provides a score for a given location as to whether the conditions are suitable for landscape photography",
	Run: func(cmd *cobra.Command, args []string) {
		location, err := location.NewLocation(lat, long)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//todo: add some verbose/debug logging (with flag)
		fmt.Printf("Location: %+v", location)

		_, _ = meteoblue.GetForecast(location)

		panic("not implemented")
	},
}
