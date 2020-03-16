package resource

import (
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/firmeve/firmeve/support/strings"
	"github.com/guregu/null"
	"github.com/ulule/paging"
	"math"
	"net/http"
)

type Paginator struct {
	resolveData contract.ResourceDataCollection
	store       *paging.GORMStore
	request     *http.Request
	pageOption  *paging.Options
	option      *Option
	meta        contract.ResourceMetaData
	link        contract.ResourceLinkData
}

func NewPaginator(store *paging.GORMStore, option *Option, request *http.Request, pageOption *paging.Options) *Paginator {
	return &Paginator{
		store:       store,
		request:     request,
		pageOption:  setDefaultPageOption(pageOption),
		option:      option,
		meta:        make(contract.ResourceMetaData, 0),
		link:        make(contract.ResourceLinkData, 0),
		resolveData: make(contract.ResourceDataCollection, 0),
	}
}

func setDefaultPageOption(option *paging.Options) *paging.Options {
	if option.LimitKeyName == `` {
		option.LimitKeyName = `limit`
	}
	if option.OffsetKeyName == `` {
		option.OffsetKeyName = `offset`
	}

	return option
}

func (p *Paginator) CollectionData() contract.ResourceDataCollection {
	if len(p.resolveData) > 0 {
		return p.resolveData
	}

	paginator, _ := paging.NewOffsetPaginator(p.store, p.request, p.pageOption)

	if err := paginator.Page(); err != nil {
		panic(err)
	}

	p.SetLink(contract.ResourceLinkData{
		"prev":  p.fullUrl(paginator.PreviousURI, paginator.Options),
		"next":  p.fullUrl(paginator.NextURI, paginator.Options),
		"first": p.fullUrl(p.firstUrl(paginator.Limit, paginator.Options), paginator.Options),
		"last":  p.fullUrl(p.lastUrl(paginator.Count, paginator.Limit, paginator.Options), paginator.Options),
	})

	p.SetMeta(contract.ResourceMetaData{
		"current_page": int64(math.Ceil(float64(paginator.Offset)/float64(paginator.Limit)) + 1), //当前页
		"total":        paginator.Count,                                                          //总条数
		"per_page":     paginator.Limit,                                                          //每页多少条
		"from":         paginator.Offset + 1,                                                     //从多少条
		"to":           p.metaTo(paginator.Count, paginator.Limit, paginator.Offset),             //到多少条
		"last_page":    int64(math.Ceil(float64(paginator.Count) / float64(paginator.Limit))),
	})

	p.resolveData = NewCollection(p.store.GetItems(), p.option).CollectionData()
	return p.resolveData
}

func (p *Paginator) SetLink(links contract.ResourceLinkData) {
	p.link = links
}

func (p *Paginator) Link() contract.ResourceLinkData {
	return p.link
}

func (p *Paginator) SetMeta(meta contract.ResourceMetaData) {
	p.meta = meta
}

func (p *Paginator) Meta() contract.ResourceMetaData {
	return p.meta
}

func (p *Paginator) firstUrl(limit int64, options *paging.Options) null.String {
	return null.StringFrom(paging.GenerateOffsetURI(limit, 0, options))
}

func (p *Paginator) lastUrl(count, limit int64, options *paging.Options) null.String {
	return null.StringFrom(paging.GenerateOffsetURI(
		limit,
		count-int64(count%limit), options))
}

func (p *Paginator) metaTo(count, limit, offset int64) int64 {
	total := limit + offset
	if count < total {
		return count
	}

	return total
}

func (p *Paginator) fullUrl(uri null.String, options *paging.Options) string {
	request := p.request
	var protocol string
	if request.TLS != nil || request.Header.Get(`X-Forwarded-Proto`) == `https` {
		protocol = `https://`
	} else {
		protocol = `http://`
	}

	query := request.URL.Query()
	query.Del(options.LimitKeyName)
	query.Del(options.OffsetKeyName)
	var queryString string
	if uri.Valid {
		queryString = strings.Join(``, protocol, request.Host, request.URL.Path, uri.String)
		if len(query) > 0 {
			queryString = strings.Join(`&`, queryString, query.Encode())
		}
	} else {
		queryString = ``
	}

	return queryString
}
