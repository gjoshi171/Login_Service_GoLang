package models

//Expenses is a representation of a peExpensionrson
type Expenses struct {
	ExpenseDate   string `bson:"expenseDate" json:"expenseDate"`
	Item          string `bson:"item" json:"item"`
	ExpenseAmount int    `bson:"expenseAmount" json:"expenseAmount"`
	ForWhom       string `bson:"forWhom" json:"forWhom"`
	ByWhom        string `bson:"byWhom" json:"byWhom"`
}

func (e Expenses) GetExpenseDate() string {
	return e.ExpenseDate
}

func (e *Expenses) SetExpensedate(ExpenseDate string) {
	e.ExpenseDate = ExpenseDate
}
func (e Expenses) GetItem() string {
	return e.Item
}

func (e *Expenses) SetItem(Item string) {
	e.Item = Item
}
func (e Expenses) GetExpenseAmount() int {
	return e.ExpenseAmount
}

func (e *Expenses) SetExpenseAmount(ExpenseAmount int) {
	e.ExpenseAmount = ExpenseAmount
}

func (e *Expenses) GetForWhom() string {
	return e.ForWhom
}

func (e *Expenses) SetForWhom(ForWhom string) {
	e.ForWhom = ForWhom
}

func (e *Expenses) GetByWhom() string {
	return e.ForWhom
}

func (e *Expenses) SetByWhom(ByWhom string) {
	e.ByWhom = ByWhom
}
