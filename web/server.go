package web

import (
	"context"
	"fmt"
	"github.com/raedahgroup/dcrextdata/netsnapshot"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/raedahgroup/dcrextdata/commstats"
	"github.com/raedahgroup/dcrextdata/exchanges/ticks"
	"github.com/raedahgroup/dcrextdata/mempool"
	"github.com/raedahgroup/dcrextdata/postgres/models"
	"github.com/raedahgroup/dcrextdata/pow"
	"github.com/raedahgroup/dcrextdata/vsp"
)

type DataQuery interface {
	ExchangeTickCount(ctx context.Context) (int64, error)
	AllExchangeTicks(ctx context.Context, currencyPair string, defaultInterval, offset, limit int) ([]ticks.TickDto, int64, error)
	AllExchange(ctx context.Context) (models.ExchangeSlice, error)
	FetchExchangeTicks(ctx context.Context, currencyPair, name string, defaultInterval, offset, limit int) ([]ticks.TickDto, int64, error)
	AllExchangeTicksCurrencyPair(ctx context.Context) ([]ticks.TickDtoCP, error)
	ExchangeTicksChartData(ctx context.Context, filter string, currencyPair string, selectedInterval int, exchanges string) ([]ticks.TickChartData, error)
	AllExchangeTicksInterval(ctx context.Context) ([]ticks.TickDtoInterval, error)

	VspTickCount(ctx context.Context) (int64, error)
	FetchVSPs(ctx context.Context) ([]vsp.VSPDto, error)
	FiltredVSPTicks(ctx context.Context, vspName string, offset, limit int) ([]vsp.VSPTickDto, int64, error)
	AllVSPTicks(ctx context.Context, offset, limit int) ([]vsp.VSPTickDto, int64, error)
	FetchChartData(ctx context.Context, attribute, vspName string) (records []vsp.ChartData, err error)
	GetVspTickDistinctDates(ctx context.Context, vsps []string) ([]time.Time, error)

	PowCount(ctx context.Context) (int64, error)
	FetchPowData(ctx context.Context, offset, limit int) ([]pow.PowDataDto, int64, error)
	FetchPowDataBySource(ctx context.Context, source string, offset, limit int) ([]pow.PowDataDto, int64, error)
	FetchPowSourceData(ctx context.Context) ([]pow.PowDataSource, error)
	FetchPowChartData(ctx context.Context, pool string, dataType string) ([]pow.PowChartData, error)
	GetPowDistinctDates(ctx context.Context, vsps []string) ([]time.Time, error)

	MempoolCount(ctx context.Context) (int64, error)
	Mempools(ctx context.Context, offtset int, limit int) ([]mempool.MempoolDto, error)
	MempoolsChartData(ctx context.Context, chartFilter string) (models.MempoolSlice, error)

	BlockCount(ctx context.Context) (int64, error)
	Blocks(ctx context.Context, offset int, limit int) ([]mempool.BlockDto, error)
	BlockHeights(ctx context.Context) ([]int64, error)
	BlocksWithoutVotes(ctx context.Context, offset int, limit int) ([]mempool.BlockDto, error)

	Votes(ctx context.Context, offset int, limit int) ([]mempool.VoteDto, error)
	VotesCount(ctx context.Context) (int64, error)
	PropagationVoteChartData(ctx context.Context) ([]mempool.PropagationChartData, error)
	PropagationBlockChartData(ctx context.Context) ([]mempool.PropagationChartData, error)
	FetchBlockReceiveTime(ctx context.Context) ([]mempool.BlockReceiveTime, error)

	CountRedditStat(ctx context.Context, subreddit string) (int64, error)
	RedditStats(ctx context.Context, subreddit string, offtset int, limit int) ([]commstats.Reddit, error)
	CountTwitterStat(ctx context.Context, handle string) (int64, error)
	TwitterStats(ctx context.Context, handle string, offtset int, limit int) ([]commstats.Twitter, error)
	CountYoutubeStat(ctx context.Context, channel string) (int64, error)
	YoutubeStat(ctx context.Context, channel string, offtset int, limit int) ([]commstats.Youtube, error)
	CountGithubStat(ctx context.Context, repository string) (int64, error)
	GithubStat(ctx context.Context, repository string, offtset int, limit int) ([]commstats.Github, error)
	CommunityChart(ctx context.Context, platform string, dataType string, filters map[string]string) ([]commstats.ChartData, error)

	LastSnapshotTime(ctx context.Context) (timestamp int64)
	FindNetworkSnapshot(ctx context.Context, timestamp int64) (*netsnapshot.SnapShot, error)
	PreviousSnapshot(ctx context.Context, timestamp int64) (*netsnapshot.SnapShot, error)
	NextSnapshot(ctx context.Context, timestamp int64) (*netsnapshot.SnapShot, error)
	TotalPeerCount(ctx context.Context, timestamp int64) (int64, error)
	NetworkPeers(ctx context.Context, timestamp int64, q string, offset int, limit int) ([]netsnapshot.NetworkPeer, int64, error)
	LastSnapshotForNode(ctx context.Context, address string) (*netsnapshot.NetworkPeer, error)
	TotalPeerCountByProtocol(ctx context.Context, timestamp int64, protocolVersion int) (int64, error)
	PeerCountByUserAgents(ctx context.Context, timestamp int64) (userAgents []netsnapshot.UserAgentInfo, err error)
	PeerCountByIPVersion(ctx context.Context, timestamp int64, iPVersion int) (int64, error)
	PeerCountByCountries(ctx context.Context, timestamp int64) (countries []netsnapshot.CountryInfo, err error) 
}

type Server struct {
	templates    map[string]*template.Template
	lock         sync.RWMutex
	db           DataQuery
	extDbFactory func(name string) (DataQuery, error)
}

func StartHttpServer(httpHost, httpPort string, db DataQuery, extDbFactory func(name string) (DataQuery, error)) {
	server := &Server{
		templates:    map[string]*template.Template{},
		db:           db,
		extDbFactory: extDbFactory,
	}

	router := chi.NewRouter()
	workDir, _ := os.Getwd()

	filesDir := filepath.Join(workDir, "web/public/dist")
	FileServer(router, "/static", http.Dir(filesDir))
	server.registerHandlers(router)

	// load templates
	server.loadTemplates()

	address := net.JoinHostPort(httpHost, httpPort)

	log.Infof("Starting http server on %s", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Error("Error starting web server")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func (s *Server) registerHandlers(r *chi.Mux) {
	r.Get("/", s.homePage)
	r.Get("/exchanges", s.getExchangeTicks)
	r.Get("/exchangedata", s.getFilteredExchangeTicks)
	r.Get("/exchangechart", s.getExchangeChartData)
	r.Get("/vsp", s.getVspTicks)
	r.Get("/vspchartdata", s.vspChartData)
	r.Get("/vsps", s.getFilteredVspTicks)
	r.Get("/pow", s.powPage)
	r.Get("/filteredpow", s.getFilteredPowData)
	r.Get("/powchart", s.getPowChartData)
	r.Get("/mempool", s.mempoolPage)
	r.Get("/mempoolcharts", s.getMempoolChartData)
	r.Get("/getmempool", s.getMempool)
	r.Get("/propagation", s.propagation)
	r.Get("/getpropagationdata", s.getPropagationData)
	r.Get("/blockschartdata", s.blocksChartData)
	r.Get("/voteschartdata", s.votesChartDate)
	r.Get("/propagationchartdata", s.propagationChartData)
	r.Get("/getblocks", s.getBlocks)
	r.Get("/blockdata", s.getBlockData)
	r.Get("/getvotes", s.getVotes)
	r.Get("/votesdata", s.getVoteData)

	r.Get("/community", s.community)
	r.Get("/getCommunityStat", s.getCommunityStat)
	r.Get("/communitychat", s.communityChat)

	r.With().Get("/snapshot", s.snapshot)
	r.With(addTimestampToCtx).Get("/snapshot/{timestamp}", s.snapshot)
	r.With(addTimestampToCtx).Get("/api/snapshot/{timestamp}/nodes", s.nodes)
	r.With(addTimestampToCtx).Get("/api/snapshot/{timestamp}/user-agents", s.nodesCountbUserAgents)
	r.With(addTimestampToCtx).Get("/api/snapshot/{timestamp}/countries", s.nodesCountByCountries)

	r.With(syncDataType).Get("/api/sync/{dataType}", s.sync)
}
