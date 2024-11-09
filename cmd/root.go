package cmd

import (
	"fmt"
	"log"

	"github.com/navicore/idionautic-server/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "idionautic-server",
	Short: "Run the Idionautic telemetry server",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("port")
		iface := viper.GetString("interface")

		fmt.Printf("Starting server on %s:%d...\n", iface, port)
		api.StartServer(iface, port) // Pass the interface and port to the server
	},
}

func init() {
	// Define flags
	rootCmd.Flags().Int("port", 8080, "Port to run the server on")
	rootCmd.Flags().String("interface", "0.0.0.0", "Interface to bind the server to")

	// Bind flags with Viper
	// Bind flags with Viper and check for errors
	if err := viper.BindPFlag("port", rootCmd.Flags().Lookup("port")); err != nil {
		log.Fatalf("Error binding port flag: %v", err)
	}
	if err := viper.BindPFlag("interface", rootCmd.Flags().Lookup("interface")); err != nil {
		log.Fatalf("Error binding interface flag: %v", err)
	}
	// Enable reading from environment variables
	viper.SetEnvPrefix("IDIONAUTIC") // Use IDIONAUTIC_PORT, IDIONAUTIC_INTERFACE
	viper.AutomaticEnv()

	// Optionally read from a config file
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // Required if using a config file
	viper.AddConfigPath(".")      // Optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file found, using defaults or ENV variables")
	}
}

// Execute starts the root command
func Execute() error {
	return rootCmd.Execute()
}
