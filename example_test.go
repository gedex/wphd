// Copyright 2016 The wphd Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package wphd_test

import (
	"fmt"

	"github.com/gedex/wphd"
)

func ExampleGetPluginData() {
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

func ExampleGetThemeData() {
	in := `/*
Theme Name: _s
Theme URI: http://underscores.me/
Author: Automattic
Author URI: http://automattic.com/
Description: Hi. I'm a starter theme called <code>_s</code>, or <em>underscores</em>, if you like. I'm a theme meant for hacking so don't use me as a <em>Parent Theme</em>. Instead try turning me into the next, most awesome, WordPress theme out there. That's what I'm here for.
Version: 1.0.0
License: GNU General Public License v2 or later
License URI: http://www.gnu.org/licenses/gpl-2.0.html
Text Domain: _s
Tags:
*/
	`

	t := wphd.GetThemeData([]byte(in))
	fmt.Printf("%+v\n", t)
	// Output:
	// &{Name:_s ThemeURI:http://underscores.me/ Version:1.0.0 Description:Hi. I'm a starter theme called <code>_s</code>, or <em>underscores</em>, if you like. I'm a theme meant for hacking so don't use me as a <em>Parent Theme</em>. Instead try turning me into the next, most awesome, WordPress theme out there. That's what I'm here for. Author:Automattic AuthorURI:http://automattic.com/ Template: Status: Tags:[] TextDomain:_s DomainPath: License:GNU General Public License v2 or later LicenseURI:http://www.gnu.org/licenses/gpl-2.0.html}
}
