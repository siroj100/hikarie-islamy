package model

type (
	QuranKemenagSurat struct {
		ID              int    `json:"id"`
		SuratName       string `json:"surat_name"`
		SuratText       string `json:"surat_text"`
		SuratTerjemahan string `json:"surat_terjemahan"`
		CountAyat       int    `json:"count_ayat"`
	}

	QuranKemenagAyat struct {
		AyaID              int    `json:"aya_id"`
		AyaNumber          int    `json:"aya_number"`
		AyaText            string `json:"aya_text"`
		SuraID             int    `json:"sura_id"`
		JuzID              int    `json:"juz_id"`
		PageNumber         int    `json:"page_number"`
		TranslationAyaText string `json:"translation_aya_text"`
	}

	QuranKemenagV1ListSuratReq struct {
		StartSurat int `param:"startSurat"`
		Count      int `param:"count"`
	}

	QuranKemenagV1ListAyatReq struct {
		SuratID   int `param:"suratID"`
		StartAyat int `param:"startAyat"`
		Count     int `param:"count"`
	}
)
