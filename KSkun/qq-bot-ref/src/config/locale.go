package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TypeLocale struct {
	Success          string `json:"success"`
	Failed           string `json:"failed"`
	FailedDebug      string `json:"failed_debug"`
	UnknownOperation string `json:"unknown_operation"`

	FuncHelp   string `json:"func_help"`
	FuncAdd    string `json:"func_add"`
	FuncDelete string `json:"func_delete"`
	FuncList   string `json:"func_list"`
	FuncRemind string `json:"func_remind"`

	HelpPrefix string `json:"help_prefix"`
	Help       string `json:"help"`

	AddPrefix                string `json:"add_prefix"`
	AddRequireTime           string `json:"add_require_time"`
	AddRequireRemind         string `json:"add_require_remind"`
	AddRequireRemindBefore   string `json:"add_require_remind_before"`
	AddRequireRemindInterval string `json:"add_require_remind_interval"`

	DeletePrefix  string `json:"delete_prefix"`
	DeleteConfirm string `json:"delete_confirm"`
	DeleteAbort   string `json:"delete_abort"`

	ListPrefix      string `json:"list_prefix"`
	ListEntry       string `json:"list_entry"`
	ListEntryRemind string `json:"list_entry_remind"`
	ListEmpty       string `json:"list_empty"`

	Remind       string `json:"remind"`
	RemindExpire string `json:"remind_expire"`
}

var Locale TypeLocale

func InitLocale() {
	localeFilename := "default.json" // use `default.json` as default filename
	// set env variable CONFIG_FILE to use other config file
	if filename, ok := os.LookupEnv("LOCALE_FILE"); ok {
		localeFilename = filename
	}

	localeFile, err := ioutil.ReadFile("./locale/" + localeFilename)
	if err != nil {
		log.Println("config: error when read locale file " + localeFilename)
		log.Panic(err)
	}

	err = json.Unmarshal(localeFile, &Locale)
	if err != nil {
		log.Println("config: error when unmarshal locale")
		log.Panic(err)
	}

	log.Println("config: locale " + localeFilename + " loaded")
}
