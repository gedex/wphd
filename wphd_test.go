package wphd

import (
	"reflect"
	"testing"
)

func TestGetPluginData(t *testing.T) {
	var tests = []struct {
		in       string
		expected *Plugin
	}{
		{`<?php
/**
 * Plugin Name
 *
 * @package     PluginPackage
 * @author      Your Name
 * @copyright   2016 Your Name or Company Name
 * @license     GPL-2.0+
 *
 * @wordpress-plugin
 * Plugin Name: Plugin Name
 * Plugin URI:  https://example.com/plugin-name
 * Description: Description of the plugin.
 * Version:     1.0.0
 * Author:      Your Name
 * Author URI:  https://example.com
 * Text Domain: plugin-name
 * License:     GPL-2.0+
 * License URI: http://www.gnu.org/licenses/gpl-2.0.txt
 */`,
			&Plugin{
				Name:        "Plugin Name",
				PluginURI:   "https://example.com/plugin-name",
				Description: "Description of the plugin.",
				Version:     "1.0.0",
				Author:      "Your Name",
				AuthorURI:   "https://example.com",
				TextDomain:  "plugin-name",
				License:     "GPL-2.0+",
				LicenseURI:  "http://www.gnu.org/licenses/gpl-2.0.txt",
			},
		},
		{`<?php
/*
 * Plugin Name: Jetpack by WordPress.com
 * Plugin URI: http://jetpack.com
 * Description: Bring the power of the WordPress.com cloud to your self-hosted WordPress. Jetpack enables you to connect your blog to a WordPress.com account to use the powerful features normally only available to WordPress.com users.
 * Author: Automattic
 * Version: 4.4-alpha
 * Author URI: http://jetpack.com
 * License: GPL2+
 * Text Domain: jetpack
 * Domain Path: /languages/
 */

define( 'JETPACK__MINIMUM_WP_VERSION', '4.5' );`,
			&Plugin{
				Name:        "Jetpack by WordPress.com",
				PluginURI:   "http://jetpack.com",
				Description: "Bring the power of the WordPress.com cloud to your self-hosted WordPress. Jetpack enables you to connect your blog to a WordPress.com account to use the powerful features normally only available to WordPress.com users.",
				Author:      "Automattic",
				Version:     "4.4-alpha",
				AuthorURI:   "http://jetpack.com",
				License:     "GPL2+",
				TextDomain:  "jetpack",
				DomainPath:  "/languages/",
			},
		},
	}

	for _, test := range tests {
		actual := GetPluginData([]byte(test.in))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Got %+v, want %+v", actual, test.expected)
		}
	}
}

func TestGetThemeData(t *testing.T) {
	var tests = []struct {
		in       string
		expected *Theme
	}{
		{`/*
Theme Name: Twenty Thirteen
Theme URI: http://wordpress.org/themes/twentythirteen
Author: the WordPress team
Author URI: http://wordpress.org/
Description: The 2013 theme for WordPress takes us back to the blog, featuring a full range of post formats, each displayed beautifully in their own unique way. Design details abound, starting with a vibrant color scheme and matching header images, beautiful typography and icons, and a flexible layout that looks great on any device, big or small.
Version: 1.0
License: GNU General Public License v2 or later
License URI: http://www.gnu.org/licenses/gpl-2.0.html
Tags: black, brown, orange, tan, white, yellow, light, one-column, two-columns, right-sidebar, flexible-width, custom-header, custom-menu, editor-style, featured-images, microformats, post-formats, rtl-language-support, sticky-post, translation-ready
Text Domain: twentythirteen

This theme, like WordPress, is licensed under the GPL.
Use it to make something cool, have fun, and share what you've learned with others.
*/`,
			&Theme{
				Name:        "Twenty Thirteen",
				ThemeURI:    "http://wordpress.org/themes/twentythirteen",
				Author:      "the WordPress team",
				AuthorURI:   "http://wordpress.org/",
				Description: "The 2013 theme for WordPress takes us back to the blog, featuring a full range of post formats, each displayed beautifully in their own unique way. Design details abound, starting with a vibrant color scheme and matching header images, beautiful typography and icons, and a flexible layout that looks great on any device, big or small.",
				Version:     "1.0",
				License:     "GNU General Public License v2 or later",
				LicenseURI:  "http://www.gnu.org/licenses/gpl-2.0.html",
				Tags:        []string{"black", "brown", "orange", "tan", "white", "yellow", "light", "one-column", "two-columns", "right-sidebar", "flexible-width", "custom-header", "custom-menu", "editor-style", "featured-images", "microformats", "post-formats", "rtl-language-support", "sticky-post", "translation-ready"},
				TextDomain:  "twentythirteen",
			},
		},
	}

	for _, test := range tests {
		actual := GetThemeData([]byte(test.in))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Got %+v, want %+v", actual, test.expected)
		}
	}
}
