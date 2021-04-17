// Package api contains the types and request/reply/event payloads associated
// with the JSON protocol used by Music Flow Player apps.
package api

// LG Music Flow protocol messages. Decides what command is sent to the player.
//
// Example:
//
// 	<= {"msg":"EQ_INFO_REQ"}
// 	=> {"data":{"bass":7,"currenteq":12,"lrbal":20,"treble":5},"msg":"EQ_INFO_REQ","result":"OK"}
const (
	MessageAddClient                         = "ADD_CLIENT"
	MessageAddPlaylist                       = "ADD_PLAYLIST"
	MessageAddVMSPlaylist                    = "ADD_VMS_PLAYLIST"
	MessageAlarmBegin                        = "ALARM_BEGIN"
	MessageAlarmListRequest                  = "ALARM_LIST_REQ"
	MessageAlarmSet                          = "ALARM_SET"
	MessageAlarmStateNotification            = "ALARM_STATE_NOTI"
	MessageAlarmStateRequest                 = "ALARM_STATE_REQ"
	MessageAutoDisplaySet                    = "AUTO_DISPLAY_SET"
	MessageAutoPowerSet                      = "AUTO_POWER_SET"
	MessageAutoVolumeSet                     = "AUTO_VOL_SET"
	MessageAVSyncSet                         = "AV_SYNC_SET"
	MessageBluetoothConnection               = "BLUETOOTH_CONNECTION"
	MessageBluetoothDisconnection            = "BLUETOOTH_DISCONNECTION"
	MessageBluetoothInfoRequest              = "BLUETOOTH_INFO_REQ"
	MessageBluetoothPairingResult            = "BLUETOOTH_PAIRING_RESULT"
	MessageBluetoothLimitSet                 = "BT_LIMIT_SET"
	MessageBluetoothLimitSetNotification     = "BT_LIMIT_SET_NOTI"
	MessageBluetoothPartymodeSet             = "BT_PARTYMODE_SET"
	MessageBluetoothStandbySet               = "BT_STANDBY_SET"
	MessageBluetoothStandbyStateNotification = "BT_STANDBY_STATE_NOTI"
	MessageC4AGroupCancelNotification        = "C4A_GROUP_CANCEL_NOTI"
	MessageC4ATOSGet                         = "C4A_TOS_GET"
	MessageChangePlaylistIndex               = "CHANGE_PLAYLIST_IDX"
	MessageChannelInfoRequest                = "CHANNEL_INFO_REQ"
	MessageChannelSet                        = "CHANNEL_SET"
	MessageChannelChangeStatus               = "CH_CHANGE_STATUS"
	MessageCPAddPlaylist                     = "CP_ADD_PLAYLIST"
	MessageCPInfoRequest                     = "CP_INFO_REQ"
	MessageCPPlaylistRequest                 = "CP_PLAYLIST_REQ"
	MessageCPPlayURL                         = "CP_PLAY_URL"
	MessageDeletePlaylist                    = "DELETE_PLAYLIST"
	MessageDRCSet                            = "DRC_SET"
	MessageEqualizerInfoRequest              = "EQ_INFO_REQ"
	MessageEqualizerChangeNotification       = "EQ_NOTI_CHG"
	MessageEqualizerSetting                  = "EQ_SETTING"
	MessageFactorySet                        = "FACTORY_SET"
	MessageFunctionSet                       = "FUNCTION_SET"
	MessageFunctionInfo                      = "FUNC_INFO"
	MessageFunctionInfoRequest               = "FUNC_INFO_REQ"
	MessageGroupCompressSet                  = "GROUP_COMPRESS_SET"
	MessageGroupCompressStateNotification    = "GROUP_COMPRESS_STATE_NOTI"
	MessageGroupDestroy                      = "GROUP_DESTROY"
	MessageGroupSet                          = "GROUP_SET"
	MessageIHRLogon                          = "IHR_LOGON"
	MessageInitializationSet                 = "INITIALIZATION_SET"
	MessageLedSet                            = "LED_SET"
	MessageLocalPlayURL                      = "LOCAL_PLAY_URL"
	MessageLocalTimeSearch                   = "LOCAL_TIME_SEARCH"
	MessageParsingError                      = "MSG_PARSING_ERROR"
	MessageMusicIndexUpdate                  = "MUSIC_INDEX_UPDATE"
	MessageMusicIndexUpdateInfoRequest       = "MUSIC_INDEX_UPDATE_INFO_REQ"
	MessageMuteChange                        = "MUTE_CHANGE"
	MessageMuteSet                           = "MUTE_SET"
	MessageNetworkInfoRequest                = "NETWORK_INFO_REQ"
	MessageNetworkStatusNotification         = "NET_STATUS_NOTI"
	MessageNewVersionSearch                  = "NEW_VER_SEARCH"
	MessageNightModeSet                      = "NIGHT_MODE_SET"
	MessageOnSurroundSet                     = "ON_SURROUND_SET"
	MessagePlaylistChange                    = "PLAYLIST_CHANGE"
	MessagePlaylistTransRequest              = "PLAYLIST_TRANS_REQ"
	MessagePlayCmd                           = "PLAY_CMD"
	MessagePlayInfo                          = "PLAY_INFO"
	MessagePlayInfoRequest                   = "PLAY_INFO_REQ"
	MessagePlayTime                          = "PLAY_TIME"
	MessagePlayTimeSet                       = "PLAY_TIME_SET"
	MessagePowerOff                          = "POWER_OFF"
	MessageProductInfo                       = "PRODUCT_INFO"
	MessageProductInfoUpdate                 = "PRODUCT_INFO_UPDATE"
	MessageRearboxLevelSet                   = "REARBOX_LEVEL_SET"
	MessageReturnLGGroupRequest              = "RETURN_LG_GRP_REQ"
	MessageRhapsodyEvent                     = "RHAPSODY_EVENT"
	MessageRhapsodyLogon                     = "RHAPSODY_LOGON"
	MessageSDPCPListRequest                  = "SDP_CPLIST_REQ"
	MessageSettingInfoRequest                = "SETTING_INFO_REQ"
	MessageSettingInfoNotification           = "SETTING_INFO_NOTI"
	MessageSetAlarmPlaylist                  = "SET_ALARM_PLAYLIST"
	MessageShareHomeInfo                     = "SHARE_HOME_INFO"
	MessageShareHomeSSID                     = "SHARE_HOME_SSID"
	MessageShareNWWired                      = "SHARE_NW_WIRED"
	MessageShareNWWireless                   = "SHARE_NW_WIRELESS"
	MessageSleepInfoRequest                  = "SLEEP_INFO_REQ"
	MessageSleepSet                          = "SLEEP_SET"
	MessageSoundEffectSet                    = "SND_EFFECT_SET"
	MessageSpeakerAddNotification            = "SPK_ADD_NOTI"
	MessageSpeakerAddSet                     = "SPK_ADD_SET"
	MessageSpeakerAlive                      = "SPK_ALIVE"
	MessageSpeakerChannelNotification        = "SPK_CH_NOTI"
	MessageSpeakerChannelSet                 = "SPK_CH_SET"
	MessageSpeakerInfoModify                 = "SPK_INFO_MODIFY"
	MessageSpeakerNameChange                 = "SPK_NAME_CHANGE"
	MessageStartupSoundSet                   = "STARTUP_SOUND_SET"
	MessageSurroundDestroy                   = "SURROUND_DESTROY"
	MessageSurroundSet                       = "SURROUND_SET"
	MessageSystemVersionRequest              = "SYSTEM_VER_REQ"
	MessageTestTone                          = "TEST_TONE"
	MessageTimezoneSet                       = "TIMEZONE_SET"
	MessageTVRemoteSet                       = "TV_REMOTE_SET"
	MessageUpdateComplete                    = "UPDATE_COMPLETE"
	MessageUpdateDownResult                  = "UPDATE_DOWN_RESULT"
	MessageUpdateProgress                    = "UPDATE_PROGRESS"
	MessageUpdateResult                      = "UPDATE_RESULT"
	MessageUpdateStart                       = "UPDATE_START"
	MessageUpdateStartReboot                 = "UPDATE_START_REBOOT"
	MessageUpdateStartWrite                  = "UPDATE_START_WRITE"
	MessageUsageShareGet                     = "USAGE_SHARE_GET"
	MessageUsageShareSet                     = "USAGE_SHARE_SET"
	MessageUsageShareSetNotification         = "USAGE_SHARE_SET_NOTI"
	MessageVMSScanResult                     = "VMS_SCAN_RESULT"
	MessageVolumeChange                      = "VOLUME_CHANGE"
	MessageVolumeDown                        = "VOLUME_DOWN"
	MessageVolumeSetting                     = "VOLUME_SETTING"
	MessageVolumeUp                          = "VOLUME_UP"
	MessageWooferLevelSet                    = "WOOFER_LEVEL_SET"
)

// EqualizerInfo represents the currently saved equalizer.
type EqualizerInfo struct {
	Bass             int       `json:"bass"`
	CurrentEqualizer Equalizer `json:"currenteq"`
	LeftRightBalance int       `json:"lrbal"`
	Treble           int       `json:"treble"`
}

// FunctionInfo represents the active function.
type FunctionInfo struct {
	BluetoothName string   `json:"btname"`
	Type          Function `json:"type"`
	Mute          bool     `json:"mute"`
	Connect       int      `json:"connect"`
}

type Initialization struct {
	Agree    bool   `json:"agree"`
	Sharing  bool   `json:"sharing"`
	Timezone string `json:"timezone"`
}

type NetworkInfo struct {
	Network  Network `json:"network"`
	MeshID   int     `json:"meshid"`
	Password string  `json:"pswd"`
	MeshCh   int     `json:"meshch"`
	SSID     string  `json:"ssid"`
}

type PlayInfo struct {
	AlbumTitle string `json:"albumtitle"`
	Shuffle    bool   `json:"shuffle"`
	C4AAppName string `json:"c4aappname"`
	Repeat     int    `json:"repeat"`
	Source     int    `json:"source"`
	Position   int    `json:"position"`
	Index      int    `json:"idx"`
	AlbumArt   string `json:"albumart"`
	CPType     int    `json:"cptype"`
	Artist     string `json:"artist"`
	URI        string `json:"uri"`
	C4AAppID   string `json:"c4aappid"`
	ObjID      string `json:"objID"`
	Duration   int    `json:"duration"`
	Title      string `json:"title"`
	Playing    int    `json:"playing"`
}

type ProductInfo struct {
	Reg       bool            `json:"reg"`
	ModelType Model           `json:"modeltype"`
	Network   Network         `json:"network"`
	ModelName string          `json:"modelname"`
	ModelNum  int             `json:"modelnum"`
	PetName   string          `json:"petname"`
	ProtoVer  int             `json:"protover"`
	Region    string          `json:"region"`
	Info      ProductInfoInfo `json:"info"`
}

type ProductInfoInfo struct {
	Name         string      `json:"name"`
	Function     Function    `json:"function"`
	GroupID      int         `json:"groupid"`
	Volume       int         `json:"vol"`
	Icon         int         `json:"icon"`
	SpeakerType  Role        `json:"spktype"`
	GroupColor   int         `json:"groupcolor"`
	Channel      int         `json:"channel"`
	Equalizers   []Equalizer `json:"eqlist"`       // Available equalizers.
	Functions    []Function  `json:"functionlist"` // Available functions.
	Playing      bool        `json:"playing"`
	Mute         bool        `json:"mute"`
	LedSet       bool        `json:"ledset"`
	Update       bool        `json:"update"`
	Demo         bool        `json:"demo"`
	BeVer        string      `json:"bever"`
	BluetoothMAC string      `json:"btmac"`
	WirelessMAC  string      `json:"wirelessmac"`
}

type Settings struct {
	AutoDisplay              bool   `json:"autodisplay"` // Not available on SJ 6.
	AutoPower                bool   `json:"autopower"`
	AutoVolume               bool   `json:"autovol"`
	AvSync                   int    `json:"avsync"`
	BatteryUsage             int    `json:"battusage"`
	BluetoothParty           bool   `json:"btparty"`
	BluetoothStandby         bool   `json:"btstandby"`
	DRC                      bool   `json:"drc"`
	EnableBalance            bool   `json:"enable_balance"`
	GroupCompress            int    `json:"groupcompress"`
	IPv4Addr                 string `json:"ipv4addr"`
	IPv6Addr                 string `json:"ipv6addr"`
	LedSet                   bool   `json:"ledset"`
	LimitBluetoothConnection bool   `json:"limit_bt_conn"`
	LimitGroupGcast          bool   `json:"limit_group_gcast"`
	NightMode                bool   `json:"nightmode"`
	RearBoxLevel             int    `json:"rearboxlevel"`  // Not available on SJ 6.
	RearBoxMax               int    `json:"rearboxmax"`    // Not available on SJ 6.
	RearBoxOffset            int    `json:"rearboxoffset"` // Not available on SJ 6.
	RearBoxOn                bool   `json:"rearboxon"`     // Not available on SJ 6.
	SettingInfoVer           int    `json:"settinginfover"`
	SoundEffect              bool   `json:"soundeffect"`  // Not available on SJ 6.
	StartSoundOn             bool   `json:"startsoundon"` // Not available on SJ 6.
	STBTVRemote              bool   `json:"stbtvremote"`  // Not available on SJ 6.
	TVRemote                 bool   `json:"tvremote"`     // Not available on SJ 6.
	VisibleAlarm             bool   `json:"visible_alarm"`
	VisibleInit              bool   `json:"visible_init"`
	VisibleMlibSync          bool   `json:"visible_mlib_sync"`
	VisibleReserveSleep      bool   `json:"visible_reserve_sleep"`
	VisibleToneControl       bool   `json:"visible_tonectrl"`
	VisibleTVConnection      bool   `json:"visible_tvconn"`
	WooferLevel              int    `json:"wooferlevel"`
	WooferMax                int    `json:"woofermax"`
	WooferOffset             int    `json:"wooferoffset"`
}

type SystemVersion struct {
	Be        string `json:"be"`
	Micom     string `json:"micom"`
	Meq       string `json:"meq"`
	HDMI      string `json:"hdmi"`
	C4A       string `json:"c4a"`
	DSP       string `json:"dsp"`
	DemoMusic string `json:"demomusic"`
}
