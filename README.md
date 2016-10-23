wphd
====

> WordPress [Header Data](https://codex.wordpress.org/File_Header) parser.

The wphd package allows Go programs to retrieve header data from plugin and 
theme file headers.

[![Build Status](https://travis-ci.org/gedex/wphd.png?branch=master)](https://travis-ci.org/gedex/wphd)
[![Coverage Status](https://coveralls.io/repos/gedex/wphd/badge.png?branch=master)](https://coveralls.io/r/gedex/wphd?branch=master)
[![GoDoc](https://godoc.org/github.com/gedex/wphd?status.svg)](https://godoc.org/github.com/gedex/wphd)

## Example

~~~go
package main

import (
	"fmt"

	"github.com/gedex/wphd"
)

func main() {
	in := `<?php
/**
 * @package Hello_Dolly
 * @version 1.6
 */
/*
Plugin Name: Hello Dolly
Plugin URI: https://wordpress.org/plugins/hello-dolly/
Description: This is not just a plugin, it symbolizes the hope and enthusiasm of an entire generation summed up in two words sung most famously by Louis Armstrong: Hello, Dolly. When activated you will randomly see a lyric from <cite>Hello, Dolly</cite> in the upper right of your admin screen on every page.
Author: Matt Mullenweg
Version: 1.6
Author URI: http://ma.tt/
*/`
	p := wphd.GetPluginData([]byte(in))
	fmt.Printf("%+v\n", p)
	// Output:
	// &{Name:Hello Dolly PluginURI:https://wordpress.org/plugins/hello-dolly/ Version:1.6 Description:This is not just a plugin, it symbolizes the hope and enthusiasm of an entire generation summed up in two words sung most famously by Louis Armstrong: Hello, Dolly. When activated you will randomly see a lyric from <cite>Hello, Dolly</cite> in the upper right of your admin screen on every page. Author:Matt Mullenweg AuthorURI:http://ma.tt/ TextDomain: DomainPath: Network:false License: LicenseURI:}
}
~~~

## Known Limitations

* Return the raw value for each field
* No translation for each field
* Markup is not applied for each field
* Extra fields are not considered
