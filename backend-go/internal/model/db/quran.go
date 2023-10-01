package db

import "database/sql"

type (
	QuranL10N struct {
		LangID     int `gorm:"primaryKey"`
		LangCode   string
		SourceID   string `gorm:"index:idx_quran_l10n_src_id,unique"`
		SourceName string
		SourceDesc string
	}

	QuranSurat struct {
		SuratID   int    `gorm:"primaryKey" json:"suratId" yaml:"suratId"`
		Name      string `json:"name" yaml:"name"`
		AyatCount int    `json:"ayatCount" yaml:"ayatCount"`
	}

	QuranSuratL10N struct {
		SuratID   int    `gorm:"index:idx_quran_surat_l10n,unique" yaml:"suratId"`
		LangID    int    `gorm:"index:idx_quran_surat_l10n,unique" yaml:"langId"`
		Translit  string `yaml:"translit"`
		Translate string `yaml:"translate"`

		Surat QuranSurat `gorm:"foreignKey:SuratID"`
	}

	QuranAyat struct {
		AyatID     int `gorm:"primaryKey"`
		SuratID    int `gorm:"index:idx_quran_ayat,unique"`
		AyatNumber int `gorm:"index:idx_quran_ayat,unique"`
		JuzID      int
		AyatText   string
		PageNumber int
	}

	QuranAyatL10N struct {
		AyatID int `gorm:"index:idx_quran_ayat_l10n,unique"`
		LangID int `gorm:"index:idx_quran_ayat_l10n,unique"`
		//PageNumber int
		Translit  string
		Translate string

		Ayat QuranAyat `gorm:"foreignKey:AyatID"`
	}

	QuranLayout struct {
		LayoutID    int `gorm:"primaryKey"`
		Name        string
		Description string
		NumOfLines  int
	}

	QuranPageDetail struct {
		LayoutID     int `gorm:"index:idx_quran_page_detail,unique"`
		PageNumber   int `gorm:"index:idx_quran_page_detail,unique"`
		LineNumber   int `gorm:"index:idx_quran_page_detail,unique"`
		SuratID      int `gorm:"index:idx_quran_page_detail,unique"`
		AyatNumber   int `gorm:"index:idx_quran_page_detail,unique"`
		IsBasmalah   sql.NullBool
		IsSuratTitle sql.NullBool
		StartIndex   int
		EndIndex     int
	}
)

func (m QuranL10N) TableName() string {
	return "quran_l10n"
}

func (m QuranSurat) TableName() string {
	return "quran_surat"
}

func (m QuranSuratL10N) TableName() string {
	return "quran_surat_l10n"
}

func (m QuranAyat) TableName() string {
	return "quran_ayat"
}

func (m QuranAyatL10N) TableName() string {
	return "quran_ayat_l10n"
}

func (m QuranLayout) TableName() string {
	return "quran_layout"
}

func (m QuranPageDetail) TableName() string {
	return "quran_page_detail"
}
