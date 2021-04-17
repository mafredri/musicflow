package api

type emptyMessage struct{}

func (emptyMessage) IsZero() bool {
	return true
}

type (
	AlarmSetRequest struct {
		Alarm
	}
	AlarmSetReply struct {
		ID int `json:"id"`
	}
)

func (AlarmSetRequest) Message() string       { return MessageAlarmSet }
func (AlarmSetRequest) Reply() *AlarmSetReply { return &AlarmSetReply{ID: -1} } // Starts at zero.

type (
	AlarmListRequest struct {
		emptyMessage
	}
	AlarmListReply struct {
		Info []Alarm `json:"info"`
	}
)

func (AlarmListRequest) Message() string        { return MessageAlarmListRequest }
func (AlarmListRequest) Reply() *AlarmListReply { return &AlarmListReply{} }

type (
	AlarmStateRequest struct {
		emptyMessage
	}
	AlarmStateReply struct {
		On bool `json:"on"`
	}
)

func (AlarmStateRequest) Message() string         { return MessageAlarmStateRequest }
func (AlarmStateRequest) Reply() *AlarmStateReply { return &AlarmStateReply{} }

type SleepSetRequest struct {
	Time int `json:"time"`
}

func (SleepSetRequest) Message() string { return MessageSleepSet }

type SpeakerInfoModifyRequest struct {
	Icon int    `json:"icon"`
	Name string `json:"name"`
}

func (SpeakerInfoModifyRequest) Message() string { return MessageSpeakerInfoModify }

type SpeakerNameChangeEvent struct {
	Icon int    `json:"icon"`
	Name string `json:"name"`
}

func (SpeakerNameChangeEvent) Message() string { return MessageSpeakerNameChange }

type (
	ProductInfoRequest struct {
		ID     string `json:"id"`  // Device ID.
		Day    Day    `json:"day"` // Day of week.
		Hour   int    `json:"hour"`
		Min    int    `json:"min"`
		Option int    `json:"option"` // Either 0 or 1, what does it do? Update time on speaker?
	}
)

func (ProductInfoRequest) Message() string     { return MessageProductInfo }
func (ProductInfoRequest) Reply() *ProductInfo { return &ProductInfo{} }

type (
	NightModeSetRequest struct {
		NightMode bool `json:"nightmode"`
	}
	NightModeSetReply struct {
		NightMode bool `json:"nightmode"`
	}
)

func (NightModeSetRequest) Message() string           { return MessageNightModeSet }
func (NightModeSetRequest) Reply() *NightModeSetReply { return &NightModeSetReply{} }

type (
	VolumeSettingRequest struct {
		FadeTime int `json:"fadetime"`
		Volume   int `json:"vol"`
	}
	VolumeSettingReply struct {
		Volume int `json:"vol"`
	}
)

func (VolumeSettingRequest) Message() string { return MessageVolumeSetting }

type MuteSetRequest struct {
	Mute bool `json:"mute"`
}

func (MuteSetRequest) Message() string { return MessageMuteSet }

type MuteChangeEvent struct {
	Mute bool `json:"mute"`
}

type SystemVersionRequest struct {
	emptyMessage
}

func (SystemVersionRequest) Message() string       { return MessageSystemVersionRequest }
func (SystemVersionRequest) Reply() *SystemVersion { return &SystemVersion{} }

type SettingInfoRequest struct {
	emptyMessage
}

func (SettingInfoRequest) Message() string  { return MessageSettingInfoRequest }
func (SettingInfoRequest) Reply() *Settings { return &Settings{} }

type TestToneRequest struct {
	Stat bool `json:"stat"`
}

func (TestToneRequest) Message() string { return MessageTestTone }

type InitializationSetRequest struct {
	Initialization
}

func (InitializationSetRequest) Message() string { return MessageInitializationSet }

type NetworkInfoRequest struct {
	emptyMessage
}

func (NetworkInfoRequest) Message() string     { return MessageNetworkInfoRequest }
func (NetworkInfoRequest) Reply() *NetworkInfo { return &NetworkInfo{} }

type PlayInfoRequest struct {
	emptyMessage
}

func (PlayInfoRequest) Message() string  { return MessagePlayInfoRequest }
func (PlayInfoRequest) Reply() *PlayInfo { return &PlayInfo{} }

type (
	WooferLevelSetRequest struct {
		Level int `json:"wooferlevel"`
	}
	WooferLevelSetReply struct {
		Level int `json:"wooferlevel"`
	}
)

func (WooferLevelSetRequest) Message() string             { return MessageWooferLevelSet }
func (WooferLevelSetRequest) Reply() *WooferLevelSetReply { return &WooferLevelSetReply{} }

type EqualizerInfoRequest struct {
	emptyMessage
}

func (EqualizerInfoRequest) Message() string       { return MessageEqualizerInfoRequest }
func (EqualizerInfoRequest) Reply() *EqualizerInfo { return &EqualizerInfo{} }

type EqualizerSetRequest struct {
	Type  EqualizerType `json:"type"`
	Value int           `json:"value"` // E.g. for SetEqualizer, use int(Equalizer).
}

func (EqualizerSetRequest) Message() string { return MessageEqualizerSetting }

type FunctionSetRequest struct {
	Type Function `json:"type"`
}

func (FunctionSetRequest) Message() string { return MessageFunctionSet }

type FunctionInfoRequest struct {
	emptyMessage
}

func (FunctionInfoRequest) Message() string      { return MessageFunctionInfoRequest }
func (FunctionInfoRequest) Reply() *FunctionInfo { return &FunctionInfo{} }
