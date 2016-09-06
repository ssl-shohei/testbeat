package beater

import (
	"fmt"
	"time"
        "os"
        "os/exec"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/ssl-shohei/testbeat/config"
)

type Testbeat struct {
	beatConfig *config.Config
	done       chan struct{}
	period     time.Duration
	client     publisher.Client
}

// Creates beater
func New() *Testbeat {
	return &Testbeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Testbeat) Config(b *beat.Beat) error {

	// Load beater beatConfig
	err := b.RawConfig.Unpack(&bt.beatConfig)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	return nil
}

func (bt *Testbeat) Setup(b *beat.Beat) error {

	// Setting default period if not set
	if bt.beatConfig.Testbeat.Period == "" {
		bt.beatConfig.Testbeat.Period = "1s"
	}

	bt.client = b.Publisher.Connect()

	var err error
	bt.period, err = time.ParseDuration(bt.beatConfig.Testbeat.Period)
	if err != nil {
		return err
	}

	return nil
}

func (bt *Testbeat) Run(b *beat.Beat) error {
	logp.Info("testbeat is running! Hit CTRL-C to stop it.")

        ticker := time.NewTicker(bt.period)
        counter := 1
        for {
                select {
                case <-bt.done:
                        return nil
                case <-ticker.C:
                }

                // コマンド実行
                cmd := exec.Command(bt.beatConfig.Testbeat.Command)
                out, err := cmd.Output()

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                event := common.MapStr{
                        "@timestamp": common.Time(time.Now()),
                        "type":       b.Name,
                        "counter":    counter,
                        "use%":       out,
                }
                bt.client.PublishEvent(event)
                logp.Info("Event sent")
                counter++
        }
}

func (bt *Testbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Testbeat) Stop() {
	close(bt.done)
}
