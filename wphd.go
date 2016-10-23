// Copyright 2016 The wphd Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package wphd (WordPress Header Data) parses header data from plugin and
// theme file headers
package wphd

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Plugin represents plugin metadata.
type Plugin struct {
	Name        string `wphd:"Plugin Name"`
	PluginURI   string `wphd:"Plugin URI"`
	Version     string `wphd:"Version"`
	Description string `wphd:"Description"`
	Author      string `wphd:"Author"`
	AuthorURI   string `wphd:"Author URI"`
	TextDomain  string `wphd:"Text Domain"`
	DomainPath  string `wphd:"Domain Path"`
	Network     bool   `wphd:"Network"`
	License     string `wphd:"License"`
	LicenseURI  string `wphd:"License URI"`
}

// set sets Plugin's field with given value.
func (p *Plugin) set(field, value string) {
	switch field {
	case "Name":
		p.Name = value
	case "PluginURI":
		p.PluginURI = value
	case "Version":
		p.Version = value
	case "Description":
		p.Description = value
	case "Author":
		p.Author = value
	case "AuthorURI":
		p.AuthorURI = value
	case "TextDomain":
		p.TextDomain = value
	case "DomainPath":
		p.DomainPath = value
	case "Network":
		p.Network = strings.ToLower(value) == "true"
	case "License":
		p.License = value
	case "LicenseURI":
		p.LicenseURI = value
	}
}

// Theme represents Theme metadata.
type Theme struct {
	Name        string   `wphd:"Theme Name"`
	ThemeURI    string   `wphd:"Theme URI"`
	Version     string   `wphd:"Version"`
	Description string   `wphd:"Description"`
	Author      string   `wphd:"Author"`
	AuthorURI   string   `wphd:"Author URI"`
	Template    string   `wphd:"Template"`
	Status      string   `wphd:"Status"`
	Tags        []string `wphd:"Tags"`
	TextDomain  string   `wphd:"Text Domain"`
	DomainPath  string   `wphd:"Domain Path"`
	License     string   `wphd:"License"`
	LicenseURI  string   `wphd:"License URI"`
}

// set sets Theme's field with given value.
func (t *Theme) set(field, value string) {
	switch field {
	case "Name":
		t.Name = value
	case "ThemeURI":
		t.ThemeURI = value
	case "Version":
		t.Version = value
	case "Description":
		t.Description = value
	case "Author":
		t.Author = value
	case "AuthorURI":
		t.AuthorURI = value
	case "Template":
		t.Template = value
	case "Status":
		t.Status = value
	case "Tags":
		tags := strings.Split(value, ",")
		for i, v := range tags {
			tags[i] = strings.TrimSpace(v)
		}
		t.Tags = tags
	case "TextDomain":
		t.TextDomain = value
	case "DomainPath":
		t.DomainPath = value
	case "License":
		t.License = value
	case "LicenseURI":
		t.LicenseURI = value
	}
}

var (
	crRe    *regexp.Regexp            // CR-only regex
	pfRe    map[string]*regexp.Regexp // Plugin fields regex
	tfRe    map[string]*regexp.Regexp // Theme fields regex
	cleanRe *regexp.Regexp            // Clean fields metadata regex
)

func init() {
	crRe = regexp.MustCompile(`\r`)

	p := reflect.TypeOf(Plugin{})
	pfRe = make(map[string]*regexp.Regexp)
	for i := 0; i < p.NumField(); i++ {
		pfRe[p.Field(i).Name] = regexp.MustCompile(fmt.Sprintf(`(?mi)^[ \t\/*#@]*%s:(.*)$`, p.Field(i).Tag.Get("wphd")))
	}

	t := reflect.TypeOf(Theme{})
	tfRe = make(map[string]*regexp.Regexp)
	for i := 0; i < t.NumField(); i++ {
		tfRe[t.Field(i).Name] = regexp.MustCompile(fmt.Sprintf(`(?mi)^[ \t\/*#@]*%s:(.*)$`, t.Field(i).Tag.Get("wphd")))
	}

	cleanRe = regexp.MustCompile(`\s*(?:\*\/|\?>).*`)
}

// cleanup returns clean string of found field data. Strips close comment and
// close PHP tags.
func cleanup(in string) string {
	clean := cleanRe.ReplaceAllString(in, "")
	clean = strings.TrimSpace(clean)

	return clean
}

// GetPluginData parses the plugin header from plugin's main file to retrieve
// plugin's metadata.
func GetPluginData(src []byte) *Plugin {
	src = crRe.ReplaceAll(src, []byte(""))
	p := &Plugin{}
	for field, re := range pfRe {
		sub := re.FindSubmatch(src)
		if len(sub) >= 2 {
			p.set(field, cleanup(string(sub[1])))
		}
	}

	return p
}

// GetThemeData parses the theme header from theme's style.css to retrieve
// theme's metadata.
func GetThemeData(src []byte) *Theme {
	src = crRe.ReplaceAll(src, []byte(""))
	t := &Theme{}
	for field, re := range tfRe {
		sub := re.FindSubmatch(src)
		if len(sub) >= 2 {
			t.set(field, cleanup(string(sub[1])))
		}
	}

	return t
}
