package web

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func (s *Server) loadTemplates() {
	layout := "web/views/layout.html"
	tpls := map[string]string{
		"error.html":       "web/views/error.html",
		"home.html":        "web/views/home.html",
		"exchange.html":    "web/views/exchange.html",
		"vsp.html":         "web/views/vsp.html",
		"pow.html":         "web/views/pow.html",
		"mempool.html":     "web/views/mempool.html",
		"propagation.html": "web/views/propagation.html",
		"community.html":   "web/views/community.html",
		"nodes.html":       "web/views/nodes.html",
		"node.html":       "web/views/node.html",
	}

	for i, v := range tpls {
		tpl, err := template.New(i).Funcs(templateFuncMap()).ParseFiles(v, layout)
		if err != nil {
			log.Errorf("Error loading templates: %s", err.Error())
		}

		s.lock.Lock()
		s.templates[i] = tpl
		s.lock.Unlock()
	}
}

var pairMap = map[string]string{
	"BTC/DCR": "DCR/BTC",
	"USD/BTC": "BTC/USD",
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"incByOne": func(number int) int {
			return number + 1
		},
		"formatDate": func(date time.Time) string {
			return date.Format("2006-01-02 15:04")
		},
		"formatDateMilli": func(date time.Time) string {
			return date.Format("2006-01-02 15:04:05.99")
		},
		"normalizeBalance": func(balance float64) string {
			return fmt.Sprintf("%010.8f DCR", balance)
		},
		"timestamp": func() int64 {
			return time.Now().Unix()
		},
		"timeSince": func(timestamp int64) string {
			return time.Since(time.Unix(timestamp, 0).UTC()).String()
		},
		"formatUnixTime": func(timestamp int64) string {
			return time.Unix(timestamp, 0).Format(time.UnixDate)
		},
		"unixTimeAgo": func(timestamp int64) string {
			return time.Since(time.Unix(timestamp, 0)).String()
		},
		"strListContains": func(stringList []string, needle string) bool {
			for _, value := range stringList {
				if value == needle {
					return true
				}
			}
			return false
		},
		"stringsReplace": func(input string, old string, new string) string {
			return strings.Replace(input, old, new, -1)
		},
		"humanizeInt": func(number int64) string {
			s := strconv.Itoa(int(number))
			r1 := ""
			idx := 0

			// Reverse and interleave the separator.
			for i := len(s) - 1; i >= 0; i-- {
				idx++
				if idx == 4 {
					idx = 1
					r1 = r1 + ","
				}
				r1 = r1 + string(s[i])
			}

			// Reverse back and return.
			r2 := ""
			for i := len(r1) - 1; i >= 0; i-- {
				r2 = r2 + string(r1[i])
			}
			return r2
		},
		"commonPair": func(pair string) string {
			if v, f := pairMap[pair]; f {
				return v
			}
			return pair
		},
		"percentage": func(total int64, actual int64) float64 {
			return 100 * float64(actual)/float64(total)
		},
	}
}
