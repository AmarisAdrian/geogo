// models/task.go
package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PrediosModel struct {
	TID                           int            `json:"t_id"`
	TIliTID                       uuid.UUID      `json:"t_ili_tid"`
	TipoCatastro                  sql.NullString `json:"tipo_catastro"`
	NumeroPredial                 sql.NullString `json:"numero_predial"`
	NumeroPredialAnterior         sql.NullString `json:"numero_predial_anterior"`
	Nupre                         sql.NullString `json:"nupre" gorm:"type:varchar(11)"`
	CirculoRegistral              sql.NullString `json:"circulo_registral"`
	MatriculaInmobiliariaCatastro sql.NullString `json:"matricula_inmobiliaria_catastro"`
	TipoPredio                    sql.NullString `json:"tipo_predio"`
	CondicionPredio               sql.NullInt64  `json:"condicion_predio"`
	DestinacionEconomica          sql.NullString `json:"destinacion_economica"`
	SistemaProcedenciaDatos       sql.NullString `json:"sistema_procedencia_datos"`
	FechaDatos                    time.Time      `json:"fecha_datos"`
}

func (PrediosModel) TableName() string {
	return "ladm.gc_prediocatastro"
}
