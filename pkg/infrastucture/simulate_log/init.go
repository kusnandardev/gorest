package simulate_log

import "os"

func InitLogFile() error {
	f1, err := os.Create("trx_log/transaction.log")
	if err != nil {
		return err
	}
	defer f1.Close()
	f2, err := os.Create("trx_log/movement.log")
	if err != nil {
		return err
	}
	defer f2.Close()
	f1.WriteString("id | type | src | dest | amount | date \n")
	f2.WriteString("id | user | trxId | drcr | amount | blncbefore | blncafter \n")
	return nil
}
