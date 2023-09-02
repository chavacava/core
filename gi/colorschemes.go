// Code generated by "goki colorgen"; DO NOT EDIT

package gi

import (
	"goki.dev/matcolor"
	"goki.dev/colors"
)

// ColorSchemes contains the color schemes used to style
// the app. You should set them in your mainrun funtion
// if you want to change the color schemes.
// [ColorScheme] is set based on ColorSchemes and
// the user preferences. The default color schemes
// are generated through the "goki colorgen" command,
// and you should use the same command to generate your
// custom color scheme. You should not directly access
// ColorSchemes when styling things; instead, you should
// access [ColorScheme].
var ColorSchemes = matcolor.Schemes{
	Light: matcolor.Scheme{
		Primary:                 colors.MustFromHex("#005AC1"),
		OnPrimary:               colors.MustFromHex("#FFFFFF"),
		PrimaryContainer:        colors.MustFromHex("#D8E2FF"),
		OnPrimaryContainer:      colors.MustFromHex("#001A41"),
		Secondary:               colors.MustFromHex("#575E71"),
		OnSecondary:             colors.MustFromHex("#FFFFFF"),
		SecondaryContainer:      colors.MustFromHex("#DBE2F9"),
		OnSecondaryContainer:    colors.MustFromHex("#141B2C"),
		Tertiary:                colors.MustFromHex("#00658B"),
		OnTertiary:              colors.MustFromHex("#FFFFFF"),
		TertiaryContainer:       colors.MustFromHex("#C4E7FF"),
		OnTertiaryContainer:     colors.MustFromHex("#001E2C"),
		Error:                   colors.MustFromHex("#BA1A1A"),
		ErrorContainer:          colors.MustFromHex("#FFDAD6"),
		OnError:                 colors.MustFromHex("#FFFFFF"),
		OnErrorContainer:        colors.MustFromHex("#410002"),
		Background:              colors.MustFromHex("#FEFBFF"),
		OnBackground:            colors.MustFromHex("#1B1B1F"),
		Surface:                 colors.MustFromHex("#FEFBFF"),
		OnSurface:               colors.MustFromHex("#1B1B1F"),
		SurfaceVariant:          colors.MustFromHex("#E1E2EC"),
		OnSurfaceVariant:        colors.MustFromHex("#44474F"),
		Outline:                 colors.MustFromHex("#74777F"),
		InverseOnSurface:        colors.MustFromHex("#F2F0F4"),
		InverseSurface:          colors.MustFromHex("#303033"),
		InversePrimary:          colors.MustFromHex("#ADC6FF"),
		Shadow:                  colors.MustFromHex("#000000"),
		SurfaceTint:             colors.MustFromHex("#005AC1"),
		OutlineVariant:          colors.MustFromHex("#C4C6D0"),
		Scrim:                   colors.MustFromHex("#000000"),
		SurfaceDim:              colors.MustFromHex("#DED8E1"),
		SurfaceBright:           colors.MustFromHex("#FEF7FF"),
		SurfaceContainerLowest:  colors.MustFromHex("#FFFFFF"),
		SurfaceContainerLow:     colors.MustFromHex("#F7F2FA"),
		SurfaceContainer:        colors.MustFromHex("#F3EDF7"),
		SurfaceContainerHigh:    colors.MustFromHex("#ECE6F0"),
		SurfaceContainerHighest: colors.MustFromHex("#E6E0E9"),
	},
	Dark: matcolor.Scheme{
		Primary:                 colors.MustFromHex("#ADC6FF"),
		OnPrimary:               colors.MustFromHex("#002E69"),
		PrimaryContainer:        colors.MustFromHex("#004494"),
		OnPrimaryContainer:      colors.MustFromHex("#D8E2FF"),
		Secondary:               colors.MustFromHex("#BFC6DC"),
		OnSecondary:             colors.MustFromHex("#293041"),
		SecondaryContainer:      colors.MustFromHex("#3F4759"),
		OnSecondaryContainer:    colors.MustFromHex("#DBE2F9"),
		Tertiary:                colors.MustFromHex("#7CD0FF"),
		OnTertiary:              colors.MustFromHex("#00344A"),
		TertiaryContainer:       colors.MustFromHex("#004C69"),
		OnTertiaryContainer:     colors.MustFromHex("#C4E7FF"),
		Error:                   colors.MustFromHex("#FFB4AB"),
		ErrorContainer:          colors.MustFromHex("#93000A"),
		OnError:                 colors.MustFromHex("#690005"),
		OnErrorContainer:        colors.MustFromHex("#FFDAD6"),
		Background:              colors.MustFromHex("#1B1B1F"),
		OnBackground:            colors.MustFromHex("#E3E2E6"),
		Surface:                 colors.MustFromHex("#1B1B1F"),
		OnSurface:               colors.MustFromHex("#E3E2E6"),
		SurfaceVariant:          colors.MustFromHex("#44474F"),
		OnSurfaceVariant:        colors.MustFromHex("#C4C6D0"),
		Outline:                 colors.MustFromHex("#8E9099"),
		InverseOnSurface:        colors.MustFromHex("#1B1B1F"),
		InverseSurface:          colors.MustFromHex("#E3E2E6"),
		InversePrimary:          colors.MustFromHex("#005AC1"),
		Shadow:                  colors.MustFromHex("#000000"),
		SurfaceTint:             colors.MustFromHex("#ADC6FF"),
		OutlineVariant:          colors.MustFromHex("#44474F"),
		Scrim:                   colors.MustFromHex("#000000"),
		SurfaceDim:              colors.MustFromHex("#141218"),
		SurfaceBright:           colors.MustFromHex("#3B383E"),
		SurfaceContainerLowest:  colors.MustFromHex("#0F0D13"),
		SurfaceContainerLow:     colors.MustFromHex("#1D1B20"),
		SurfaceContainer:        colors.MustFromHex("#211F26"),
		SurfaceContainerHigh:    colors.MustFromHex("#2B2930"),
		SurfaceContainerHighest: colors.MustFromHex("#36343B"),
	},
}
