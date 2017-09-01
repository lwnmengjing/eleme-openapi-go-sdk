package elemeOpenApi

type APIClient struct {
	Message  Message
	Order    Order
	Shop     Shop
	Product  Product
	User     User
	File     File
	Packs    Packs
	Finance  Finance
  Activity Activity
  UGC      UGC
	config   Config
}

type User struct {
	config *Config
}

type Shop struct {
	config *Config
}

type Product struct {
	Item     Item
	Category Category
	config   *Config
}

type Order struct {
	config *Config
}

type Message struct {
	config *Config
}

type File struct {
	config *Config
}

type Item struct {
	config *Config
}

type Category struct {
	config *Config
}

type Packs struct {
	config *Config
}

type Finance struct {
	config *Config
}

type Flash struct {
  config *Config
}

type Activity struct {
  Flash Flash
  config *Config
}

type UGC struct {
  config *Config
}

func NewAPIClient(config Config) APIClient {
	client := APIClient{}
	client.SetConfig(config)
	return client
}

func (client *APIClient) SetConfig(config Config) {
	client.config = config
	client.Message.config = &client.config
	client.Order.config = &client.config
	client.Product.config = &client.config
	client.Product.Item.config = &client.config
	client.Product.Category.config = &client.config
	client.User.config = &client.config
	client.Shop.config = &client.config
	client.File.config = &client.config
	client.Packs.config = &client.config
	client.Finance.config = &client.config
  client.Activity.config = &client.config
  client.Activity.flash.config = &client.config
  client.UGC.config = &client.config
}

