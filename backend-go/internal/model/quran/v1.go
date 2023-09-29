package quran

import "github.com/siroj100/hikarie-islamy/internal/model/db"

type (
	V1ListSuratItem struct {
		db.QuranSurat `yaml:",inline"`

		L10N map[string]db.QuranSuratL10N `yaml:"l10n"`
	}
	V1ListSuratResp struct {
		LayoutID   int `yaml:"layoutId"`
		NumOfLines int `yaml:"numOfLines"`

		ListSurat []V1ListSuratItem `yaml:"listSurat"`
	}

	V1PageReq struct {
		LayoutID   int `param:"layoutID"`
		PageNumber int `param:"pageNumber"`
	}

	V1PageLineAyatResp struct {
		SuratID int    `json:"suratId" yaml:"suratId"`
		AyatID  int    `json:"ayatId" yaml:"ayatId"`
		Text    string `json:"text"`
		AyatEnd bool   `json:"ayatEnd"`

		CharStart int `json:"charStart" yaml:"charStart"`
		TotalChar int `json:"totalChar" yaml:"totalChar"`
	}

	V1PageLineResp struct {
		StartEndOfSurat bool `json:"startEndOfSurat" yaml:"startEndOfSurat"`
		EndEndOfSurat   bool `json:"endEndOfSurat" yaml:"endEndOfSurat"`
		Basmalah        bool `json:"basmalah" yaml:"basmalah"`

		ListAyat []V1PageLineAyatResp `json:"listAyat" yaml:"listAyat"`
	}

	V1PageResp struct {
		Surat      db.QuranSurat `json:"surat" yaml:"surat"`
		LayoutID   int           `json:"layoutId" yaml:"layoutId"`
		PageNumber int           `json:"pageNumber" yaml:"pageNumber"`
		NumOfLines int           `json:"numOfLines" yaml:"numOfLines"`

		ListLine []V1PageLineResp `json:"listLine" yaml:"listLine"`
	}
)
