package api

import "fmt"

// Equalizer represents an equalizer. A list of possible equalizers are
// available when requesting product info.
type Equalizer int

// Equalizer enums.
// {"data": {"type": 0, "value": 12}, "msg": "EQ_SETTING"}
const (
	EqualizerStandard      Equalizer = 0
	EqualizerBass          Equalizer = 1
	EqualizerFlat          Equalizer = 2
	EqualizerBoost         Equalizer = 3
	EqualizerTrebleBass    Equalizer = 4
	EqualizerUser          Equalizer = 5
	EqualizerMusic         Equalizer = 6
	EqualizerCinema        Equalizer = 7
	EqualizerNight         Equalizer = 8
	EqualizerNews          Equalizer = 9
	EqualizerVoice         Equalizer = 10
	EqualizerISound        Equalizer = 11
	EqualizerASC           Equalizer = 12
	EqualizerMovie         Equalizer = 13
	EqualizerBassBlast     Equalizer = 14
	EqualizerDolbyAtmos    Equalizer = 15
	EqualizerDTSVirtualX   Equalizer = 16
	EqualizerBassBoostPlus Equalizer = 17
)

func (e Equalizer) String() string {
	switch e {
	case EqualizerStandard:
		return "Standard"
	case EqualizerBass:
		return "Bass"
	case EqualizerFlat:
		return "Flat"
	case EqualizerBoost:
		return "Boost"
	case EqualizerTrebleBass:
		return "Treble and Bass"
	case EqualizerUser:
		return "User"
	case EqualizerMusic:
		return "Music"
	case EqualizerCinema:
		return "Cinema"
	case EqualizerNight:
		return "Night"
	case EqualizerNews:
		return "News"
	case EqualizerVoice:
		return "Voice"
	case EqualizerISound:
		return "ISound"
	case EqualizerASC:
		return "ASC"
	case EqualizerMovie:
		return "Movie"
	case EqualizerBassBlast:
		return "Bass Blast"
	case EqualizerDolbyAtmos:
		return "Dolby Atmos"
	case EqualizerDTSVirtualX:
		return "DTS Virtual X"
	case EqualizerBassBoostPlus:
		return "Bass Boost Plus"
	default:
		return fmt.Sprintf("Equalizer(%d)", e)
	}
}

// EqualizerType is a setting type for changing equalizer settings.
type EqualizerType int

// EqualizerType enums.
// {"data": {"type": 1, "value": 5}, "msg": "EQ_SETTING"}
// {"data": {"type": 3, "value": 20}, "msg": "EQ_SETTING"}
const (
	SetEqualizer        EqualizerType = 0 // Sets EQ, e.g. Cinema, ASC, etc.
	SetBass             EqualizerType = 1
	SetTreble           EqualizerType = 2
	SetLeftRightBalance EqualizerType = 3
	// Save or restore equalizer setting. Value 1 saves, 0 restores.
	// Reads of EQ_INFO_REQ will echo the previously saved value
	// until the new values are saved.
	// 2021-04-17 Update: On firmware NB9.504.00629.C EQ_INFO_REQ
	// seems to return the current value, saved or not.
	SetSaveRestore EqualizerType = 4
)

func (t EqualizerType) String() string {
	switch t {
	case SetEqualizer:
		return "Equalizer"
	case SetBass:
		return "Bass"
	case SetTreble:
		return "Treble"
	case SetLeftRightBalance:
		return "LeftRightBalance"
	case SetSaveRestore:
		return "SaveRestore"
	default:
		return fmt.Sprintf("EqualizerType(%d)", t)
	}
}

// Function represents a function mode for the speaker. A list of
// available modes are available when requesting product info.
type Function int

// Function mode enums.
// {"data": {"type": 7}, "msg": "FUNCTION_SET"}
const (
	FunctionWiFi       Function = 0
	FunctionBluetooth  Function = 1
	FunctionPortable   Function = 2
	FunctionAUX        Function = 3
	FunctionOptical    Function = 4
	FunctionCP         Function = 5
	FunctionHDMI       Function = 6
	FunctionARC        Function = 7
	FunctionSpotify    Function = 8
	FunctionOptical2   Function = 9
	FunctionHDMI2      Function = 10
	FunctionHDMI3      Function = 11
	FunctionLGTV       Function = 12
	FunctionMic        Function = 13
	FunctionC4A        Function = 14
	FunctionOpticalARC Function = 15
	FunctionLGOptical  Function = 16
	FunctionFM         Function = 17
	FunctionUSB        Function = 18
)

func (f Function) String() string {
	switch f {
	case FunctionWiFi:
		return "WiFi"
	case FunctionBluetooth:
		return "Bluetooth"
	case FunctionPortable:
		return "Portable"
	case FunctionAUX:
		return "AUX"
	case FunctionOptical:
		return "Optical"
	case FunctionCP:
		return "CP"
	case FunctionHDMI:
		return "HDMI"
	case FunctionARC:
		return "ARC"
	case FunctionSpotify:
		return "Spotify"
	case FunctionOptical2:
		return "Optical2"
	case FunctionHDMI2:
		return "HDMI2"
	case FunctionHDMI3:
		return "HDMI3"
	case FunctionLGTV:
		return "LGTV"
	case FunctionMic:
		return "Microphone"
	case FunctionC4A:
		return "C4A"
	case FunctionOpticalARC:
		return "Optical / HDMI ARC"
	case FunctionLGOptical:
		return "LG Optical"
	case FunctionFM:
		return "FM"
	case FunctionUSB:
		return "USB"
	default:
		return fmt.Sprintf("Function(%d)", f)
	}
}

// Model represents the speaker model.
type Model int

// Model type (product info "modeltype") enums.
const (
	ModelBridge    Model = 0
	ModelBasic     Model = 1
	ModelSoundBar  Model = 2
	ModelMono      Model = 3
	ModelConnector Model = 4
	ModelPortable  Model = 5
)

func (m Model) String() string {
	switch m {
	case ModelBridge:
		return "Bridge"
	case ModelBasic:
		return "Basic"
	case ModelSoundBar:
		return "SoundBar"
	case ModelMono:
		return "Mono"
	case ModelConnector:
		return "Connector"
	case ModelPortable:
		return "Portable"
	default:
		return fmt.Sprintf("Model(%d)", m)
	}
}

// Network represents the connection type.
type Network int

// Network type (product info "network") enums.
const (
	NetworkWired    Network = 0
	NetworkWireless Network = 1
	NetworkMeshed   Network = 2
)

func (n Network) String() string {
	switch n {
	case NetworkWired:
		return "Wired"
	case NetworkWireless:
		return "Wireless"
	case NetworkMeshed:
		return "Meshed"
	default:
		return fmt.Sprintf("Network(%d)", n)
	}
}

// Role represents the speakers role.
type Role int

// Speaker role (product info "spktype") enums.
const (
	RoleIndividual     Role = 0
	RoleMaster         Role = 1
	RoleSlave          Role = 2
	RoleSurroundMaster Role = 3
	RoleSurroundSlave  Role = 4
)

func (r Role) String() string {
	switch r {
	case RoleIndividual:
		return "Individual"
	case RoleMaster:
		return "Master"
	case RoleSlave:
		return "Slave"
	case RoleSurroundMaster:
		return "Surround Master"
	case RoleSurroundSlave:
		return "Surround Slave"
	default:
		return fmt.Sprintf("Role(%d)", r)
	}
}
