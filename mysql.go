package smsoh

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func (m *Middleware) mysqlInsert(ud, scts, oa, da string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", m.Username, m.Password, m.Database))
	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO inbox (ReceivingDateTime, Text, SenderNumber, RecipientID, UDH, TextDecoded) VALUES( ?, ?, ?, ?, '', '' )")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(scts, ud, oa, da)
	if err != nil {
		return err
	}

	return nil
}
