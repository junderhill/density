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
var persistLocation string

func init() {
	ShootCmd.Flags().StringVar(&lat, "lat", "", "Location Lat to check")
	ShootCmd.Flags().StringVar(&long, "long", "", "Location Lon to check")
	ShootCmd.Flags().StringVar(&persistLocation, "persist-dir", "", "File system location to save JSON forecast data")
	viper.BindPFlag("latitude", ShootCmd.Flags().Lookup("latitude"))
	viper.BindPFlag("longitude", ShootCmd.Flags().Lookup("longitude"))
	viper.BindPFlag("persist-dir", ShootCmd.Flags().Lookup("persist-dir"))
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

		request := meteoblue.ForecastRequest{
			Location:   location,
			PersistDir: persistLocation,
		}

		//todo: add some verbose/debug logging (with flag)
		fmt.Printf("Location: %+v", location)

		_, _ = meteoblue.GetForecast(&request)

		panic("not implemented")
	},
}
