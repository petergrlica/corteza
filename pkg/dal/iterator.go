package dal

import (
	"context"
	"encoding/json"
	"io"

	"github.com/cortezaproject/corteza-server/pkg/filter"
)

type (
	// Iterator provides an interface for loading data from the underlying store
	Iterator interface {
		Next(ctx context.Context) bool
		More(uint, ValueGetter) error
		Err() error
		Scan(ValueSetter) error
		Close() error

		BackCursor(ValueGetter) (*filter.PagingCursor, error)
		ForwardCursor(ValueGetter) (*filter.PagingCursor, error)

		// // -1 means unknown
		// Total() int
		// Cursor() any
		// // ... do we need anything else here?
	}
)

// IteratorEncodeJSON helper function that encodes each item from the iterator as JSON
// and writes it to th given io.Writer.
//
// target initialization function is intentionally used to avoid use of reflection
func IteratorEncodeJSON(ctx context.Context, w io.Writer, iter Iterator, initTarget func() ValueSetter) (err error) {
	var (
		target   ValueSetter
		firstOut = false
	)

	for iter.Next(ctx) {
		if err = iter.Err(); err != nil {
			return
		}

		if firstOut {
			if _, err = w.Write([]byte(`,`)); err != nil {
				return
			}
		}

		firstOut = true

		target = initTarget()

		if err = iter.Scan(target); err != nil {
			return
		}

		err = json.NewEncoder(w).Encode(target)
		if err != nil {
			return
		}
	}

	return
}

// IteratorPaging helper function for record paging cursor and total
func IteratorPaging(ctx context.Context, iter Iterator, infp filter.Paging, fn func(i Iterator) (ValueGetter, bool)) (out filter.Paging, err error) {
	// @todo: temp fix
	// it was breaking due to paging was not properly cloned
	if val := infp.Clone(); val != nil {
		out = *val
	}

	out.PageNavigation = []*filter.Page{}

	const howMuchMore = 1000

	var (
		counter uint
		total   uint

		// when calculating totals or when page navigation needs to be included
		// we need to fetch ALL records
		fetchAll = infp.IncTotal || infp.IncPageNavigation

		cur  *filter.PagingCursor
		page = filter.Page{
			Page:   1,
			Count:  infp.Limit,
			Cursor: nil,
		}
	)

	for iter.Next(ctx) {
		if err = iter.Err(); err != nil {
			return
		}

		// callback (access-control)
		//
		// if user can not read the record
		// we should not include it in the total count
		// (and adjust page navigation as well)
		r, ok := fn(iter)
		if !ok {
			continue
		}

		counter++
		total++
		page.Count++
		hasNextPage := infp.Limit > 0 && total%infp.Limit == 0

		if infp.PageCursor != nil && out.PrevPage == nil {
			// when input has a cursor + we do not have previous-page-cursor
			// we need to generate one
			out.PrevPage, err = iter.BackCursor(r)
			if err != nil {
				return
			}
		}

		// cursor for each page
		cur, err = iter.ForwardCursor(r)
		if err != nil {
			return
		}

		// We fetched enough, we don't need count/all pages; end because anything
		// extra would be useless processing
		if hasNextPage && !fetchAll {
			// @todo why are we doing this here?
			//       next-page cursor is not created if we break too early
			// break
		}

		// Additional processing only happens when we
		// get to the next page so that we can safely skip it
		if !hasNextPage {
			continue
		}

		// Update paging params for the initial filtering
		if out.NextPage == nil {
			out.NextPage = cur
		}

		// Paging params for the current chunk
		out.PageNavigation = append(out.PageNavigation, &filter.Page{
			Page:   page.Page,
			Count:  page.Count,
			Cursor: page.Cursor,
		})

		// Prepare paging params for the next chunk
		page = filter.Page{
			Page:   uint(len(out.PageNavigation) + 1),
			Count:  0,
			Cursor: cur,
		}

		// Request more
		// If this was the first page, request more because the limit was exceeded
		// If this wasn't the first page, request more after we've reached the re-fetch count
		if len(out.PageNavigation) == 1 || counter == howMuchMore {
			counter = 0

			// request more items
			if err = iter.More(howMuchMore, r); err != nil {
				return
			}
		}
	}

	// push the last page to page navigation
	if page.Count > 0 {
		out.PageNavigation = append(out.PageNavigation, &filter.Page{
			Page:   page.Page,
			Count:  page.Count,
			Cursor: page.Cursor,
		})
	}

	if total < infp.Limit {
		out.NextPage = nil
	}

	if infp.IncTotal {
		// @todo cal based on pages
		out.Total = total
	}

	if !infp.IncPageNavigation {
		out.PageNavigation = nil
	}

	return
}
