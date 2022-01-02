package utils_test

import (
	"log"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/sgash708/scraping_with_ddd/utils"
)

// TestCalPages ページ数計算
func TestCalPages(t *testing.T) {
	cases := []struct {
		count string
		pages int
		err   bool
	}{
		{count: "1000", pages: 50, err: false},
		{count: "1001", pages: 51, err: false},
		{count: "0", pages: 0, err: true},
		{count: "test", pages: 0, err: true},
		{count: "1000000000000000000000000000000000000000000000000000000000000000000000000000000000", pages: 0, err: true},
		{count: "+1000", pages: 50, err: false},
		{count: "+1001", pages: 51, err: false},
		{count: "-1001", pages: 0, err: false},
	}

	for _, c := range cases {
		pageNum, err := common.CalPages(c.count)
		if err != nil && !c.err {
			t.Errorf("数値変換ができていないですよ。\n詳細:%v", err)
		}
		if pageNum != c.pages {
			t.Errorf("ページの計算ができていません。\nreal:%v\nexpected:%v", pageNum, c.pages)
		}
	}
}

// TestSetLog ログ出力テスト
func TestSetLog(t *testing.T) {
	cases := []struct {
		file    string
		comment string
	}{
		{file: "test.log", comment: "こんにちは"},
		{file: "test.txt", comment: "こんばんは"},
	}

	for _, c := range cases {
		common.SetLog(c.file)
		// Log書き込み
		log.Println(c.comment)

		out, err := os.Open(c.file)
		defer func() {
			if err := out.Close(); err != nil {
				log.Fatalln(err)
			}
		}()

		if err != nil {
			t.Errorf("ファイルを開けませんでした。\n%s", err)
		}

		cont, err := os.ReadFile(c.file)
		if err != nil {
			t.Errorf("ファイルを読み込めませんでした。\n%v", err)
		}

		ret := string(cont)
		if !strings.Contains(ret, c.comment) {
			t.Errorf("ログを書き込めていません。\nexpected:%v\nreal:%v", c.comment, ret)
		}

		if err := os.Remove(c.file); err != nil {
			log.Fatalln(err)
		}
	}
}

// TestConvertImageToStr_Success 画像変換テスト成功
func TestConvertImageToStr_Success(t *testing.T) {
	cases := []struct {
		file  string
		kanji string
	}{
		{file: "o.png", kanji: "緒"},
		{file: "saki.png", kanji: "崎"},
	}

	for _, c := range cases {
		kanji, err := common.ConvertImageToStr("../../test/image/" + c.file)
		if err != nil {
			t.Fatalf("漢字の変換に失敗しました。\n詳細:%v", err)
		}

		rs := []rune(kanji)
		if !unicode.In(rs[0], unicode.Han) {
			t.Fatal("漢字ではありません。")
		}

		if kanji != c.kanji {
			t.Errorf("期待する漢字と一致していません。\nreal:%v\nexpected:%v", kanji, c.kanji)
		}
	}
}

// TestConvertImageToStr_Fail 画像変換テスト失敗
func TestConvertImageToStr_Fail(t *testing.T) {
	cases := []struct {
		file  string
		kanji string
	}{
		{file: "vue.png", kanji: "Y"},
	}

	for _, c := range cases {
		kanji, err := common.ConvertImageToStr("../../test/image/" + c.file)
		if err != nil {
			t.Fatalf("漢字の変換に失敗しました。\n詳細:%v", err)
		}

		// REF: https: //qiita.com/tomtwinkle/items/d52a01d5a020b00b4b8e
		rs := []rune(kanji)
		if unicode.In(rs[0], unicode.Han) {
			t.Fatal("漢字にならない画像です。")
		}
	}
}

// TODO
// TestDownloadFile ファイルダウンロード記述
// TestDownloadFile(t *testing.T) {}

// TestCreateDirectory ディレクトリ作成テスト
func TestCreateDirectory(t *testing.T) {
	cases := []struct {
		dirname string
	}{
		{dirname: "./test"},
		{dirname: "./test.png\\/"},
		{dirname: "../test.png\\/"},
	}

	for _, c := range cases {
		if err := common.CreateDirectory(c.dirname); err != nil {
			t.Errorf("ディレクトリの作成に失敗しています。\n詳細:%v", err)
		}

		if _, err := os.Stat(c.dirname); os.IsNotExist(err) {
			t.Errorf("ディレクトリが存在していませんでした。\n詳細:%v", err)
		}

		if err := os.Remove(c.dirname); err != nil {
			t.Fatal(err)
		}
	}
}

// TestB2i booleanのint変換テスト
func TestB2i(t *testing.T) {
	cases := []struct {
		flg      bool
		expected int8
	}{
		{flg: true, expected: 1},
		{flg: false, expected: 0},
	}

	for _, c := range cases {
		resInt := common.B2i(c.flg)
		if resInt != c.expected {
			t.Errorf("要素が一致しません。\nreal:%v\nexpected:%v", resInt, c.expected)
		}
	}
}
