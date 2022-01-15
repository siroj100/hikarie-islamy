package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

type (
	KemenagRepo struct {
		client http.Client
	}
)

func NewKemenag(client http.Client) KemenagRepo {
	return KemenagRepo{client: client}
}

func (r KemenagRepo) GetListKemenagSurat(ctx context.Context) ([]model.QuranKemenagSurat, error) {
	kemenagResp := struct {
		Data []model.QuranKemenagSurat `json:"data"`
	}{}
	resp, err := r.client.Get("https://quran.kemenag.go.id/index.php/api/v1/surat/0/114")
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return kemenagResp.Data, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(errorx.PrintTrace(errorx.ErrInvalidResponse), resp.StatusCode, resp.Status)
		return kemenagResp.Data, errorx.ErrInvalidResponse
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return kemenagResp.Data, err
	}

	err = json.Unmarshal(body, &kemenagResp)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return kemenagResp.Data, err
	}

	return kemenagResp.Data, nil
}

func (r KemenagRepo) GetListKemenagAyat(ctx context.Context, suratID, ayatStart, ayatCount int) ([]model.QuranKemenagAyat, error) {
	kemenagResp := struct {
		Data []model.QuranKemenagAyat `json:"data"`
	}{}

	resp, err := r.client.Get(fmt.Sprintf("https://quran.kemenag.go.id/index.php/api/v1/ayatweb/%d/0/%d/%d", suratID, ayatStart, ayatCount))
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return kemenagResp.Data, err
	}
	if resp.StatusCode != 200 {
		log.Println(errorx.PrintTrace(errorx.ErrInvalidResponse), resp.StatusCode, resp.Status)
		return kemenagResp.Data, errorx.ErrInvalidResponse
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return kemenagResp.Data, err
	}

	err = json.Unmarshal(body, &kemenagResp)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return kemenagResp.Data, err
	}

	return kemenagResp.Data, nil

}
