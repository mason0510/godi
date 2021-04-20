package demo

import (
	"database/sql"
)

type A struct {
	Db  *sql.DB `di:"db"`
	Db1 *sql.DB `di:"db"`
	B   *B      `di:"b,prototype"`
	B1  *B      `di:"b,prototype"`
}

func NewA() *A {
	return &A{}
}

//单例对象在整个容器中只有一个实例，所以不管在何处注入，获取到的指针一定是一样的。
//实例对象是通过同一个工厂方法创建的，所以每个实例的指针不可以相同。
func (p *A) Version() (string, error) {
	rows, err := p.Db.Query("SELECT VERSION() as version")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var version string
	if rows.Next() {
		if err := rows.Scan(&version); err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	return version, nil
}