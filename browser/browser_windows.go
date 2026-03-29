//go:build windows

package browser

import (
	"github.com/moond4rk/hackbrowserdata/types"
)

var (
	chromiumList = map[string]struct {
		name        string
		profilePath string
		storage     string
		dataTypes   []types.DataType
	}{
		"chrome": {
			name:        chromeName,
			profilePath: chromeUserDataPath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"edge": {
			name:        edgeName,
			profilePath: edgeProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"chromium": {
			name:        chromiumName,
			profilePath: chromiumUserDataPath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"chrome-beta": {
			name:        chromeBetaName,
			profilePath: chromeBetaUserDataPath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"opera": {
			name:        operaName,
			profilePath: operaProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"opera-gx": {
			name:        operaGXName,
			profilePath: operaGXProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"vivaldi": {
			name:        vivaldiName,
			profilePath: vivaldiProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"coccoc": {
			name:        coccocName,
			profilePath: coccocProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"brave": {
			name:        braveName,
			profilePath: braveProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"yandex": {
			name:        yandexName,
			profilePath: yandexProfilePath,
			dataTypes:   types.DefaultYandexTypes,
		},
		"360": {
			name:        speed360Name,
			profilePath: speed360ProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"qq": {
			name:        qqBrowserName,
			profilePath: qqBrowserProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"dc": {
			name:        dcBrowserName,
			profilePath: dcBrowserProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"sogou": {
			name:        sogouName,
			profilePath: sogouProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
	}
	firefoxList = map[string]struct {
		name        string
		storage     string
		profilePath string
		dataTypes   []types.DataType
	}{
		"firefox": {
			name:        firefoxName,
			profilePath: firefoxProfilePath,
			dataTypes:   types.DefaultFirefoxTypes,
		},
	}
)

var (
	chromeUserDataPath     = GetBrowserProfilePath("AppData/Local/Google/Chrome/User Data/Default/")
	chromeBetaUserDataPath = GetBrowserProfilePath("AppData/Local/Google/Chrome Beta/User Data/Default/")
	chromiumUserDataPath   = GetBrowserProfilePath("AppData/Local/Chromium/User Data/Default/")
	edgeProfilePath        = GetBrowserProfilePath("AppData/Local/Microsoft/Edge/User Data/Default/")
	braveProfilePath       = GetBrowserProfilePath("AppData/Local/BraveSoftware/Brave-Browser/User Data/Default/")
	speed360ProfilePath    = GetBrowserProfilePath("AppData/Local/360chrome/Chrome/User Data/Default/")
	qqBrowserProfilePath   = GetBrowserProfilePath("AppData/Local/Tencent/QQBrowser/User Data/Default/")
	operaProfilePath       = GetBrowserProfilePath("AppData/Roaming/Opera Software/Opera Stable/")
	operaGXProfilePath     = GetBrowserProfilePath("AppData/Roaming/Opera Software/Opera GX Stable/")
	vivaldiProfilePath     = GetBrowserProfilePath("AppData/Local/Vivaldi/User Data/Default/")
	coccocProfilePath      = GetBrowserProfilePath("AppData/Local/CocCoc/Browser/User Data/Default/")
	yandexProfilePath      = GetBrowserProfilePath("AppData/Local/Yandex/YandexBrowser/User Data/Default/")
	dcBrowserProfilePath   = GetBrowserProfilePath("AppData/Local/DCBrowser/User Data/Default/")
	sogouProfilePath       = GetBrowserProfilePath("AppData/Roaming/SogouExplorer/Webkit/Default/")

	firefoxProfilePath = GetBrowserProfilePath("AppData/Roaming/Mozilla/Firefox/Profiles/")
)
