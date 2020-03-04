// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"net"
	"os"

	flags "github.com/jessevdk/go-flags"
)

const (
	DefaultConfigFilename      = "dcrextdata.conf"
	defaultLogFilename         = "dcrextdata.log"
	Hint                       = `Run dcrextdata < --http > to start http server or dcrextdata < --help > for help.`
	defaultDbHost              = "localhost"
	defaultDbPort              = "5432"
	defaultDbUser              = "postgres"
	defaultDbPass              = "dbpass"
	defaultDbName              = "dcrextdata"
	defaultLogLevel            = "debug"
	defaultHttpHost            = "127.0.0.1"
	defaultHttpPort            = "7770"
	defaultDcrdServer          = "127.0.0.1:9109"
	defaultDcrdUser            = "rpcuser"
	defaultDcrdPassword        = "rpcpass"
	defaultDcrdNetworkType     = "mainnet"
	defaultMempoolInterval     = 60
	defaultVSPInterval         = 300
	defaultPowInterval         = 300
	defaultSyncInterval        = 60
	defaultSnapshotInterval    = 5
	defaultRedditInterval      = 60
	defaultTwitterStatInterval = 60 * 24
	defaultGithubStatInterval  = 60 * 24
	defaultYoutubeInterval     = 60 * 24
	//dcrseeder
	defaultSeederHostAddress   = "network-seed.example.com"
	defaultSeederListonAddress = "localhost"
	defaultSeederListenPort    = "5354"
	defaultSeederNameServer    = "nameserver.example.com"
	defaultSeeder              = "127.0.0.1"
)

var (
	defaultSubreddits          = []string{"decred"}
	defaultTwitterHandles      = []string{"decredproject"}
	defaultGithubRepositories  = []string{"decred/dcrd", "decred/dcrdata", "decred/dcrwallet", "decred/politeia", "decred/decrediton"}
	defaultYoutubeChannelNames = []string{"Decred"}
	defaultYoutubeChannelId    = []string{"UCJ2bYDaPYHpSmJPh_M5dNSg"}
)

func defaultFileOptions() ConfigFileOptions {
	cfg := ConfigFileOptions{
		LogFile:         defaultLogFilename,
		DBHost:          defaultDbHost,
		DBPort:          defaultDbPort,
		DBUser:          defaultDbUser,
		DBPass:          defaultDbPass,
		DBName:          defaultDbName,
		DebugLevel:      defaultLogLevel,
		VSPInterval:     defaultVSPInterval,
		PowInterval:     defaultPowInterval,
		MempoolInterval: defaultMempoolInterval,
		DcrdNetworkType: defaultDcrdNetworkType,
		DcrdRpcServer:   defaultDcrdServer,
		DcrdRpcUser:     defaultDcrdUser,
		DcrdRpcPassword: defaultDcrdPassword,
		HTTPHost:        defaultHttpHost,
		HTTPPort:        defaultHttpPort,
		SyncInterval:    defaultSyncInterval,
	}

	cfg.RedditStatInterval = defaultRedditInterval
	cfg.Subreddit = defaultSubreddits
	cfg.TwitterStatInterval = defaultTwitterStatInterval
	cfg.TwitterHandles = defaultTwitterHandles
	cfg.GithubStatInterval = defaultGithubStatInterval
	cfg.GithubRepositories = defaultGithubRepositories
	cfg.YoutubeStatInterval = defaultYoutubeInterval
	cfg.YoutubeChannelName = defaultYoutubeChannelNames
	cfg.YoutubeChannelId = defaultYoutubeChannelId
	cfg.SnapshotInterval = defaultSnapshotInterval
	cfg.SeederHost = defaultSeederHostAddress
	cfg.Nameserver = defaultSeederNameServer
	cfg.Listen = normalizeAddress(defaultSeederListonAddress, defaultSeederListenPort)
	cfg.Seeder = defaultSeeder

	return cfg
}

type Config struct {
	ConfigFileOptions
	CommandLineOptions
}

type ConfigFileOptions struct {
	// General application behaviour
	LogFile    string `short:"L" long:"logfile" description:"File name of the log file"`
	DebugLevel string `short:"d" long:"debuglevel" description:"Logging level {trace, debug, info, warn, error, critical}"`
	Quiet      bool   `short:"q" long:"quiet" description:"Easy way to set debuglevel to error"`

	// Postgresql Configuration
	DBHost string `long:"dbhost" description:"Database host"`
	DBPort string `long:"dbport" description:"Database port"`
	DBUser string `long:"dbuser" description:"Database username"`
	DBPass string `long:"dbpass" description:"Database password"`
	DBName string `long:"dbname" description:"Database name"`

	// Http Server
	HTTPHost string `long:"httphost" description:"HTTP server host address or IP when running godcr in http mode."`
	HTTPPort string `long:"httpport" description:"HTTP server port when running godcr in http mode."`

	// pprof
	Cpuprofile string `long:"cpuprofile" description:"write cpu profile to file"`
	Memprofile string `long:"memprofile" description:"write memory profile to file"`

	// Exchange collector
	DisableExchangeTicks bool     `long:"disablexcticks" description:"Disables collection of ticker data from exchanges"`
	DisabledExchanges    []string `long:"disableexchange" description:"Disable data collection for this exchange"`

	// PoW collector
	DisablePow   bool     `long:"disablepow" description:"Disables collection of data for pows"`
	DisabledPows []string `long:"disabledpow" description:"Disable data collection for this Pow"`
	PowInterval  int64    `long:"powI" description:"Collection interval for Pow"`

	// VSP
	DisableVSP  bool  `long:"disablevsp" description:"Disables periodic voting service pool status collection"`
	VSPInterval int64 `long:"vspinterval" description:"Collection interval for pool status collection"`

	// Mempool
	DisableMempool  bool    `long:"disablemempool" description:"Disable mempool data collection"`
	MempoolInterval float64 `long:"mempoolinterval" description:"The duration of time between mempool collection"`
	DcrdRpcServer   string  `long:"dcrdrpcserver" description:"Dcrd rpc server host"`
	DcrdNetworkType string  `long:"dcrdnetworktype" description:"Dcrd rpc network type"`
	DcrdRpcUser     string  `long:"dcrdrpcuser" description:"Your Dcrd rpc username"`
	DcrdRpcPassword string  `long:"dcrdrpcpassword" description:"Your Dcrd rpc password"`
	DisableTLS		bool	`long:"dcrdisabletls" description:"DisableTLS specifies whether transport layer security should be disabled"`

	// sync
	DisableSync   bool     `long:"disablesync" description:"Disables data sharing operation"`
	SyncInterval  int      `long:"syncinterval" description:"The number of minuets between sync operations"`
	SyncSources   []string `long:"syncsource" description:"Address of remote instance to sync data from"`
	SyncDatabases []string `long:"syncdatabase" description:"Database to sync remote data to"`

	CommunityStatOptions   
	NetworkSnapshotOptions 
}

// CommandLineOptions holds the top-level options/flags that are displayed on the command-line menu
type CommandLineOptions struct {
	Reset      bool   `short:"R" long:"reset" description:"Drop all database tables and start over"`
	ConfigFile string `short:"C" long:"configfile" description:"Path to Configuration file"`
	HttpMode   bool   `long:"http" description:"Launch http server"`
}

type CommunityStatOptions struct {
	// Community stat
	DisableCommunityStat bool     `long:"disablecommstat" description:"Disables periodic community stat collection"`
	RedditStatInterval   int64    `long:"redditstatinterval" description:"Collection interval for Reddit community stat"`
	Subreddit            []string `long:"subreddit" description:"List of subreddit for community stat collection"`
	TwitterHandles       []string `long:"twitterhandle" description:"List of twitter handles community stat collection"`
	TwitterStatInterval  int      `long:"twitterstatinterval" description:"Number of minutes between Twitter stat collection"`
	GithubRepositories   []string `long:"githubrepository" description:"List of Github repositories to track"`
	GithubStatInterval   int      `long:"githubstatinterval" description:"Number of minutes between Github stat collection"`
	YoutubeChannelName   []string `long:"youtubechannelname" description:"List of Youtube channel names to be tracked"`
	YoutubeChannelId     []string `long:"youtubechannelid" description:"List of Youtube channel ID to be tracked"`
	YoutubeStatInterval  int      `long:"youtubestatinterval" description:"Number of minutes between Youtube stat collection"`
	YoutubeDataApiKey    string   `long:"youtubedataapikey" description:"Youtube data API key gotten from google developer console"`
}

type NetworkSnapshotOptions struct {
	DisableNetworkSnapshot bool   `long:"disablesnapshot" description:"Disable network snapshot"`
	SnapshotInterval       int    `long:"snapshotinterval" description:"The number of minutes between snapshot (default 5)"`
	SeederHost             string `short:"H" long:"host" description:"Seed DNS address"`
	Listen                 string `long:"listen" short:"l" description:"Listen on address:port"`
	Nameserver             string `short:"n" long:"nameserver" description:"hostname of nameserver"`
	Seeder                 string `short:"s" long:"seeder" description:"IP address of a working node"`
	SeederPort             string `short:"p" long:"seederport" description:"Port of a working node"`
	TestNet                bool   `long:"testnet" description:"Use testnet"`
	ShowDetailedLog        bool   `long:"showdetailedlog" description:"Weather or not to show detailed log for peer discovery"`
}

func defaultConfig() Config {
	return Config{
		CommandLineOptions: CommandLineOptions{
			ConfigFile: DefaultConfigFilename,
		},
		ConfigFileOptions: defaultFileOptions(),
	}
}

func LoadConfig() (*Config, []string, error) {
	cfg := defaultConfig()

	// Pre-parse the command line options to see if an alternative config file
	// or the version flag was specified. Override any environment variables
	// with parsed command line flags.
	preCfg := cfg
	preParser := flags.NewParser(&preCfg, flags.HelpFlag|flags.PassDoubleDash)
	_, flagerr := preParser.Parse()

	if flagerr != nil {
		e, ok := flagerr.(*flags.Error)
		if ok && e.Type != flags.ErrHelp {
			return nil, nil, flagerr
		}
	}

	parser := flags.NewParser(&cfg, flags.IgnoreUnknown)
	err := flags.NewIniParser(parser).ParseFile(preCfg.ConfigFile)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			fmt.Printf("Missing Config file %s in current directory\n", preCfg.ConfigFile)
		} else {
			return nil, nil, err
		}
	}

	unknownArg, err := parser.Parse()
	if err != nil {
		e, ok := err.(*flags.Error)
		if ok && e.Type == flags.ErrHelp {
			os.Exit(0)
		}
		return nil, nil, err
	}

	// network snapshot validation

	if len(cfg.SeederHost) == 0 {
		return nil, nil, fmt.Errorf("Please specify a hostname")
	}

	if len(cfg.Nameserver) == 0 {
		return nil, nil, fmt.Errorf("Please specify a nameserver")
	}

	if len(cfg.Seeder) == 0 {
		return nil, nil, fmt.Errorf("Please specify a seeder")
	}

	if net.ParseIP(cfg.Seeder) == nil {
		str := "\"%s\" is not a valid textual representation of an IP address"
		return nil, nil, fmt.Errorf(str, cfg.Seeder)
	}

	cfg.Listen = normalizeAddress(cfg.Listen, defaultSeederListenPort)

	return &cfg, unknownArg, nil
}

// normalizeAddress returns addr with the passed default port appended if
// there is not already a port specified.
func normalizeAddress(addr, defaultPort string) string {
	_, _, err := net.SplitHostPort(addr)
	if err != nil {
		return net.JoinHostPort(addr, defaultPort)
	}
	return addr
}
