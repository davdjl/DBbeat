package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/davdjl/dbbeat/config"
)

// dbbeat configuration.
type dbbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of dbbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &dbbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts dbbeat.
func (bt *dbbeat) Run(b *beat.Beat) error {
	logp.Info("dbbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	
	if err != nil {
		logp.Info("error: %s",err)
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	startConection(bt.config.Server,bt.config.Port,bt.config.User,bt.config.Password,bt.config.Database)
	//startConection()
	
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		// Read employees
		//producto, err := ReadEmployees(bt.config.Query)
		producto, err := ReadEmployees(bt.config.Query)
		if err != nil {
			logp.Error(err)
		}
		if producto != nil {
			for _,p:= range producto{
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"type": b.Info.Name,
						"uiCodSucursal": p.uiCodSucursal,
						"uiCodProducto": p.uiCodProducto,
						"bExistenciaINV": p.bExistenciaINV,
						"bExistenciaPro": p.bExistenciaPro,
						"nExistencia": p.nExistencia,
						"nExistenciaApartada": p.nExistenciaApartada,
						"nExistenciaDisponible": p.nExistenciaDisponible,
						"bAgotado": p.bAgotado,
						"bDescontinuado": p.bDescontinuado,
						"bAlertaSanitaria": p.bAlertaSanitaria,
						"dFechaVenta": p.dFechaVenta,
						"cControlCOFEPRIS": p.cControlCOFEPRIS,
					},
				}
				bt.client.Publish(event)
				//logp.Info(fmt.Sprintf("Producto-sucursal (%s %s) enviado",p.uiCodProducto, p.uiCodSucursal))
			}
		} else {
			logp.Info("Estas al dia!")
		}
	}
}

// Stop stops dbbeat.
func (bt *dbbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
