package png

import (
	"image/color"
	"strconv"
	"strings"
)

func getBool(s string, def bool) bool {
	if s == "" {
		return def
	}

	switch s {
	case "True", "true", "1":
		return true
	case "False", "false", "0":
		return false
	}

	return def
}

func getString(s string, def string) string {
	if s == "" {
		return def
	}

	return s
}

func getFloat64(s string, def float64) float64 {
	if s == "" {
		return def
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}

	return n
}

func getInt(s string, def int) int {
	if s == "" {
		return def
	}
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return def
	}
	return int(n)
}

func string2RGBA(clr string) color.RGBA {
	if c, ok := colors[clr]; ok {
		return c
	}
	c, err := hexToRGBA(clr)
	if err != nil {
		return color.RGBA{0, 0, 0, 255}
	}
	return *c
}

// https://code.google.com/p/sadbox/source/browse/color/hex.go
// hexToRGBA converts an Hex string to a RGB triple.
func hexToRGBA(h string) (*color.RGBA, error) {
	var r, g, b uint8
	if len(h) > 0 && h[0] == '#' {
		h = h[1:]
	}

	if len(h) == 3 {
		h = h[:1] + h[:1] + h[1:2] + h[1:2] + h[2:] + h[2:]
	}

	alpha := byte(255)

	if len(h) == 6 {
		if rgb, err := strconv.ParseUint(string(h), 16, 32); err == nil {
			r = uint8(rgb >> 16)
			g = uint8(rgb >> 8)
			b = uint8(rgb)
		} else {
			return nil, err
		}
	}

	if len(h) == 8 {
		if rgb, err := strconv.ParseUint(string(h), 16, 32); err == nil {
			r = uint8(rgb >> 24)
			g = uint8(rgb >> 16)
			b = uint8(rgb >> 8)
			alpha = uint8(rgb)
		} else {
			return nil, err
		}
	}

	return &color.RGBA{r, g, b, alpha}, nil
}

var colors = map[string]color.RGBA{
	// Graphite default colors
	"black":     {0x00, 0x00, 0x00, 0xff},
	"white":     {0xff, 0xff, 0xff, 0xff},
	"blue":      {0x00, 0x00, 0xff, 0xff},
	"green":     {0x00, 0xff, 0x00, 0xff},
	"red":       {0xff, 0x00, 0x00, 0xff},
	"yellow":    {0xff, 0xff, 0x00, 0xff},
	"orange":    {0xff, 0xa5, 0x00, 0xff},
	"purple":    {0xc8, 0x64, 0xff, 0xff},
	"brown":     {0x96, 0x64, 0x32, 0xff},
	"cyan":      {0x00, 0xff, 0xff, 0xff},
	"aqua":      {0x00, 0x96, 0x96, 0xff},
	"gray":      {0xaf, 0xaf, 0xaf, 0xff},
	"grey":      {0xaf, 0xaf, 0xaf, 0xff},
	"magenta":   {0xff, 0x00, 0xff, 0xff},
	"pink":      {0xff, 0x64, 0x64, 0xff},
	"gold":      {0xc8, 0xc8, 0x00, 0xff},
	"rose":      {0xc8, 0x96, 0xc8, 0xff},
	"darkblue":  {0x00, 0x21, 0x73, 0xff},
	"darkgreen": {0x00, 0xc8, 0x00, 0xff},
	"darkred":   {0xc8, 0x00, 0x32, 0xff},
	"darkgray":  {0x6f, 0x6f, 0x6f, 0xff},
	"darkgrey":  {0x6f, 0x6f, 0x6f, 0xff},

	// Custom colors
	"navy":                 {0x00, 0x00, 0x80, 0xff},
	"mediumblue":           {0x00, 0x00, 0xcd, 0xff},
	"teal":                 {0x00, 0x80, 0x80, 0xff},
	"darkcyan":             {0x00, 0x8b, 0x8b, 0xff},
	"deepskyblue":          {0x00, 0xbf, 0xff, 0xff},
	"darkturquoise":        {0x00, 0xce, 0xd1, 0xff},
	"mediumspringgreen":    {0x00, 0xfa, 0x9a, 0xff},
	"lime":                 {0x00, 0xff, 0x00, 0xff},
	"springgreen":          {0x00, 0xff, 0x7f, 0xff},
	"midnightblue":         {0x19, 0x19, 0x70, 0xff},
	"dodgerblue":           {0x1e, 0x90, 0xff, 0xff},
	"lightseagreen":        {0x20, 0xb2, 0xaa, 0xff},
	"forestgreen":          {0x22, 0x8b, 0x22, 0xff},
	"seagreen":             {0x2e, 0x8b, 0x57, 0xff},
	"darkslategray":        {0x2f, 0x4f, 0x4f, 0xff},
	"limegreen":            {0x32, 0xcd, 0x32, 0xff},
	"mediumseagreen":       {0x3c, 0xb3, 0x71, 0xff},
	"turquoise":            {0x40, 0xe0, 0xd0, 0xff},
	"royalblue":            {0x41, 0x69, 0xe1, 0xff},
	"steelblue":            {0x46, 0x82, 0xb4, 0xff},
	"darkslateblue":        {0x48, 0x3d, 0x8b, 0xff},
	"mediumturquoise":      {0x48, 0xd1, 0xcc, 0xff},
	"indigo":               {0x4b, 0x00, 0x82, 0xff},
	"darkolivegreen":       {0x55, 0x6b, 0x2f, 0xff},
	"cadetblue":            {0x5f, 0x9e, 0xa0, 0xff},
	"cornflowerblue":       {0x64, 0x95, 0xed, 0xff},
	"mediumaquamarine":     {0x66, 0xcd, 0xaa, 0xff},
	"dimgray":              {0x69, 0x69, 0x69, 0xff},
	"slateblue":            {0x6a, 0x5a, 0xcd, 0xff},
	"olivedrab":            {0x6b, 0x8e, 0x23, 0xff},
	"slategray":            {0x70, 0x80, 0x90, 0xff},
	"lightslategray":       {0x77, 0x88, 0x99, 0xff},
	"mediumslateblue":      {0x7b, 0x68, 0xee, 0xff},
	"lawngreen":            {0x7c, 0xfc, 0x00, 0xff},
	"chartreuse":           {0x7f, 0xff, 0x00, 0xff},
	"aquamarine":           {0x7f, 0xff, 0xd4, 0xff},
	"lavender":             {0xe6, 0xe6, 0xfa, 0xff},
	"darksalmon":           {0xe9, 0x96, 0x7a, 0xff},
	"violet":               {0xee, 0x82, 0xee, 0xff},
	"palegoldenrod":        {0xee, 0xe8, 0xaa, 0xff},
	"lightcoral":           {0xf0, 0x80, 0x80, 0xff},
	"khaki":                {0xf0, 0xe6, 0x8c, 0xff},
	"aliceblue":            {0xf0, 0xf8, 0xff, 0xff},
	"honeydew":             {0xf0, 0xff, 0xf0, 0xff},
	"azure":                {0xf0, 0xff, 0xff, 0xff},
	"sandybrown":           {0xf4, 0xa4, 0x60, 0xff},
	"wheat":                {0xf5, 0xde, 0xb3, 0xff},
	"beige":                {0xf5, 0xf5, 0xdc, 0xff},
	"whitesmoke":           {0xf5, 0xf5, 0xf5, 0xff},
	"mintcream":            {0xf5, 0xff, 0xfa, 0xff},
	"ghostwhite":           {0xf8, 0xf8, 0xff, 0xff},
	"salmon":               {0xfa, 0x80, 0x72, 0xff},
	"antiquewhite":         {0xfa, 0xeb, 0xd7, 0xff},
	"linen":                {0xfa, 0xf0, 0xe6, 0xff},
	"lightgoldenrodyellow": {0xfa, 0xfa, 0xd2, 0xff},
	"oldlace":              {0xfd, 0xf5, 0xe6, 0xff},
	"fuchsia":              {0xff, 0x00, 0xff, 0xff},
	"deeppink":             {0xff, 0x14, 0x93, 0xff},
	"orangered":            {0xff, 0x45, 0x00, 0xff},
	"tomato":               {0xff, 0x63, 0x47, 0xff},
	"hotpink":              {0xff, 0x69, 0xb4, 0xff},
	"coral":                {0xff, 0x7f, 0x50, 0xff},
	"darkorange":           {0xff, 0x8c, 0x00, 0xff},
	"lightsalmon":          {0xff, 0xa0, 0x7a, 0xff},
	"lightpink":            {0xff, 0xb6, 0xc1, 0xff},
	"peachpuff":            {0xff, 0xda, 0xb9, 0xff},
	"navajowhite":          {0xff, 0xde, 0xad, 0xff},
	"moccasin":             {0xff, 0xe4, 0xb5, 0xff},
	"bisque":               {0xff, 0xe4, 0xc4, 0xff},
	"mistyrose":            {0xff, 0xe4, 0xe1, 0xff},
	"blanchedalmond":       {0xff, 0xeb, 0xcd, 0xff},
	"papayawhip":           {0xff, 0xef, 0xd5, 0xff},
	"lavenderblush":        {0xff, 0xf0, 0xf5, 0xff},
	"seashell":             {0xff, 0xf5, 0xee, 0xff},
	"cornsilk":             {0xff, 0xf8, 0xdc, 0xff},
	"lemonchiffon":         {0xff, 0xfa, 0xcd, 0xff},
	"floralwhite":          {0xff, 0xfa, 0xf0, 0xff},
	"snow":                 {0xff, 0xfa, 0xfa, 0xff},
	"lightyellow":          {0xff, 0xff, 0xe0, 0xff},
	"ivory":                {0xff, 0xff, 0xf0, 0xff},
}

func SetColor(name, rgba string) error {
	color, err := hexToRGBA(rgba)
	if err != nil {
		return err
	}

	name = strings.ToLower(name)
	colors[name] = *color
	return nil
}
