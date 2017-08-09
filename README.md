# Pizzaman - Order your pizza !

## Order a pizza

`curl localhost:9876/yourSauce/yourMeat/yourCheese`

If you want one of your pizza's part to be skipped, just say "no"

### Examples

#### Pizza with only ketchup and emmental

`curl localhost:9876/ketchup/no/emmental`

#### Pizza with buffalo sauce, bacon and mozzarella

`curl localhost:9876/buffalo/bacon/mozzarella`

### Fridge

Available sauces:

	- tomato
	- ketchup
	- classico
	- italian
	- buffalo
	- supreme

Available meats:

	- chicken
	- bacon

Available cheeses :

	- goat
	- mozzarella
	- emmental
	- ricotta
	- provolone
	- cheddar
	- romano

## Ask questions

You can ask question to the pizzaman, as simply as `curl localhost:9876/yourQuestionFormat`


| Question | Format |
| --- | --- |
|What sauces do you have ? | whatsaucesdoyouhave |
|What meats do you have ? | whatmeatsdoyouhave |
|What cheeses do you have ? | whatcheesesdoyouhave |
