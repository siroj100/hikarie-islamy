package v1

type (
	QuranSura struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		AyaCount    int    `json:"aya_count"`
		PageStart   int    `json:"page_start"`
		Translation string `json:"translation"`
	}

	ListSuraResp struct {
		ListSura []QuranSura `json:"list_sura"`
	}

	QuranPageAya struct {
		ID          int    `json:"id"`
		Number      int    `json:"number"`
		Text        string `json:"text"`
		Translation string `json:"translation"`
	}

	QuranPageSura struct {
		ID          int            `json:"id"`
		Name        string         `json:"name"`
		AyaCount    int            `json:"aya_count"`
		Translation string         `json:"translation"`
		Ayas        []QuranPageAya `json:"ayas"`
	}

	QuranPage struct {
		JuzID  int             `json:"juz_id"`
		Number int             `json:"number"`
		Suras  []QuranPageSura `json:"suras"`
	}
)
