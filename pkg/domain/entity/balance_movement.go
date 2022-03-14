package entity

type BalanceMovement struct {
	Id            int
	UserId        string
	TransactionId int
	Drcr          int // debit = 1, credit =2
	Amount        int
	BalanceBefore int
	BalanceAfter  int
}
