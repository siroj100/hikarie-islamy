package db

type (
	QuranL10N struct {
		LangID     int `gorm:"primaryKey"`
		LangCode   string
		SourceID   string `gorm:"index:idx_quran_l10n_src_id,unique"`
		SourceName string
		SourceDesc string
	}

	QuranSurat struct {
		SuratID   int `gorm:"primaryKey"`
		Name      string
		AyatCount int
	}

	QuranSuratL10N struct {
		SuratID   int `gorm:"index:idx_quran_surat_l10n,unique"`
		LangID    int `gorm:"index:idx_quran_surat_l10n,unique"`
		Translit  string
		Translate string
	}

	QuranAyat struct {
		AyatID     int `gorm:"primaryKey"`
		SuratID    int `gorm:"index:idx_quran_ayat,unique"`
		AyatNumber int `gorm:"index:idx_quran_ayat,unique"`
		JuzID      int
		AyatText   string
	}

	QuranAyatL10N struct {
		AyatID     int `gorm:"index:idx_quran_ayat_l10n,unique"`
		LangID     int `gorm:"index:idx_quran_ayat_l10n,unique"`
		PageNumber int
		Translit   string
		Translate  string
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
