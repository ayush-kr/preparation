package coffeemachine

//Beverage interface declares methods for a Beverage item
type Beverage interface {
	Prepare() error
}

//BeverageItem implements Beverage interface
type BeverageItem struct {
	ingredients map[string]int
	name        string
	machine     BeverageMaker
}

//Prepare makes the beverage if all the required ingredients are available in required amount and throws error otherwise
func (bi *BeverageItem) Prepare() error {
	for k, v := range bi.ingredients {
		err := bi.machine.ConsumeIngredient(k, v, false)
		if err != nil {
			return err
		}
	}
	for k, v := range bi.ingredients {
		_ = bi.machine.ConsumeIngredient(k, v, true)
	}
	return nil
}
