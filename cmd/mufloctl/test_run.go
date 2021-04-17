package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/mafredri/goodspeaker"
	"github.com/mafredri/musicflow"
	"github.com/mafredri/musicflow/api"
)

func testRun(ctx context.Context, addr, key, iv string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Printf("Connecting to %s...", addr)

	var gsOpt []goodspeaker.Option
	if key != "" && iv != "" {
		aes, err := goodspeaker.WithAES([]byte(key), []byte(iv))
		if err != nil {
			return err
		}
		gsOpt = append(gsOpt, aes)
	}
	opt := []musicflow.DialOption{
		musicflow.WithGoodspeakerOption(gsOpt...),
		musicflow.WithLogger(log.New(os.Stderr, "[musicflow] ", log.Flags())),
	}

	c, err := musicflow.Dial(ctx, addr, opt...)
	if err != nil {
		return err
	}
	defer c.Close()

	c.OnBroadcast(func(message string, data []byte) {
		log.Printf("Broadcast: %s %s", message, data)
	})

	_, err = c.ProductInfo(ctx, time.Now(), true)
	if err != nil {
		return err
	}
	prodinfo, err := c.ProductInfo(ctx, time.Now(), false)
	if err != nil {
		return err
	}
	log.Printf("%#v", prodinfo)

	sysver, err := c.SystemVersion(ctx)
	if err != nil {
		return err
	}
	log.Printf("%#v", sysver)

	netinfo, err := c.NetworkInfo(ctx)
	if err != nil {
		return err
	}
	log.Printf("%#v", netinfo)

	playinfo, err := c.PlayInfo(ctx)
	if err != nil {
		return err
	}
	log.Printf("%#v", playinfo)

	settings, err := c.Settings(ctx)
	if err != nil {
		return err
	}

	eqinfo, err := c.EqualizerInfo(ctx)
	if err != nil {
		return err
	}

	// Change equalizer settings (without saving).
	err = c.Equalizer(ctx,
		musicflow.SetBass(eqinfo.Bass+2),
		musicflow.SetTreble(eqinfo.Treble+2),
		musicflow.SetEqualizer(api.EqualizerASC),
		musicflow.SetLeftRightBalance(eqinfo.LeftRightBalance-2),
	)
	if err != nil {
		return err
	}

	// Check that they've changed.
	eqinfo2, err := c.EqualizerInfo(ctx)
	if err != nil {
		return err
	}
	if *eqinfo == *eqinfo2 {
		panic("eqinfo did not change")
	}

	err = c.Equalizer(ctx, musicflow.RestoreEqualizerSettings())
	if err != nil {
		return err
	}

	eqinfo3, err := c.EqualizerInfo(ctx)
	if err != nil {
		return err
	}
	if *eqinfo != *eqinfo3 {
		panic("eqinfo did not restore")
	}

	funcinfo, err := c.FunctionInfo(ctx)
	if err != nil {
		return err
	}
	log.Printf("%#v", funcinfo)

	err = c.Function(ctx, api.FunctionARC)
	if err != nil {
		return err
	}

	err = c.NightMode(ctx, !settings.NightMode)
	if err != nil {
		return err
	}
	err = c.NightMode(ctx, settings.NightMode)
	if err != nil {
		return err
	}

	err = c.Volume(ctx, prodinfo.Info.Volume-1, 0)
	if err != nil {
		return err
	}
	err = c.Volume(ctx, prodinfo.Info.Volume, 0)
	if err != nil {
		return err
	}

	err = c.Mute(ctx, !prodinfo.Info.Mute)
	if err != nil {
		return err
	}
	err = c.Mute(ctx, prodinfo.Info.Mute)
	if err != nil {
		return err
	}

	err = c.SetName(ctx, prodinfo.Info.Name+"-test")
	if err != nil {
		return err
	}
	err = c.SetName(ctx, prodinfo.Info.Name)
	if err != nil {
		return err
	}

	err = c.WooferLevel(ctx, settings.WooferLevel)
	if err != nil {
		return err
	}

	alarmOn, err := c.AlarmState(ctx)
	if err != nil {
		return err
	}
	log.Printf("AlarmSate: %v", alarmOn)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			req := musicflow.Request{}
			in := s.Bytes()
			err = json.Unmarshal(in, &req)
			if err != nil {
				log.Println(err)
				continue
			}

			err = c.Send(ctx, req, nil)
			if err != nil {
				if err == io.EOF {
					log.Println("Connection lost")
					break
				}
				log.Println(err)
				cancel()
				return
			}
		}
		if err := s.Err(); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
