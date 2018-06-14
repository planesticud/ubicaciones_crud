package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RelacionLugares struct {
	Id                int                `orm:"column(id);pk;auto"`
	LugarPadre        *Lugar             `orm:"column(lugar_padre);rel(fk)"`
	LugarHijo         *Lugar             `orm:"column(lugar_hijo);rel(fk)"`
	TipoRelacionLugar *TipoRelacionLugar `orm:"column(tipo_relacion_lugar);rel(fk)"`
	Activo            bool               `orm:"column(activo)"`
}

func (t *RelacionLugares) TableName() string {
	return "relacion_lugares"
}

func init() {
	orm.RegisterModel(new(RelacionLugares))
}

// AddRelacionLugares insert a new RelacionLugares into database and returns
// last inserted Id on success.
func AddRelacionLugares(m *RelacionLugares) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRelacionLugaresById retrieves RelacionLugares by Id. Returns error if
// Id doesn't exist
func GetRelacionLugaresById(id int) (v *RelacionLugares, err error) {
	o := orm.NewOrm()
	v = &RelacionLugares{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRelacionLugares retrieves all RelacionLugares matches certain condition. Returns empty list if
// no records exist
func GetAllRelacionLugares(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RelacionLugares)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RelacionLugares
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRelacionLugares updates RelacionLugares by Id and returns error if
// the record to be updated doesn't exist
func UpdateRelacionLugaresById(m *RelacionLugares) (err error) {
	o := orm.NewOrm()
	v := RelacionLugares{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRelacionLugares deletes RelacionLugares by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRelacionLugares(id int) (err error) {
	o := orm.NewOrm()
	v := RelacionLugares{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RelacionLugares{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetRelacionesPadre retrieves padre de un lugar
func GetRelacionesPadre(id int) (v []RelacionLugares) {
	o := orm.NewOrm()
	if _, err := o.Raw(`Select lugar_padre,lugar_hijo from core_new.relacion_lugares 
		where lugar_hijo=` + strconv.Itoa(id)).QueryRows(&v); err == nil {
	}
	return v
}

// GetJerarquiaLugarById retrieves JerarquiaLugar by Id. Returns error if
func GetJerarquiaLugarById(id int) (v map[string]interface{}, err error) {

	var relacionesLugares []RelacionLugares
	var resultado map[string]interface{}
	resultado = make(map[string]interface{})

	//buscar en lugar el Id de tipo_lugar
	o := orm.NewOrm()
	lugar := &Lugar{Id: id}
	if err = o.Read(lugar); err == nil {
		if _, err := o.Raw(`Select lugar_padre,lugar_hijo from core_new.relacion_lugares 
			where lugar_hijo=` + strconv.Itoa(id)).QueryRows(&relacionesLugares); err == nil {
			if relacionesLugares == nil { //no tiene padre
				err := o.Read(lugar.TipoLugar)
				if err == nil {
					resultado[lugar.TipoLugar.Nombre] = lugar
				}

			} else {
				for relacionesLugares != nil {
					lugar := Lugar{Id: relacionesLugares[0].LugarHijo.Id}
					err = o.Read(&lugar)
					err := o.Read(lugar.TipoLugar)
					if err == nil {
						resultado[lugar.TipoLugar.Nombre] = lugar
					}

					padre := relacionesLugares[0].LugarPadre.Id
					relacionesLugares = GetRelacionesPadre(relacionesLugares[0].LugarPadre.Id)

					if relacionesLugares == nil {
						lugar := Lugar{Id: padre}
						err = o.Read(&lugar)
						err := o.Read(lugar.TipoLugar)
						if err == nil {
							resultado[lugar.TipoLugar.Nombre] = lugar
						}
					}
				}
			}

		}
		return resultado, nil
	}
	return nil, err
}
