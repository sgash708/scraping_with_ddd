package configutil_test

import (
	"testing"

	"github.com/sgash708/scraping_with_ddd/utils/configutil"
)

// TestSetup_Success セットアップ成功
func TestSetup_Success(t *testing.T) {
	cases := map[string]struct {
		LAWURL  string
		OFFURL  string
		CWToken string
		CWURL   string
		CWMes   string
		DBUser  string
		DBPass  string
		DBName  string
		DBHost  string
		DBPort  int
		SSHHost string
		SSHPort int
		SSHUser string
		SSHPass string
	}{
		"config.ini.sample": {
			LAWURL:  "https://example.com/lawyer_url",
			OFFURL:  "https://example.com/office_url",
			CWToken: "xxxxxxxxxxxxx",
			CWURL:   "https://api.chatwork.com/v2/rooms/1111111111/messages",
			CWMes: `[toall]
実行日時: `,
			DBUser:  "xxxxxxx",
			DBPass:  "xxxxxxx",
			DBName:  "xxxxxxxxx",
			DBHost:  "xxxxxxxxx",
			DBPort:  3306,
			SSHHost: "xxxxxx",
			SSHPort: 22,
			SSHUser: "xxxxx",
			SSHPass: `-----BEGIN RSA PRIVATE KEY-----
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
-----END RSA PRIVATE KEY-----`,
		},
	}

	for key, c := range cases {
		t.Run("読み込み成功", func(t *testing.T) {
			config := configutil.NewConfig(key)
			Scraping, err := config.Setup()
			if err != nil {
				t.Fatalf("ファイルの読み込みに失敗しました。\n詳細:%v", err)
			}

			if Scraping.LAWURL != c.LAWURL {
				t.Fatal("URLが間違っています。")
			}
			if Scraping.OFFURL != c.OFFURL {
				t.Fatal("URLが間違っています。")
			}
			if Scraping.CWToken != c.CWToken {
				t.Fatal("chatworkのtokenが間違っています。")
			}
			if Scraping.CWURL != c.CWURL {
				t.Fatal("chatworkのendpointが間違っています。")
			}
			if Scraping.CWMes != c.CWMes {
				t.Fatal("chatworkのmessageが間違っています。")
			}
			if Scraping.DBUser != c.DBUser {
				t.Fatal("dbのuserが間違っています。")
			}
			if Scraping.DBPass != c.DBPass {
				t.Fatal("dbのpasswordが間違っています。")
			}
			if Scraping.DBName != c.DBName {
				t.Fatal("dbのnameが間違っています。")
			}
			if Scraping.DBHost != c.DBHost {
				t.Fatal("dbのhostが間違っています。")
			}
			if Scraping.DBPort != c.DBPort {
				t.Fatal("dbのportが間違っています。")
			}
			if Scraping.SSHHost != c.SSHHost {
				t.Fatal("sshのhostが間違っています。")
			}
			if Scraping.SSHPort != c.SSHPort {
				t.Fatal("sshのportが間違っています。")
			}
			if Scraping.SSHUser != c.SSHUser {
				t.Fatal("sshのuserが間違っています。")
			}
			if Scraping.SSHPass != c.SSHPass {
				t.Fatal("sshのpassが間違っています。")
			}
		})
	}
}

// TestSetup_FailReadingFile セットアップ失敗
func TestSetup_FailReadingFile(t *testing.T) {
	cases := map[string]struct {
		URL       string
		UserName  string
		UserPass  string
		CWURL     string
		CWToken   string
		CWMessage string
		Profile   string
		S3Bucket  string
		S3URL     string
		ALBww9    string
		ALBwww    string
	}{
		"config.example": {
			URL: "https://test.example.com/",
		},
	}

	for key := range cases {
		t.Run("読み込み失敗_存在しないファイル", func(t *testing.T) {
			config := configutil.NewConfig(key)
			_, err := config.Setup()
			if err == nil {
				t.Fatalf("ファイルを読み込めています。確認してください。\n詳細:%v", err)
			}
		})
	}
}
