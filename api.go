package musicflow

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	errors "golang.org/x/xerrors"

	"github.com/mafredri/musicflow/api"
)

// Request to the speaker.
type Request struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg"`
}

func (r Request) String() string {
	return fmt.Sprintf("Request{Message=%s, Data=%#v}", r.Message, r.Data)
}

func newRequest(data interface{ Message() string }) Request {
	return Request{
		Data:    data,
		Message: data.Message(),
	}
}

// Response from the speaker.
type Response struct {
	Data    json.RawMessage `json:"data,omitempty"`
	Message string          `json:"msg"`
	Result  string          `json:"result,omitempty"`
}

func (r Response) String() string {
	return fmt.Sprintf("Response{Result=%s, Message=%s, Data=%s}", r.Result, r.Message, r.Data)
}

// OnBroadcast sets the handler for broadcasted messages.
func (c *Client) OnBroadcast(fn func(message string, data []byte)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.broadcast = fn
}

// SetName sets the speakers name.
func (c *Client) SetName(ctx context.Context, name string) error {
	req := api.SpeakerInfoModifyRequest{
		Name: name,
		Icon: 0,
	}
	err := c.Send(ctx, newRequest(req), nil)
	if err != nil {
		return errors.Errorf("SetName failed: %w", err)
	}
	return nil
}

// ProductInfo returns the product info and optionally sets the time on
// the soundbar.
func (c *Client) ProductInfo(ctx context.Context, now time.Time, setTime bool) (*api.ProductInfo, error) {
	option := 0
	if setTime {
		option = 1
	}
	req := api.ProductInfoRequest{
		ID:     clientID,
		Day:    api.FromWeekday(now.Weekday()),
		Hour:   now.Hour(),
		Min:    now.Minute(),
		Option: option,
	}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("ProductInfo failed: %w", err)
	}
	return reply, nil
}

// Settings returns the system version information.
func (c *Client) Settings(ctx context.Context) (*api.Settings, error) {
	req := api.SettingInfoRequest{}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("Settings failed: %w", err)
	}
	return reply, nil
}

// SystemVersion returns the system version information.
func (c *Client) SystemVersion(ctx context.Context) (*api.SystemVersion, error) {
	req := api.SystemVersionRequest{}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("SystemVersion failed: %w", err)
	}
	return reply, nil
}

// NetworkInfo returns the network information.
//
// WARNING: Returns WiFi password in plain text.
func (c *Client) NetworkInfo(ctx context.Context) (*api.NetworkInfo, error) {
	req := api.NetworkInfoRequest{}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("NetworkInfo failed: %w", err)
	}
	return reply, nil
}

// PlayInfo returns information on what's playing.
func (c *Client) PlayInfo(ctx context.Context) (*api.PlayInfo, error) {
	req := api.PlayInfoRequest{}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("PlayInfo failed: %w", err)
	}
	return reply, nil
}

// EqualizerInfo returns information on what's playing.
func (c *Client) EqualizerInfo(ctx context.Context) (*api.EqualizerInfo, error) {
	req := api.EqualizerInfoRequest{}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("EqualizerInfo failed: %w", err)
	}
	return reply, nil
}

// EqualizerSetting represents an equalizer setting.
type EqualizerSetting func(context.Context, *Client) error

// SetBass sets the bass level.
func SetBass(value int) EqualizerSetting {
	return func(ctx context.Context, c *Client) error {
		req := api.EqualizerSetRequest{Type: api.SetBass, Value: value}
		err := c.Send(ctx, newRequest(req), nil)
		if err != nil {
			return errors.Errorf("SetBass failed: %w", err)
		}
		return nil
	}
}

// SetEqualizer sets the equalizer.
func SetEqualizer(value api.Equalizer) EqualizerSetting {
	return func(ctx context.Context, c *Client) error {
		req := api.EqualizerSetRequest{Type: api.SetBass, Value: int(value)}
		err := c.Send(ctx, newRequest(req), nil)
		if err != nil {
			return errors.Errorf("SetEqualizer failed: %w", err)
		}
		return nil
	}
}

// SetLeftRightBalance sets the left right balance.
func SetLeftRightBalance(value int) EqualizerSetting {
	return func(ctx context.Context, c *Client) error {
		req := api.EqualizerSetRequest{Type: api.SetLeftRightBalance, Value: value}
		err := c.Send(ctx, newRequest(req), nil)
		if err != nil {
			return errors.Errorf("SetLeftRightBalance failed: %w", err)
		}
		return nil
	}
}

// SetTreble sets treble.
func SetTreble(value int) EqualizerSetting {
	return func(ctx context.Context, c *Client) error {
		req := api.EqualizerSetRequest{Type: api.SetTreble, Value: value}
		err := c.Send(ctx, newRequest(req), nil)
		if err != nil {
			return errors.Errorf("SetTreble failed: %w", err)
		}
		return nil
	}
}

// SaveEqualizerSettings saves the current equalizer settings as default.
func SaveEqualizerSettings() EqualizerSetting {
	return func(ctx context.Context, c *Client) error {
		req := api.EqualizerSetRequest{Type: api.SetSaveRestore, Value: 1}
		err := c.Send(ctx, newRequest(req), nil)
		if err != nil {
			return errors.Errorf("SaveEqualizerSettings failed: %w", err)
		}
		return nil
	}
}

// RestoreEqualizerSettings restores the previously saved equalizer settings.
func RestoreEqualizerSettings() EqualizerSetting {
	return func(ctx context.Context, c *Client) error {
		req := api.EqualizerSetRequest{Type: api.SetSaveRestore, Value: 0}
		err := c.Send(ctx, newRequest(req), nil)
		if err != nil {
			return errors.Errorf("RestoreEqualizerSettings failed: %w", err)
		}
		return nil
	}
}

// Equalizer sets the provided equalizer settings.
func (c *Client) Equalizer(ctx context.Context, eq ...EqualizerSetting) error {
	for _, e := range eq {
		err := e(ctx, c)
		if err != nil {
			return errors.Errorf("Equalizer failed: %w", err)
		}
	}
	return nil
}

// FunctionInfo sets the provided equalizer settings.
func (c *Client) FunctionInfo(ctx context.Context) (*api.FunctionInfo, error) {
	req := api.FunctionInfoRequest{}
	reply := req.Reply()

	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("FunctionInfo failed: %w", err)
	}
	return reply, nil
}

// Function activates the provided function.
func (c *Client) Function(ctx context.Context, f api.Function) error {
	req := api.FunctionSetRequest{Type: f}
	err := c.Send(ctx, newRequest(req), nil)
	if err != nil {
		return errors.Errorf("EqualizerInfo failed: %w", err)
	}
	return nil
}

// NightMode sets the nightmode on or off.
func (c *Client) NightMode(ctx context.Context, on bool) error {
	req := api.NightModeSetRequest{NightMode: on}
	reply := req.Reply()
	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return errors.Errorf("NightMode failed: %w", err)
	}
	if reply.NightMode != on {
		return errors.New("NightMode: wrong return value")
	}
	return nil
}

// Volume sets the volume.
func (c *Client) Volume(ctx context.Context, volume, fadetime int) error {
	req := api.VolumeSettingRequest{Volume: volume, FadeTime: fadetime}
	err := c.Send(ctx, newRequest(req), nil)
	if err != nil {
		return errors.Errorf("Volume failed: %w", err)
	}
	return nil
}

// WooferLevel sets the woofer level.
func (c *Client) WooferLevel(ctx context.Context, level int) error {
	req := api.WooferLevelSetRequest{Level: level}
	reply := req.Reply()
	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return errors.Errorf("WooferLevel failed: %w", err)
	}
	if reply.Level != level {
		return errors.New("WooferLevel: wrong return value")
	}
	return nil
}

// Mute the speaker.
func (c *Client) Mute(ctx context.Context, on bool) error {
	req := api.MuteSetRequest{Mute: on}
	err := c.Send(ctx, newRequest(req), nil)
	if err != nil {
		return errors.Errorf("Mute failed: %w", err)
	}
	return nil
}

// Alarms lists all alarms.
func (c *Client) Alarms(ctx context.Context) ([]api.Alarm, error) {
	req := api.AlarmListRequest{}
	reply := req.Reply()
	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return nil, errors.Errorf("Alarms failed: %w", err)
	}
	return reply.Info, nil
}

// AlarmCreate creates a new alarm.
func (c *Client) AlarmCreate(ctx context.Context, a api.Alarm) (id int, err error) {
	a.ID = -1
	a.Mode = api.AlarmCreate
	a.Day = api.AlarmDays(time.Monday, time.Tuesday)
	req := api.AlarmSetRequest{
		Alarm: a,
	}
	reply := req.Reply()
	err = c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return 0, errors.Errorf("SetAlarm failed: %w", err)
	}
	return reply.ID, nil
}

// AlarmDelete deletes the alarm.
func (c *Client) AlarmDelete(ctx context.Context, a api.Alarm) error {
	a.Mode = api.AlarmDelete
	req := api.AlarmSetRequest{
		Alarm: a,
	}
	reply := req.Reply()
	err := c.Send(ctx, newRequest(req), reply)
	if err != nil {
		return errors.Errorf("AlarmDelete failed: %w", err)
	}
	return nil
}

// AlarmState returns true if an alarm is active right now.
func (c *Client) AlarmState(ctx context.Context) (on bool, err error) {
	req := api.AlarmStateRequest{}
	reply := req.Reply()
	err = c.Send(ctx, newRequest(req), reply, WaitFor(api.MessageAlarmStateNotification, ""))
	if err != nil {
		return false, errors.Errorf("AlarmState failed: %w", err)
	}
	return reply.On, nil
}

// SleepAfter sets the sleep timer. Set -1 to disable.
func (c *Client) SleepAfter(ctx context.Context, minutes int) error {
	req := api.SleepSetRequest{Time: minutes}
	err := c.Send(ctx, newRequest(req), nil)
	if err != nil {
		return errors.Errorf("SleepAfter failed: %w", err)
	}
	return nil
}

// TestTone plays the test tone.
func (c *Client) TestTone(ctx context.Context) error {
	req := api.TestToneRequest{Stat: true}
	err := c.Send(ctx, newRequest(req), nil)
	if err != nil {
		return errors.Errorf("TestTone failed: %w", err)
	}
	return nil
}
