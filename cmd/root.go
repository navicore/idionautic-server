package cmd

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
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
		log.Fatal().Err(err).Msg("Failed to bind port flag")
	}
	if err := viper.BindPFlag("interface", rootCmd.Flags().Lookup("interface")); err != nil {
		log.Fatal().Err(err).Msg("Failed to bind interface flag")
	}
	// Enable reading from environment variables
	viper.SetEnvPrefix("IDIONAUTIC") // Use IDIONAUTIC_PORT, IDIONAUTIC_INTERFACE
	viper.AutomaticEnv()

	// Optionally read from a config file
	viper.SetConfigName("idionautic") // Name of config file (without extension)
	viper.SetConfigType("yml")        // Required if using a config file
	viper.AddConfigPath(".")          // Optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		log.Debug().Msg("No config file found, using defaults and ENV variables")
	}

	rootCmd.AddCommand(completionCmd)
}

// Execute starts the root command
func Execute() error {

	log.Debug().Msg("No config file found, using defaults and ENV variables 2*************")
	if err := viper.ReadInConfig(); err != nil {
		log.Debug().Msg("No config file found, using defaults and ENV variables 2")
	}
	return rootCmd.Execute()
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
			log.Fatal().Err(err).Msg("Failed to generate bash completion")
		}
	},
}
