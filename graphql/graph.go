package main

// type Server struct{
// 	accountClient *account.Client
// 	productClient *product.Client
// 	orderClient *order.Client
// }

func NewGraphQLServer(accountUrl String, productUrl String, orderUrl String)(*Server, error){
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	productClient, err := product.NewClient(productUrl)
	if err != nil {
		return nil, err
	}
	orderClient. err := order.NewClient(orderUrl)
	if err != nil {
		return nil, err
	}
}