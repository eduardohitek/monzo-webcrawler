package main

import (
	"reflect"
	"testing"
)

func Test_formatURL(t *testing.T) {
	type args struct {
		base string
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1", args: args{base: "http://monzo.com", link: "/about"}, want: "http://monzo.com/about"},
		{name: "Test2", args: args{base: "http://monzo.com", link: "/legal/terms"}, want: "http://monzo.com/legal/terms"},
		{name: "Test3", args: args{base: "http://monzo.com", link: "/features/shared-tabs-more"}, want: "http://monzo.com/features/shared-tabs-more"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatURL(tt.args.base, tt.args.link); got != tt.want {
				t.Errorf("formatURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returnLocalLinks(t *testing.T) {
	type args struct {
		baseURL string
		links   []string
	}
	tests := []struct {
		name           string
		args           args
		wantLocalLinks []string
	}{
		{name: "Test1",
			args: args{
				baseURL: "http://monzo.com",
				links: []string{
					"http://monzo.com/about",
					"http://monzo.com/legal",
					"http://twitter.com/monzo"},
			},
			wantLocalLinks: []string{
				"http://monzo.com/about",
				"http://monzo.com/legal"}},
		{name: "Test2",
			args: args{
				baseURL: "http://monzo.com",
				links: []string{
					"http://instagram.com/monzo",
					"http://facebook.com/monzo",
					"http://twitter.com/monzo"},
			},
			wantLocalLinks: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLocalLinks := returnLocalLinks(tt.args.baseURL, tt.args.links); !reflect.DeepEqual(gotLocalLinks, tt.wantLocalLinks) {
				if !(len(returnLocalLinks(tt.args.baseURL, tt.args.links)) == 0 && len(tt.wantLocalLinks) == 0) {
					t.Errorf("returnLocalLinks() = %v, want %v", gotLocalLinks, tt.wantLocalLinks)
				}
			}
		})
	}
}

func Test_crawl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test1",
			args: args{
				url: "http://mock.eduardohitek.com/"},
			want: []string{
				"http://mock.eduardohitek.com/a.html",
				"http://mock.eduardohitek.com/b.html",
				"http://mock.eduardohitek.com/c.html",
				"http://mock.eduardohitek.com/d.html",
				"http://twitter.com/eduardohitek",
				"http://github.com/eduardohitek",
				"http://mock.eduardohitek.com/e.html",
				"http://mock.eduardohitek.com/f.html",
				"http://mock.eduardohitek.com/g.html"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crawl(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("crawl() = %v, want %v", got, tt.want)
			}
		})
	}
}
