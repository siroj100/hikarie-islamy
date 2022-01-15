package model

import "github.com/siroj100/hikarie-islamy/internal/model/db"

type (
	QuranData struct {
		LangData      db.QuranL10N
		ListSurat     []db.QuranSurat
		ListSuratL10N []db.QuranSuratL10N
		ListAyat      map[int][]db.QuranAyat
		ListAyatL10N  map[int][]db.QuranAyatL10N
	}
)
