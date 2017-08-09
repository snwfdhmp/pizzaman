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

## See your pizza's state

Ask the pizzaman by send a request on `howismypizza/yourTicketNumber`

Example :

```
$ curl 0:9876/order/supreme/chicken/goat
{"message":"Your order has been taken","pizza":{"sauce":"supreme","meat":"chicken","cheese":"goat"},"ticket":"1fb1"}
```

My ticket number is 1fb1, so let's ask the pizzaman

```
$ curl 0:9876/howismypizza/1fb1
pending
```

A few seconds later ...

```
curl 0:9876/howismypizza/1fb1
prepared
```

## Ask questions

You can ask questions to the pizzaman, as simply as `curl localhost:9876/yourQuestionFormat`


| Question | Format |
| --- | --- |
|What sauces do you have ? | whatsaucesdoyouhave |
|What meats do you have ? | whatmeatsdoyouhave |
|What cheeses do you have ? | whatcheesesdoyouhave |
