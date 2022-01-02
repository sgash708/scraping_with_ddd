package utils

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/otiai10/gosseract/v2"
)

// CalPages ページ計算
func CalPages(cnt string) (int, error) {
	tCount, err := strconv.ParseFloat(cnt, 64)
	if err != nil {
		return 0, err
	}

	num := int(math.Ceil(float64(tCount / 20)))
	if num < 0 {
		return 0, nil
	}

	return num, nil
}

// SetLog ログ出力先を決定
func SetLog(file string) {
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stderr, logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

// REF: https://github.com/tesseract-ocr/tesseract/wiki/Compiling#windows
// ConvertImageToStr 画像を漢字に変換する
func ConvertImageToStr(imageFileName string) (string, error) {
	client := gosseract.NewClient()
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// REF: https://developers.freee.co.jp/entry/2019/12/11/130506
	if err := client.SetLanguage("jpn"); err != nil {
		return "", err
	}
	if err := client.SetImage(imageFileName); err != nil {
		return "", err
	}

	text, err := client.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}

// DownloadFile ファイルダウンロード
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	_, err = io.Copy(out, resp.Body)

	return err
}

// CreateDirectory ディレクトリ作成
func CreateDirectory(baseDirName string) error {
	if _, err := os.Stat(baseDirName); os.IsNotExist(err) {
		if err = os.Mkdir(baseDirName, 0777); err != nil {
			return fmt.Errorf("ディレクトリの作成ができませんでした。：%v", err)
		}
	}

	return nil
}

// B2i boolを数値変換
func B2i(b bool) int8 {
	if b {
		return 1
	}

	return 0
}
